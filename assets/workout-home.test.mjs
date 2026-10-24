import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import vm from 'node:vm';

const source = readFileSync(new URL('./workout-home.js', import.meta.url), 'utf8');
const sharedSource = readFileSync(new URL('./workout-state.js', import.meta.url), 'utf8');
const initialNow = Date.parse('2026-10-24T12:00:00Z');
const receipt = (event, fields = {}) => ({ saved: true, eventId: event.eventId, position: 0, streak: 2, ...fields });
const response = value => ({ ok: true, status: 200, json: async () => value });
const event = (letter = 'a', fields = {}) => ({ eventId: letter.repeat(32), userId: 1, day: 1, position: 0,
  completedAt: '2026-10-24T12:00:00Z', completedOn: '2026-10-24', ...fields });
const payload = (fields = {}) => ({ userId: 1, position: 0, day: 1, programKey: 'plan', resumeKey: 'plan:0:1', ...fields });
const settle = async () => { for (let i = 0; i < 12; i++) await Promise.resolve(); };

function target(extra = {}) {
  const listeners = new Map();
  return Object.assign(extra, {
    addEventListener(name, fn) { if (!listeners.has(name)) listeners.set(name, new Set()); listeners.get(name).add(fn); },
    removeEventListener(name, fn) { listeners.get(name)?.delete(fn); },
    emit(name, event = {}) { [...(listeners.get(name) || [])].forEach(fn => fn(event)); },
  });
}

function node(attributes = {}) {
  const classes = new Set(['hidden']);
  return { textContent: '', href: '/', getAttribute: name => attributes[name],
    classList: {
      contains: name => classes.has(name),
      add: name => classes.add(name),
      remove: name => classes.delete(name),
      toggle(name, force) { const on = force ?? !classes.has(name); if (on) classes.add(name); else classes.delete(name); return on; },
    },
  };
}

function dashboard({ config = {}, fetch = async (_url, options) => response(receipt(JSON.parse(options.body))), prepare,
  visible = true } = {}) {
  let now = initialNow, nextTimer = 1, reloads = 0;
  const timers = new Map(), disk = new Map(), cleanups = [];
  const link = node(), status = node({ 'data-pending': 'Pending', 'data-saved': 'Saved' });
  const document = target({ visibilityState: visible ? 'visible' : 'hidden',
    querySelector: selector => selector === '.resume-workout' ? link : selector === '.home-save-status' ? status : null });
  const window = target({ fetch, navigator: { onLine: true }, location: { reload() { reloads++; } },
    localStorage: {
      get length() { return disk.size; }, key: i => [...disk.keys()][i],
      getItem: key => disk.get(key) ?? null, setItem: (key, value) => disk.set(key, value), removeItem: key => disk.delete(key),
    },
  });
  class ClockDate extends Date {
    constructor(...args) { super(...(args.length ? args : [now])); }
    static now() { return now; }
  }
  function schedule(fn, delay, interval = false) {
    const id = nextTimer++;
    timers.set(id, { fn, at: now + delay, interval: interval ? delay : 0 });
    return id;
  }
  const values = typeof config === 'string' ? config : { userId: 1, position: 0, programKey: 'plan', streak: 2, restartAt: '', ...config };
  const context = vm.createContext({ window, document, Date: ClockDate, AbortController,
    setTimeout: (fn, ms) => schedule(fn, ms), clearTimeout: id => timers.delete(id),
    setInterval: (fn, ms) => schedule(fn, ms, true), clearInterval: id => timers.delete(id),
    $data: () => values, $sys: { clean: fn => cleanups.push(fn) },
  });
  vm.runInContext(sharedSource, context);
  const state = window.MorningWorkoutState;
  prepare?.(state);
  vm.runInContext(source, context);
  return { state, link, status, document, window, disk, get reloads() { return reloads; },
    cleanup() { cleanups.forEach(fn => fn()); },
    setNow(value) { now = value; },
    advance(ms) {
      const end = now + ms;
      for (;;) {
        const due = [...timers].filter(([, t]) => t.at <= end).sort((a, b) => a[1].at - b[1].at)[0];
        if (!due) break;
        const [id, timer] = due;
        now = timer.at;
        if (timer.interval) timer.at += timer.interval; else timers.delete(id);
        timer.fn();
      }
      now = end;
    },
  };
}

test('home keeps Pending until every queued completion is acknowledged', async () => {
  const first = event('a'), second = event('b');
  let allowFirst = false, resolveSecond;
  const home = dashboard({ prepare(state) { state.queueCompletion(first); state.queueCompletion(second); },
    fetch: async (_url, options) => {
      const sent = JSON.parse(options.body);
      if (sent.eventId === first.eventId) {
        if (!allowFirst) throw Error('offline');
        return response(receipt(sent));
      }
      return new Promise(resolve => { resolveSecond = resolve; });
    },
  });
  await settle();
  assert.equal(home.status.textContent, 'Pending');
  resolveSecond(response(receipt(second)));
  await settle();
  assert.equal(home.state.pending(1).length, 1);
  assert.equal(home.status.textContent, 'Pending', 'one saved event cannot hide another failed event');
  assert.equal(home.reloads, 0, 'a replay without progress changes does not reload');
  allowFirst = true;
  await home.state.flush(1);
  assert.equal(home.status.textContent, 'Saved');
  assert.equal(home.status.classList.contains('hidden'), false);
});

test('delayed home acknowledgment removes matching resume and refreshes updated calendar once', async () => {
  const completed = event();
  let acknowledge;
  const home = dashboard({ config: { streak: 1 }, prepare(state) {
    state.saveCheckpoint(payload(), { eventId: completed.eventId, phase: 'complete' });
    state.queueCompletion(completed);
  }, fetch: () => new Promise(resolve => { acknowledge = resolve; }) });
  await settle();
  assert.equal(home.status.textContent, 'Pending');
  assert.equal(home.link.href, '/day/1');
  assert.equal(home.link.classList.contains('hidden'), false);
  assert.equal(home.reloads, 0);
  acknowledge(response(receipt(completed, { position: 1, streak: 2 })));
  await settle();
  assert.equal(home.status.textContent, 'Saved');
  assert.equal(home.link.classList.contains('hidden'), true);
  home.advance(249);
  assert.equal(home.reloads, 0);
  home.advance(1);
  assert.equal(home.reloads, 1);
  await home.state.sendCompletion(completed);
  home.advance(500);
  assert.equal(home.reloads, 1);
});

test('resume uses the checkpoint day and current account, cycle position and program', () => {
  const home = dashboard({ config: { position: 11 }, prepare(state) {
    state.saveCheckpoint(payload({ position: 11, day: 5, resumeKey: 'plan:11:5' }), { eventId: 'x'.repeat(32) });
  } });
  assert.equal(home.link.href, '/day/5');
  assert.equal(home.link.classList.contains('hidden'), false);
  assert.equal(home.status.classList.contains('hidden'), true);
  for (const fields of [{ position: 30 }, { programKey: 'new-plan' }, { userId: 2 }]) {
    const stale = dashboard({ config: fields, prepare(state) {
      state.saveCheckpoint(payload(), { eventId: 'x'.repeat(32) });
    } });
    assert.equal(stale.link.classList.contains('hidden'), true);
  }
});

test('another account completion never changes home status or calendar', async () => {
  const other = event('b', { userId: 2 });
  const home = dashboard({ prepare(state) { state.queueCompletion(other); } });
  await settle();
  assert.equal(home.state.pending(2).length, 1);
  assert.equal(home.status.classList.contains('hidden'), true);
  await home.state.start(2);
  home.advance(500);
  assert.equal(home.status.classList.contains('hidden'), true);
  assert.equal(home.reloads, 0);
});

test('returning on the next local date refreshes once and never reloads while hidden', () => {
  const home = dashboard();
  home.document.visibilityState = 'hidden';
  home.setNow(initialNow + 24 * 60 * 60 * 1000);
  home.document.emit('visibilitychange');
  home.window.emit('pageshow');
  assert.equal(home.reloads, 0);
  home.document.visibilityState = 'visible';
  home.document.emit('visibilitychange');
  home.window.emit('pageshow');
  assert.equal(home.reloads, 1);
});

test('restart deadline refreshes a visible home even before the local date changes', () => {
  const home = dashboard({ config: { restartAt: new Date(initialNow + 60000).toISOString() } });
  home.advance(59999);
  assert.equal(home.reloads, 0);
  home.advance(1);
  assert.equal(home.reloads, 1);
  home.window.emit('pageshow');
  home.advance(60000);
  assert.equal(home.reloads, 1);
});

test('already due restart refreshes at initialization or when hidden home becomes visible', () => {
  const config = { restartAt: new Date(initialNow).toISOString() };
  assert.equal(dashboard({ config }).reloads, 1);
  const hidden = dashboard({ config, visible: false });
  assert.equal(hidden.reloads, 0);
  hidden.document.visibilityState = 'visible';
  hidden.document.emit('visibilitychange');
  assert.equal(hidden.reloads, 1);
});

test('home cleanup cancels delayed calendar refresh and lifecycle reloads', async () => {
  const home = dashboard({ prepare(state) { state.queueCompletion(event()); },
    fetch: async () => response(receipt(event(), { position: 1 })) });
  await settle();
  assert.equal(home.status.textContent, 'Saved');
  home.cleanup();
  home.advance(500);
  home.setNow(initialNow + 24 * 60 * 60 * 1000);
  home.window.emit('pageshow');
  home.document.emit('visibilitychange');
  assert.equal(home.reloads, 0);
});

test('invalid serialized home configuration does not break the page', () => {
  assert.doesNotThrow(() => dashboard({ config: '{broken' }));
});
