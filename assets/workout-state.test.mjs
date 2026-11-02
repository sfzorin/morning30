import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import vm from 'node:vm';

const source = readFileSync(new URL('./workout-state.js', import.meta.url), 'utf8');
const plain = value => JSON.parse(JSON.stringify(value));
const payload = (fields = {}) => ({ userId: 1, position: 0, day: 1, programKey: 'plan-a', resumeKey: 'plan-a:0:1', ...fields });
const event = (fields = {}) => ({ eventId: 'a'.repeat(32), userId: 1, day: 1, position: 0, completedAt: '2026-10-24T21:30:00Z', completedOn: '2026-10-25', ...fields });
const state = (fields = {}) => ({ index: 3, side: 1, phase: 'rest', remainingMs: 7000, nextIndex: 3, nextSide: 1, level: 0, swaps: { 3: 'replacement' }, eventId: 'a'.repeat(32), completion: null, ...fields });
const response = (e, fields = {}) => ({ ok: true, status: 200, json: async () => ({ saved: true, eventId: e.eventId, streak: 4, position: 1, ...fields }) });

function browser({ disk = new Map(), fetch = async (_url, options) => response(JSON.parse(options.body)), storageFails = false } = {}) {
  const events = new Map(), timers = new Map();
  let nextTimer = 1;
  const listen = (name, fn) => {
    if (!events.has(name)) events.set(name, []);
    events.get(name).push(fn);
  };
  const storage = {
    get length() { if (storageFails) throw Error('storage blocked'); return disk.size; },
    key(i) { if (storageFails) throw Error('storage blocked'); return [...disk.keys()][i]; },
    getItem(key) { if (storageFails) throw Error('storage blocked'); return disk.get(key) ?? null; },
    setItem(key, value) { if (storageFails) throw Error('storage full'); disk.set(key, value); },
    removeItem(key) { if (storageFails) throw Error('storage blocked'); disk.delete(key); },
  };
  const document = { hidden: false, visibilityState: 'visible', addEventListener: listen };
  const window = { localStorage: storage, fetch, addEventListener: listen, navigator: { onLine: true } };
  const schedule = (fn, delay) => { const id = nextTimer++; timers.set(id, { fn, delay }); return id; };
  const context = vm.createContext({ window, document, AbortController, setTimeout: schedule, clearTimeout: id => timers.delete(id), setInterval: schedule });
  vm.runInContext(source, context);
  return { api: window.MorningWorkoutState, disk, events, timers, document,
    emit(name, value = {}) { events.get(name)?.forEach(fn => fn(value)); },
    reloadScript() { vm.runInContext(source, context); },
  };
}

test('checkpoint survives reload with exact side, remaining time and replacement; stale identity is discarded per account', () => {
  const a = browser();
  assert.equal(a.api.saveCheckpoint(payload(), state()), true);
  a.api.saveCheckpoint(payload({ userId: 2 }), state({ index: 5 }));
  const b = browser({ disk: a.disk });
  assert.deepEqual(plain(b.api.loadCheckpoint(payload())), state());
  assert.deepEqual(plain(b.api.getResume(1, 0, 'plan-a')), { day: 1, url: '/day/1' });
  assert.equal(b.api.getResume(1, 30, 'plan-a'), null);
  assert.equal(b.api.loadCheckpoint(payload()), null);
  assert.equal(b.api.loadCheckpoint(payload({ userId: 2 })).index, 5);
  b.api.saveCheckpoint(payload(), state());
  assert.equal(b.api.loadCheckpoint(payload({ resumeKey: 'plan-b:0:1', programKey: 'plan-b' })), null);
});

test('persist before fetch, concurrent send deduplication, durable acknowledgment and matching checkpoint removal', async () => {
  let resolveFetch, calls = 0;
  const b = browser({ fetch: async (_url, options) => {
    calls++;
    assert.equal(options.keepalive, true);
    assert.equal(options.credentials, 'same-origin');
    assert.equal(options.headers['Content-Type'], 'application/json');
    assert.equal([...b.disk.keys()].filter(key => key.includes(':pending:')).length, 1);
    return new Promise(resolve => { resolveFetch = resolve; });
  } });
  const notifications = [];
  b.api.subscribe(e => notifications.push(e));
  await b.api.start(1);
  b.api.saveCheckpoint(payload(), state());
  const first = b.api.sendCompletion(event());
  const second = b.api.sendCompletion(event());
  assert.equal(first, second);
  assert.equal(calls, 0, 'queue must persist synchronously before fetch runs');
  await Promise.resolve();
  assert.equal(calls, 1);
  resolveFetch(response(event()));
  assert.equal((await first).saved, true);
  assert.equal(b.api.pending(1).length, 0);
  assert.equal(b.api.loadCheckpoint(payload()), null);
  assert.deepEqual(notifications.map(e => e.status), ['saving', 'saved']);
  const reloaded = browser({ disk: b.disk });
  assert.equal(reloaded.api.ack(event().eventId, 1).position, 1);
  assert.equal(reloaded.api.ack(event().eventId, 2), null);
});

test('network failure remains queued across reload and retries preserve original date and position', async () => {
  const offline = browser({ fetch: async () => { throw Error('offline'); } });
  await offline.api.start(1);
  await assert.rejects(offline.api.sendCompletion(event()), /offline/);
  assert.equal(offline.api.pending(1).length, 1);
  assert.equal(offline.api.ack(event().eventId, 1), null);
  offline.api.queueCompletion(event({ completedAt: '2026-10-26T01:00:00Z', completedOn: '2026-10-26', position: 1 }));
  let sent;
  const online = browser({ disk: offline.disk, fetch: async (_url, options) => {
    sent = JSON.parse(options.body);
    return response(sent);
  } });
  await online.api.start(1);
  assert.deepEqual(sent, event());
  assert.equal(online.api.pending(1).length, 0);
  assert.equal(online.api.ack(event().eventId, 1).saved, true);
});

test('outbox sends only the active account and keeps other accounts queued', async () => {
  const sent = [];
  const b = browser({ fetch: async (_url, options) => {
    const e = JSON.parse(options.body);
    sent.push(e.userId);
    return response(e);
  } });
  b.api.queueCompletion(event());
  b.api.queueCompletion(event({ userId: 2 }));
  await b.api.start(2);
  await b.api.flush(1);
  await assert.rejects(b.api.sendCompletion(event()), /account_mismatch/);
  assert.deepEqual(sent, [2]);
  assert.equal(b.api.pending(1).length, 1);
  assert.equal(b.api.pending(2).length, 0);
  await b.api.start(1);
  assert.deepEqual(sent, [2, 1]);
});

test('storage failures keep in-memory state and report that pending completion is not durable', async () => {
  const b = browser({ storageFails: true, fetch: async () => { throw Error('offline'); } });
  const notifications = [];
  b.api.subscribe(e => notifications.push(e));
  await b.api.start(1);
  assert.equal(b.api.saveCheckpoint(payload(), state()), false);
  assert.deepEqual(plain(b.api.loadCheckpoint(payload())), state());
  assert.equal(b.api.queueCompletion(event()).durable, false);
  await assert.rejects(b.api.sendCompletion(event()), /offline/);
  assert.equal(b.api.pending(1).length, 1);
  assert.equal(notifications.at(-1).status, 'pending');
  assert.equal(notifications.at(-1).durable, false);
  assert.match(b.api.newEventId(), /^[a-f0-9]{32}$/);
});

test('corrupt storage and invalid server acknowledgment never discard pending completion', async () => {
  const disk = new Map([
    ['morning.workout.v1:1:checkpoint', '{broken'],
    ['morning.workout.v1:1:pending:bad', '{broken'],
    ['morning.workout.v1:1:ack:' + event().eventId, '{"saved":true}'],
  ]);
  const b = browser({ disk, fetch: async () => response(event(), { eventId: 'wrong' }) });
  await b.api.start(1);
  assert.equal(b.api.loadCheckpoint(payload()), null);
  assert.equal(b.api.pending(1).length, 0);
  await assert.rejects(b.api.sendCompletion(event()), /invalid_receipt/);
  assert.equal(b.api.pending(1).length, 1);
  assert.equal(b.api.ack(event().eventId, 1), null);
});

test('acknowledging an old workout preserves a newer workout checkpoint', async () => {
  const b = browser();
  await b.api.start(1);
  b.api.saveCheckpoint(payload(), state({ eventId: 'b'.repeat(32) }));
  await b.api.sendCompletion(event());
  assert.equal(b.api.loadCheckpoint(payload()).eventId, 'b'.repeat(32));
  // Startup may flush before the player rebuilds its finished checkpoint.
  b.api.saveCheckpoint(payload(), state());
  await b.api.sendCompletion(event());
  assert.equal(b.api.loadCheckpoint(payload()), null);
  b.api.saveCheckpoint(payload(), state({ eventId: 'b'.repeat(32) }));
  await b.api.sendCompletion(event());
  assert.equal(b.api.loadCheckpoint(payload()).eventId, 'b'.repeat(32));
});

test('singleton lifecycle listeners retry current account and script reinitialization does not duplicate them', async () => {
  let calls = 0;
  const b = browser({ fetch: async (_url, options) => {
    calls++;
    return response(JSON.parse(options.body));
  } });
  const original = b.api;
  b.reloadScript();
  assert.equal(b.api, original);
  assert.equal(b.events.get('online').length, 1);
  assert.equal(b.events.get('pageshow').length, 1);
  await b.api.start(1);
  await b.api.start(1);
  assert.equal([...b.timers.values()].filter(timer => timer.delay === 30000).length, 1);
  for (const lifecycle of ['online', 'pageshow', 'visibilitychange']) {
    const e = event({ eventId: String(calls + 1).repeat(32) });
    b.api.queueCompletion(e);
    b.emit(lifecycle);
    await b.api.flush(1);
    assert.equal(b.api.ack(e.eventId, 1).saved, true);
  }
  assert.equal(calls, 3);
});

test('subscriber-triggered retry shares the current request and another tab acknowledgment notifies active account', async () => {
  let calls = 0, reentrant;
  const b = browser({ fetch: async (_url, options) => { calls++; return response(JSON.parse(options.body)); } });
  await b.api.start(1);
  const unsubscribe = b.api.subscribe(e => {
    if (e.status === 'saving') reentrant = b.api.sendCompletion(event());
  });
  const first = b.api.sendCompletion(event());
  assert.equal(reentrant, first);
  await first;
  assert.equal(calls, 1);
  unsubscribe();
  const notices = [];
  b.api.subscribe(e => notices.push(e));
  const receipt = { saved: true, eventId: 'b'.repeat(32), position: 2, streak: 4 };
  const storageKey = 'morning.workout.v1:1:ack:' + receipt.eventId;
  b.disk.set(storageKey, JSON.stringify(receipt));
  b.emit('storage', { key: storageKey });
  assert.equal(notices.at(-1).status, 'saved');
  assert.equal(notices.at(-1).receipt.position, 2);
  b.emit('storage', { key: storageKey.replace(':1:', ':2:') });
  assert.equal(notices.length, 1);
});

const checkpointDisk = (savedState = state({ phase: 'complete', completion: event() })) => new Map([
  ['morning.workout.v1:1:checkpoint', JSON.stringify({ identity: payload(), state: savedState })],
]);

test('finished checkpoint is recovered before incompatible player state is discarded or replaced', async t => {
  for (const [name, invalidate] of [
    ['changed program', api => api.loadCheckpoint(payload({ programKey: 'plan-b', resumeKey: 'plan-b:0:1' }))],
    ['advanced position', api => api.getResume(1, 1, 'plan-a')],
    ['cleared checkpoint', api => api.clearCheckpoint(payload())],
    ['new workout checkpoint', api => api.saveCheckpoint(payload({ day: 2 }), state({ eventId: 'b'.repeat(32) }))],
  ]) await t.test(name, async () => {
    // Legacy state: the app stopped after saving its complete checkpoint but
    // before creating a pending event. No outbox entry exists at startup.
    const b = browser({ disk: checkpointDisk() });
    invalidate(b.api);
    assert.deepEqual(plain(b.api.pending(1)), [event()]);
    let sent;
    const reloaded = browser({ disk: b.disk, fetch: async (_url, options) => {
      sent = JSON.parse(options.body);
      return response(sent);
    } });
    await reloaded.api.start(1);
    assert.deepEqual(sent, event(), 'recovery retains the original completion date and position');
    assert.equal(reloaded.api.pending(1).length, 0);
  });
});

test('startup flush recovers a completed checkpoint without requiring the workout page to reopen', async () => {
  const sent = [];
  const b = browser({ disk: checkpointDisk(), fetch: async (_url, options) => {
    const e = JSON.parse(options.body);
    sent.push(e);
    return response(e);
  } });
  await b.api.start(2);
  assert.deepEqual(sent, [], 'another signed-in account cannot send the completion');
  await b.api.start(1);
  assert.deepEqual(sent, [event()]);
  assert.equal(b.api.loadCheckpoint(payload()), null);
  assert.equal(b.api.ack(event().eventId, 1).saved, true);
});

test('saving a completed checkpoint also persists its completion before any UI or network work', () => {
  const b = browser();
  assert.equal(b.api.saveCheckpoint(payload(), state({ phase: 'complete', completion: event() })), true);
  assert.deepEqual(plain(browser({ disk: b.disk }).api.pending(1)), [event()]);
});

test('checkpoint recovery keeps an existing immutable pending event and respects acknowledgment', async () => {
  const disk = checkpointDisk(state({ phase: 'complete', completion: event({
    completedAt: '2026-10-26T01:00:00Z', completedOn: '2026-10-26',
  }) }));
  disk.set('morning.workout.v1:1:pending:' + event().eventId, JSON.stringify(event()));
  let sent;
  const b = browser({ disk, fetch: async (_url, options) => {
    sent = JSON.parse(options.body);
    return response(sent);
  } });
  await b.api.start(1);
  assert.deepEqual(sent, event());

  const acknowledgedDisk = checkpointDisk();
  acknowledgedDisk.set('morning.workout.v1:1:ack:' + event().eventId,
    JSON.stringify({ saved: true, eventId: event().eventId, position: 1, streak: 4 }));
  let calls = 0;
  const acknowledged = browser({ disk: acknowledgedDisk, fetch: async () => { calls++; throw Error('unexpected send'); } });
  await acknowledged.api.start(1);
  acknowledged.api.getResume(1, 1, 'plan-a');
  assert.equal(calls, 0);
  assert.equal(acknowledged.api.pending(1).length, 0);
});

test('recovery never invents a finish from incomplete or mismatched checkpoint data', async t => {
  const cases = [
    ['unfinished', state({ phase: 'exercise', remainingMs: 0, completion: event() })],
    ['missing completion', state({ phase: 'complete' })],
    ['different event', state({ phase: 'complete', completion: event({ eventId: 'b'.repeat(32) }) })],
    ['different account', state({ phase: 'complete', completion: event({ userId: 2 }) })],
    ['different day', state({ phase: 'complete', completion: event({ day: 2 }) })],
    ['different position', state({ phase: 'complete', completion: event({ position: 1 }) })],
    ['invalid timestamp', state({ phase: 'complete', completion: event({ completedAt: 'invalid' }) })],
    ['invalid date', state({ phase: 'complete', completion: event({ completedOn: '2026-02-30' }) })],
  ];
  for (const [name, snapshot] of cases) await t.test(name, async () => {
    let calls = 0;
    const b = browser({ disk: checkpointDisk(snapshot), fetch: async () => { calls++; throw Error('unexpected send'); } });
    await b.api.start(1);
    b.api.getResume(1, 1, 'plan-a');
    assert.equal(b.api.pending(1).length, 0);
    assert.equal(calls, 0);
  });
});
