import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import vm from 'node:vm';

// Drive the shipped player with a small DOM and a deterministic clock: these
// tests check actual transitions, Done/Skip buttons and visible side labels.
function memoryState(options = {}) {
  const data = { checkpoint: null, events: new Map(), receipts: new Map(), subscribers: new Set(), nextId: 0 };
  let userId;
  const copy = value => value == null ? value : JSON.parse(JSON.stringify(value));
  const notify = event => data.subscribers.forEach(fn => fn(event));
  const api = {
    start(id) { userId = id; },
    newEventId() { return (++data.nextId).toString(16).padStart(32, '0'); },
    saveCheckpoint(payload, state) { data.checkpoint = { payload: copy(payload), state: copy(state) }; return !options.storageUnavailable; },
    loadCheckpoint(payload) {
      const record = data.checkpoint;
      if (!record) return null;
      if (!['userId', 'position', 'day', 'resumeKey'].every(k => record.payload[k] === payload[k])) {
        data.checkpoint = null; return null;
      }
      return copy(record.state);
    },
    clearCheckpoint() { data.checkpoint = null; },
    queueCompletion(event) { if (!data.events.has(event.eventId)) data.events.set(event.eventId, copy(event)); return { durable: !options.storageUnavailable }; },
    sendCompletion(event) {
      api.queueCompletion(event);
      notify({ ...event, status: 'saving' });
      if (options.rejectCompletion) {
        notify({ ...event, status: 'pending', durable: !options.storageUnavailable });
        return Promise.reject(new Error('Offline'));
      }
      return Promise.resolve().then(() => {
        const receipt = { saved: true, eventId: event.eventId, streak: 1, position: event.position + 1 };
        data.receipts.set(event.eventId, receipt);
        if (data.checkpoint?.state.eventId === event.eventId) data.checkpoint = null;
        notify({ ...event, status: 'saved', receipt });
        return receipt;
      });
    },
    ack(eventId, id) { return id === userId ? data.receipts.get(eventId) : null; },
    subscribe(fn) { data.subscribers.add(fn); return () => data.subscribers.delete(fn); },
  };
  return { api, data, options };
}
function player(items, extra = {}, options = {}) {
  const nodes = new Map();
  const timers = new Map();
  const spoken = [];
  const utterances = [];
  let speechResumes = 0;
  const navigations = [];
  const images = [];
  const persistence = options.persistence || memoryState(options);
  let now = 0, nextTimer = 1, cleanup = () => {};
  function eventTarget(target = {}) {
    const events = new Map();
    return Object.assign(target, {
      addEventListener(name, fn) { if (!events.has(name)) events.set(name, new Set()); events.get(name).add(fn); },
      removeEventListener(name, fn) { events.get(name)?.delete(fn); },
      dispatch(name, event = {}) { [...(events.get(name) || [])].forEach(fn => fn(event)); },
    });
  }
  function element() {
    const classes = new Set();
    const node = eventTarget({
      textContent: '', style: {}, children: [], disabled: false,
      classList: {
        add(...xs) { xs.forEach(x => classes.add(x)); },
        remove(...xs) { xs.forEach(x => classes.delete(x)); },
        contains(x) { return classes.has(x); },
        toggle(x, yes) { if (yes ?? !classes.has(x)) classes.add(x); else classes.delete(x); },
      },
      appendChild(child) { this.children.push(child); child.parentElement = this; },
      setAttribute(name, value) { this[name] = value; },
      removeAttribute(name) { delete this[name]; },
      click() { if (!this.disabled) this.dispatch('click', { preventDefault() {} }); },
    });
    Object.defineProperty(node, 'className', {
      get: () => [...classes].join(' '),
      set: value => { classes.clear(); value.split(/\s+/).forEach(cls => { classes.add(cls); if (cls && !nodes.has('.' + cls)) nodes.set('.' + cls, node); }); },
    });
    return node;
  }
  const root = element();
  const get = selector => {
    if (!nodes.has(selector)) {
      const node = element();
      node.parentElement = element();
      if (['.info-overlay', '.done-overlay', '.rest-overlay'].includes(selector)) node.classList.add('hidden');
      nodes.set(selector, node);
    }
    return nodes.get(selector);
  };
  root.querySelector = get;
  const document = eventTarget({ querySelector: () => root, createElement: element, visibilityState: 'visible' });
  const schedule = (fn, delay, interval) => {
    const id = nextTimer++;
    timers.set(id, { fn, due: now + delay, interval });
    return id;
  };
  const window = eventTarget({
    matchMedia: () => ({ matches: !!options.reducedMotion }),
    MorningWorkoutState: persistence.api,
    speechSynthesis: { paused: !!options.speechPaused, getVoices: () => [], cancel() {},
      resume() { this.paused = false; speechResumes++; },
      speak(utterance) { utterances.push(utterance); spoken.push(utterance.text); } },
  });
  if (options.noSpeech) delete window.speechSynthesis;
  const base = Date.parse('2026-10-24T08:00:00Z');
  class ClockDate extends Date {
    constructor(...args) { super(...(args.length ? args : [base + now])); }
    static now() { return base + now; }
  }
  const context = {
    document, navigator: options.navigator || {}, window, Date: ClockDate,
    Math: options.random ? Object.assign(Object.create(Math), { random: options.random }) : Math,
    location: { assign(url) { navigations.push(url); } },
    Image: function () { images.push(this); },
    SpeechSynthesisUtterance: function (text) { this.text = text; },
    $data: () => ({ userId: 1, position: 0, day: 1, resumeKey: 'plan:0:1', programKey: 'plan', items, voiceMode: 'off',
      cues: { switch_side: 'Change sides', go: 'Go' }, labels: { side: 'Side', per_side: 'per side', resume_title: 'Paused', resume: 'Continue',
        saving: 'Saving', saved: 'Saved', save_pending: 'Waiting for connection', save_unavailable: 'Keep this page open', save_retry: 'Retry' }, ...extra }),
    $hook: () => { throw new Error('Completion must not depend on Doors hooks'); },
    $sys: { clean(fn) { cleanup = fn; } },
    setInterval: (fn, delay) => schedule(fn, delay, delay), clearInterval: id => timers.delete(id),
    setTimeout: (fn, delay) => schedule(fn, delay, 0), clearTimeout: id => timers.delete(id),
  };
  if (options.realStore) {
    const values = options.storageValues || new Map();
    delete window.MorningWorkoutState;
    window.localStorage = {
      get length() { return values.size; }, key: index => [...values.keys()][index] || null,
      getItem: key => values.get(key) ?? null, setItem: (key, value) => values.set(key, value), removeItem: key => values.delete(key),
    };
    window.fetch = options.fetch || (() => Promise.reject(new Error('Offline')));
    vm.runInNewContext(readFileSync(new URL('./workout-state.js', import.meta.url), 'utf8'), context);
  }
  vm.runInNewContext(readFileSync(new URL('./player.js', import.meta.url), 'utf8'), context);
  function advance(seconds) {
    const end = now + seconds * 1000;
    while (true) {
      const next = [...timers].filter(([, t]) => t.due <= end).sort((a, b) => a[1].due - b[1].due)[0];
      if (!next) break;
      const [id, timer] = next;
      now = timer.due;
      if (timer.interval) timer.due += timer.interval; else timers.delete(id);
      timer.fn();
    }
    now = end;
  }
  return {
    get, advance, spoken, utterances, navigations, images, persistence, window, document, stateStore: window.MorningWorkoutState, cleanup: () => cleanup(),
    get speechResumes() { return speechResumes; },
    get completed() { return [...persistence.data.events.values()]; },
    hide() { document.visibilityState = 'hidden'; document.dispatch('visibilitychange'); window.dispatch('pagehide'); },
    show() { document.visibilityState = 'visible'; document.dispatch('visibilitychange'); window.dispatch('pageshow'); window.dispatch('focus'); },
  };
}
const item = (fields = {}) => ({ id: 'a', name: 'Exercise', unit: 'reps', value: 6, slot: 'main', round: 1, rest: 0, ...fields });
const hidden = node => node.classList.contains('hidden');

const reactionItem = (fields = {}) => item({ unit: 'seconds', value: 30,
  reaction: { directions: ['left', 'right'] }, correct: ['Keep moving'], wrong: ['Do not cross feet'], ...fields });
const reactionCues = { go: 'Go', halfway: 'Half', reaction_left: 'Left', reaction_right: 'Right' };
const directionsSaid = p => p.spoken.filter(word => word === 'Left' || word === 'Right');

test('reaction cues share continuous work time and balance randomized directions without fixed alternation', () => {
  const samples = [0, .99, .99, 0, 0, 0, 0];
  const p = player([reactionItem()], { voiceMode: 'normal', cues: reactionCues }, { random: () => samples.shift() ?? 0 });
  p.advance(3); // ready complete
  assert.ok(hidden(p.get('.reaction-cue')));
  assert.ok(!hidden(p.get('.ex-visual')), 'reaction still works before illustrations are published');
  p.advance(2);
  assert.equal(directionsSaid(p).length, 0);
  p.advance(1);
  assert.equal(p.get('.reaction-word').textContent, 'Left');
  assert.equal(p.get('.reaction-arrow').textContent, '←');
  assert.ok(!hidden(p.get('.reaction-cue')));
  assert.equal(p.get('.ex-value').textContent, '27″');
  p.advance(9);
  assert.deepEqual(directionsSaid(p).slice(0, 4), ['Left', 'Left', 'Right', 'Right']);
  assert.equal(p.get('.ex-value').textContent, '18″');
  assert.ok(!p.spoken.includes('Keep moving') && !p.spoken.includes('Do not cross feet') && !p.spoken.includes('Half'));
  p.advance(15);
  assert.equal(p.get('.ex-value').textContent, '3″');
  assert.ok(hidden(p.get('.reaction-cue')));
  const count = directionsSaid(p).length;
  p.advance(3);
  assert.equal(directionsSaid(p).length, count);
  assert.deepEqual(p.spoken.slice(-3), ['3', '2', '1']);
  assert.equal(p.completed.length, 1);
});

test('reaction pacing varies from three to five seconds and muted cues remain visible', () => {
  for (const voiceMode of ['off', 'min', 'detailed']) {
    const p = player([reactionItem()], { voiceMode, cues: reactionCues }, { random: () => .99 });
    p.advance(6); // 3 ready + 3 work
    assert.equal(p.get('.reaction-word').textContent, 'Right');
    assert.ok(!hidden(p.get('.reaction-cue')));
    p.advance(2);
    assert.ok(hidden(p.get('.reaction-cue')));
    p.advance(2);
    assert.ok(hidden(p.get('.reaction-cue')), 'five-second interval has not expired');
    p.advance(1);
    assert.equal(p.get('.reaction-word').textContent, 'Left');
    assert.ok(!hidden(p.get('.reaction-cue')));
    if (voiceMode === 'off') assert.equal(p.spoken.length, 0);
    else assert.deepEqual(directionsSaid(p), ['Right', 'Left']);
  }
});

test('reaction cues pause through app backgrounding and technique cards without overdue bursts', () => {
  const p = player([reactionItem()], { voiceMode: 'normal', cues: reactionCues }, { random: () => 0 });
  p.advance(6);
  assert.equal(directionsSaid(p).length, 1);
  p.hide();
  assert.ok(hidden(p.get('.reaction-cue')));
  p.advance(60);
  p.show();
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, 27000);
  assert.equal(directionsSaid(p).length, 1);
  p.get('.resume-continue').click();
  p.advance(2);
  assert.equal(directionsSaid(p).length, 1);
  p.advance(1);
  assert.equal(directionsSaid(p).length, 2);
  p.get('.ex-visual').click();
  assert.ok(hidden(p.get('.reaction-cue')));
  p.advance(30);
  p.get('.info-close').click();
  p.advance(2);
  assert.equal(directionsSaid(p).length, 2);
  p.advance(1);
  assert.equal(directionsSaid(p).length, 3);
  assert.equal(p.get('.ex-value').textContent, '21″');
  p.cleanup();
  assert.ok(hidden(p.get('.reaction-cue')));
  p.advance(60);
  assert.equal(directionsSaid(p).length, 3);
});

test('reaction cues stop at rest, skip, back and replacement; ordinary exercises keep their narration', () => {
  const p = player([reactionItem({ rest: 10 }), item({ id: 'next', unit: 'seconds', value: 30, correct: ['Steady'] })],
    { voiceMode: 'normal', cues: reactionCues }, { random: () => 0 });
  p.advance(6);
  p.get('.ctl-skip').click();
  assert.ok(hidden(p.get('.reaction-cue')));
  p.advance(10);
  p.advance(3);
  assert.equal(directionsSaid(p).length, 1);
  assert.ok(p.spoken.includes('Steady'));
  p.get('.ctl-prev').click();
  p.advance(3);
  assert.ok(hidden(p.get('.reaction-cue')));
  p.advance(3);
  assert.equal(directionsSaid(p).length, 2);

  const replacement = item({ id: 'calm', unit: 'seconds', value: 10 });
  const swapped = player([reactionItem({ repl: replacement })], { voiceMode: 'normal', cues: reactionCues }, { random: () => 0 });
  swapped.advance(6);
  swapped.get('.ex-visual').click();
  swapped.get('.info-replace').click();
  assert.ok(hidden(swapped.get('.reaction-cue')));
  swapped.advance(5);
  assert.equal(directionsSaid(swapped).length, 1);

  const intoReaction = player([item({ repl: reactionItem({ id: 'react' }) })], { voiceMode: 'off', cues: reactionCues }, { random: () => 0 });
  intoReaction.advance(3);
  intoReaction.get('.ex-visual').click();
  intoReaction.get('.info-replace').click();
  intoReaction.advance(3);
  assert.ok(!hidden(intoReaction.get('.reaction-cue')));
});

test('reloading a reaction drill preserves work remaining and waits for a fresh signal after Continue', () => {
  const persistence = memoryState();
  const first = player([reactionItem()], { voiceMode: 'normal', cues: reactionCues }, { persistence, random: () => 0 });
  first.advance(6);
  first.hide();
  first.cleanup();
  const resumed = player([reactionItem()], { voiceMode: 'normal', cues: reactionCues }, { persistence, random: () => 0 });
  assert.equal(resumed.get('.ex-value').textContent, '27″');
  assert.ok(hidden(resumed.get('.reaction-cue')));
  resumed.advance(60);
  assert.equal(directionsSaid(resumed).length, 0);
  resumed.get('.resume-continue').click();
  resumed.advance(2);
  assert.equal(directionsSaid(resumed).length, 0);
  resumed.advance(1);
  assert.equal(directionsSaid(resumed).length, 1);
  assert.equal(resumed.get('.ex-value').textContent, '24″');
});

test('reaction remains disabled for invalid, counted, and bilateral configurations', () => {
  for (const fields of [{ unit: 'reps' }, { perSide: true }, { reaction: { directions: ['left', 'left'] } }]) {
    const p = player([reactionItem(fields)], { voiceMode: 'normal', cues: reactionCues }, { random: () => 0 });
    p.advance(10);
    assert.equal(directionsSaid(p).length, 0);
    assert.ok(hidden(p.get('.reaction-cue')));
    assert.ok(p.spoken.includes('Keep moving'));
  }
});

test('reaction figure stays neutral while random signals run; technique card can play its demonstration', () => {
  const animation = { src: '/poses/react', frames: 3, kind: 'alternating' };
  const p = player([reactionItem({ animation })], { cues: reactionCues }, { random: () => 0 });
  p.advance(6);
  assert.equal(p.get('.ex-visual').style.backgroundPosition, '0% 50%');
  p.get('.ex-name').click();
  p.advance(1);
  assert.equal(p.get('.info-visual').style.backgroundPosition, '50% 50%');
  p.get('.info-close').click();
  p.advance(5);
  assert.equal(p.get('.ex-visual').style.backgroundPosition, '0% 50%');
});

test('exercise title opens technique before illustrations arrive and pauses work time', () => {
  const p = player([item({ unit: 'seconds', value: 30, desc: 'Small quick steps' })]);
  p.advance(5);
  assert.ok(hidden(p.get('.ex-visual')));
  p.get('.ex-name').click();
  assert.ok(!hidden(p.get('.info-overlay')));
  assert.equal(p.get('.info-desc').textContent, 'Small quick steps');
  p.advance(10);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, 28000);
  p.get('.info-close').click();
  p.advance(1);
  assert.equal(p.get('.ex-value').textContent, '27″');
});

test('speech failures are visible and narration resumes after returning without losing workout time', () => {
  const p = player([item({ unit: 'seconds', value: 30, correct: ['Stay steady'] })], { voiceMode: 'normal' });
  assert.ok(hidden(p.get('.voice-status')));
  p.advance(3);
  p.utterances.at(-1).onerror({ error: 'not-allowed' });
  assert.equal(p.get('.voice-status').textContent, 'voice_allow');
  p.hide();
  const remaining = p.persistence.data.checkpoint.state.remainingMs;
  p.advance(2);
  p.show();
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, remaining);
  p.get('.resume-continue').click();
  p.utterances.at(-1).onstart();
  assert.ok(hidden(p.get('.voice-status')));
  p.advance(3);
  assert.ok(p.spoken.includes('Stay steady'));
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, remaining - 3000);
});

test('speech reports actual missing engines and synthesis failures', () => {
  const missing = player([item()], { voiceMode: 'normal' }, { noSpeech: true });
  assert.equal(missing.get('.voice-status').textContent, 'voice_engine');
  const p = player([item()], { voiceMode: 'normal' });
  p.utterances.at(-1).onerror({ error: 'language-unavailable' });
  assert.equal(p.get('.voice-status').textContent, 'voice_engine');
  p.utterances.at(-1).onerror({ error: 'network' });
  assert.equal(p.get('.voice-status').textContent, 'voice_failed');
});

test('speech recovers a paused synthesizer and ignores canceled or stale utterances', () => {
  const p = player([item()], { voiceMode: 'normal' }, { speechPaused: true });
  assert.equal(p.speechResumes, 1);
  p.utterances.at(-1).onerror({ error: 'canceled' });
  assert.ok(hidden(p.get('.voice-status')));
  const old = p.utterances.at(-1);
  p.hide();
  old.onerror({ error: 'network' });
  assert.ok(hidden(p.get('.voice-status')));
  p.cleanup();
  old.onerror({ error: 'network' });
  p.advance(10);
  assert.ok(hidden(p.get('.voice-status')));
});

test('voice errors never cover exercise details or completion', () => {
  const p = player([item()], { voiceMode: 'normal' });
  p.advance(3);
  p.utterances.at(-1).onerror({ error: 'not-allowed' });
  assert.ok(!hidden(p.get('.voice-status')));
  p.get('.ex-visual').click();
  assert.ok(hidden(p.get('.voice-status')));
  p.get('.info-close').click();
  assert.ok(!hidden(p.get('.voice-status')));
  p.get('.ex-visual').click();
  p.hide();
  p.show();
  assert.ok(!hidden(p.get('.voice-status')));
  p.get('.resume-continue').click();
  p.get('.ctl-done').click();
  assert.ok(!hidden(p.get('.done-overlay')));
  assert.ok(hidden(p.get('.voice-status')));
});

test('unillustrated exercises clear the previous image until pictures are published', () => {
  const p = player([item({ svg: '/poses/old' }), item({ id: 'new' }), item({ svg: '/poses/new' })]);
  p.advance(3);
  assert.equal(p.get('.ex-svg').src, '/poses/old');
  p.get('.ctl-done').click();
  assert.ok(hidden(p.get('.ex-svg')));
  assert.equal(p.get('.ex-svg').src, undefined);
  assert.ok(hidden(p.get('.info-media')));
  p.get('.ctl-done').click();
  assert.ok(!hidden(p.get('.ex-svg')));
  assert.equal(p.get('.ex-svg').src, '/poses/new');
});

test('stale authored side pauses cannot split counted work or change the universal preparation', () => {
  const p = player([
    item({ perSide: true, sideRest: 90, rest: 15 }),
    item({ id: 'hold', unit: 'seconds', value: 5, perSide: true, sideRest: 90, rest: 20 }),
    item({ id: 'next', name: 'Next' }),
  ]);
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '× 6 per side');
  assert.doesNotMatch(p.get('.phase').textContent, /Side [12]\/2/);
  p.get('.ctl-done').click();
  assert.equal(p.get('.rest-count').textContent, 15);
  p.advance(15);
  assert.match(p.get('.phase').textContent, /Side 1\/2/);
  p.advance(5);
  assert.equal(p.persistence.data.checkpoint.state.phase, 'side-change');
  assert.equal(p.get('.ex-value').textContent, '3″');
  assert.ok(hidden(p.get('.rest-overlay')));
  p.advance(3);
  assert.match(p.get('.phase').textContent, /Side 2\/2/);
  p.advance(5);
  assert.equal(p.get('.rest-count').textContent, 20);
  p.advance(20);
  assert.equal(p.get('.ex-name').textContent, 'Next');
});

test('counted sides use one Done tap and timed sides get one preparation', () => {
  const p = player([item({ perSide: true }), item({ unit: 'seconds', value: 5, perSide: true })]);
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '× 6 per side');
  p.get('.ctl-done').click();
  p.advance(8);
  assert.match(p.get('.phase').textContent, /Side 2\/2/);
  assert.ok(hidden(p.get('.rest-overlay')));
  p.advance(5);
  assert.equal(p.completed.length, 1);
});

test('self-managed timed sides keep one screen and 40 seconds of work around three-second preparation', () => {
  const p = player([
    item({ perSide: true }),
    item({ id: 'hold', name: 'Hold', unit: 'seconds', value: 20, perSide: true }),
    item({ id: 'next', name: 'Next' }),
  ], { selfManagedSides: true, voiceMode: 'normal', cues: { switch_side: 'Change sides.', halfway: 'Halfway' } });
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '× 6 per side');
  p.get('.ctl-done').click();
  assert.equal(p.get('.ex-value').textContent, '40″');
  assert.doesNotMatch(p.get('.phase').textContent, /Side [12]\/2/);
  p.advance(19);
  assert.equal(p.get('.ex-value').textContent, '21″');
  p.advance(1);
  assert.equal(p.persistence.data.checkpoint.state.phase, 'side-change');
  assert.equal(p.get('.ex-name').textContent, 'Hold');
  assert.equal(p.get('.ex-value').textContent, '3″');
  assert.ok(hidden(p.get('.rest-overlay')));
  assert.equal(p.spoken.at(-1), 'Change sides. 3, 2, 1.');
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '20″');
  assert.doesNotMatch(p.get('.phase').textContent, /Side [12]\/2/);
  assert.equal(p.spoken.at(-1), 'Change sides. 3, 2, 1.'); // no Go cuts off the instruction
  p.advance(19);
  assert.equal(p.get('.ex-name').textContent, 'Hold');
  assert.equal(p.get('.ex-value').textContent, '1″');
  p.advance(1);
  assert.equal(p.get('.ex-name').textContent, 'Next');
  assert.equal(p.spoken.filter(x => x.includes('Change sides')).length, 1);
  assert.ok(!p.spoken.includes('Halfway'));
});

test('self-managed second side resumes its saved remainder without repeating the first side', () => {
  const items = [item({ unit: 'seconds', value: 20, perSide: true }), item({ name: 'Next' })];
  const first = player(items, { selfManagedSides: true });
  first.advance(32); // ready + first20 + preparation3 + second6
  first.hide(); first.cleanup();
  assert.equal(first.persistence.data.checkpoint.state.remainingMs, 14000);
  const second = player(items, { selfManagedSides: true }, { persistence: first.persistence });
  assert.equal(second.get('.ex-value').textContent, '14″');
  second.advance(30);
  assert.equal(second.get('.ex-value').textContent, '14″');
  second.get('.resume-continue').click();
  second.advance(14);
  assert.equal(second.get('.ex-name').textContent, 'Next');
});

test('a difficulty change preserves both running side durations and their label across reload', () => {
  const items = [item({ unit: 'seconds', value: 20, perSide: true, scale: true }), item({ name: 'Next', unit: 'seconds', value: 10, scale: true })];
  const first = player(items, { selfManagedSides: true });
  first.advance(8); first.get('.lvl-up').click(); first.hide(); first.cleanup();
  const second = player(items, { selfManagedSides: true }, { persistence: first.persistence });
  assert.equal(second.get('.ex-value').textContent, '35″');
  assert.match(second.get('.phase').textContent, /20″ per side/);
  second.get('.resume-continue').click(); second.advance(18);
  assert.equal(second.get('.ex-value').textContent, '20″');
  assert.match(second.get('.phase').textContent, /20″ per side/);
  second.advance(20);
  assert.equal(second.get('.ex-name').textContent, 'Next');
  assert.equal(second.get('.ex-value').textContent, '11″');
});

for (const selfManagedSides of [false, true]) {
  test(`side-change restores exact remaining preparation and targets side two (combined=${selfManagedSides})`, () => {
    const items = [item({ unit: 'seconds', value: 20, perSide: true }), item({ name: 'Next' })];
    const extra = { selfManagedSides, voiceMode: 'normal' };
    const first = player(items, extra);
    first.advance(24.4);
    first.hide(); first.cleanup();
    assert.equal(first.persistence.data.checkpoint.state.phase, 'side-change');
    assert.equal(first.persistence.data.checkpoint.state.remainingMs, 1600);
    assert.equal(first.persistence.data.checkpoint.state.sideDurationMs, 20000);
    const second = player(items, extra, { persistence: first.persistence });
    assert.equal(second.get('.ex-value').textContent, '2″');
    second.advance(30);
    second.get('.resume-continue').click();
    assert.equal(second.spoken.at(-1), '2, 1.');
    second.advance(2);
    assert.equal(second.get('.ex-value').textContent, '20″');
    assert.equal(second.persistence.data.checkpoint.state.side, 2);
    assert.equal(second.spoken.filter(x => x.includes('Change sides')).length, 0);
    second.advance(20);
    assert.equal(second.get('.ex-name').textContent, 'Next');
  });
}

test('side-change details, Skip and Back preserve the exercise and intended side', () => {
  const p = player([item({ name: 'Before' }), item({ name: 'Hold', unit: 'seconds', value: 20, perSide: true }), item({ name: 'Next' })], { selfManagedSides: true });
  p.advance(3); p.get('.ctl-done').click(); p.advance(21);
  p.get('.ex-visual').click(); p.advance(10); p.get('.info-close').click();
  assert.equal(p.get('.ex-value').textContent, '2″');
  p.get('.ctl-prev').click(); p.advance(3);
  assert.equal(p.get('.ex-name').textContent, 'Hold');
  assert.equal(p.get('.ex-value').textContent, '40″');
  p.advance(20); p.get('.ctl-skip').click();
  assert.equal(p.get('.ex-name').textContent, 'Hold');
  assert.equal(p.get('.ex-value').textContent, '20″');
  p.advance(20);
  assert.equal(p.get('.ex-name').textContent, 'Next');
});

test('phase shows current and total main rounds while single-round warmup and cooldown stay compact', () => {
  const p = player([
    item({ id: 'warm', slot: 'warmup' }),
    item({ id: 'a', round: 1 }), item({ id: 'b', round: 1 }),
    item({ id: 'a', round: 2 }), item({ id: 'b', round: 2 }),
    item({ id: 'cool', slot: 'cooldown' }),
  ], { labels: { round: 'Круг' } });
  p.advance(3);
  assert.match(p.get('.phase').textContent, /Круг 1$/);
  for (const expected of ['1/2', '1/2', '2/2', '2/2', '1']) {
    p.get('.ctl-done').click();
    assert.ok(p.get('.phase').textContent.endsWith('Круг ' + expected));
  }
});

test('fixed difficulty ignores saved level and all in-session level controls', () => {
  const p = player([item({ scale: true, value: 10 })], { fixedDifficulty: true, level: 3 });
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '× 10');
  assert.ok(hidden(p.get('.lvl-row')));
  p.get('.lvl-up').click();
  p.get('.lvl-down').click();
  assert.equal(p.get('.ex-value').textContent, '× 10');
});

test('voice keeps technique and mistake cues without a spoken introduction', () => {
  const correct = ['Колени направлены по линии стоп.', 'Пятки остаются на полу.', 'Двигайся плавно.'];
  const wrong = ['Не своди колени и не отрывай пятки.', 'Не задерживай дыхание.'];
  for (const voiceMode of ['off', 'min', 'normal', 'detailed']) {
    const p = player([item({ correct, wrong, vStart: 'Начали.' })], { voiceMode });
    p.advance(60);
    const expected = voiceMode === 'off' ? [] : ['3', '2', '1', 'Начали.'];
    if (voiceMode === 'normal') expected.push(...correct.slice(0, 2), ...wrong.slice(0, 1));
    if (voiceMode === 'detailed') expected.push(...correct, ...wrong);
    assert.deepEqual(p.spoken, expected, voiceMode);
    p.get('.ex-visual').click();
    assert.deepEqual(p.get('.info-wrong').children.map(node => node.textContent), wrong);
  }
});

test('exercise details show both warning and safety, without repeating identical text', () => {
  for (const [warning, safety, expected] of [
    ['Count total repetitions.', 'Replace an unsteady hop with a step.', 'Count total repetitions. Replace an unsteady hop with a step.'],
    [' Keep your heels down. ', 'Keep your heels down.', 'Keep your heels down.'],
    ['', 'Learn this with an adult first.', 'Learn this with an adult first.'],
  ]) {
    const p = player([item({ warning, safety })]);
    p.advance(3);
    p.get('.ex-visual').click();
    assert.ok(!hidden(p.get('.info-overlay')));
    assert.ok(hidden(p.get('.ex-svg')));
    assert.equal(p.get('.info-warn').textContent, expected);
    assert.ok(!hidden(p.get('.info-warn-sec')));
  }
});

const settle = async () => { for (let n = 0; n < 20; n++) await Promise.resolve(); };

test('hiding freezes a fractional timer, stops narration and requires an explicit Continue', () => {
  const p = player([item({ unit: 'seconds', value: 20, correct: ['Stay steady'] }), item({ name: 'Next' })], { voiceMode: 'normal' });
  p.advance(8.4); // ready 3s + work 5.4s
  p.hide();
  const atHide = p.persistence.data.checkpoint.state;
  assert.equal(atHide.phase, 'exercise');
  assert.equal(atHide.remainingMs, 14600);
  const spoken = p.spoken.length;
  p.advance(120);
  p.show();
  p.advance(30);
  assert.equal(p.spoken.length, spoken);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, 14600);
  assert.ok(!hidden(p.get('.resume-overlay')));
  p.get('.ctl-skip').click(); // covered controls cannot accidentally advance
  assert.equal(p.get('.ex-name').textContent, 'Exercise');
  p.get('.resume-continue').click();
  assert.ok(hidden(p.get('.resume-overlay')));
  p.advance(14);
  assert.equal(p.get('.ex-name').textContent, 'Exercise');
  p.advance(1);
  assert.equal(p.get('.ex-name').textContent, 'Next');
});

test('reload during the second timed side preserves its side and remaining time', () => {
  const items = [item({ unit: 'seconds', value: 5, perSide: true, rest: 20 }), item({ name: 'Next' })];
  const first = player(items);
  first.advance(13); // ready, first side, preparation, two seconds of the second side
  first.hide(); first.cleanup();
  const second = player(items, {}, { persistence: first.persistence });
  assert.match(second.get('.phase').textContent, /Side 2\/2/);
  assert.equal(second.get('.ex-value').textContent, '3″');
  assert.ok(!hidden(second.get('.resume-overlay')));
  second.advance(60);
  assert.equal(second.get('.ex-value').textContent, '3″');
  second.get('.resume-continue').click();
  second.advance(3);
  assert.equal(second.get('.rest-count').textContent, 20);
});

test('a removed side-rest checkpoint cannot run a program-owned pause', () => {
  const items = [item({ unit: 'seconds', value: 5, perSide: true }), item({ name: 'Next' })];
  const first = player(items);
  first.advance(3); first.cleanup();
  Object.assign(first.persistence.data.checkpoint.state, { phase: 'side-rest', nextSide: 2, remainingMs: 90000 });
  const second = player(items, {}, { persistence: first.persistence });
  assert.ok(hidden(second.get('.resume-overlay')));
  assert.equal(second.get('.rest-count').textContent, 3);
  second.advance(3);
  assert.match(second.get('.phase').textContent, /Side 1\/2/);
  assert.equal(second.completed.length, 0);
});

test('reload during normal rest keeps its destination and Skip starts that exercise', () => {
  const items = [item({ rest: 90 }), item({ id: 'b', name: 'Next' })];
  const first = player(items);
  first.advance(3); first.get('.ctl-done').click(); first.advance(17); first.cleanup();
  const second = player(items, {}, { persistence: first.persistence });
  assert.equal(second.get('.rest-count').textContent, 73);
  assert.equal(second.get('.rest-index').textContent, '2 / 2');
  second.get('.resume-continue').click();
  second.get('.rest-skip').click();
  assert.equal(second.get('.ex-name').textContent, 'Next');
});

test('switching away with exercise details open cannot accidentally resume on closing details', () => {
  const p = player([item({ unit: 'seconds', value: 20 })]);
  p.advance(8); p.get('.ex-visual').click(); p.advance(8); p.hide(); p.show();
  p.get('.info-close').click(); p.advance(30);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, 15000);
  assert.ok(!hidden(p.get('.resume-overlay')));
  p.get('.resume-continue').click(); p.advance(1);
  assert.equal(p.get('.ex-value').textContent, '14″');
});

test('session difficulty and a manual replacement restore from canonical payload data', () => {
  const items = [item({ scale: true, value: 10, repl: { id: 'replacement', name: 'Alternative', unit: 'seconds', value: 20, scale: true } })];
  const first = player(items);
  first.advance(3); first.get('.lvl-up').click(); first.get('.ex-visual').click(); first.get('.info-replace').click(); first.advance(4); first.cleanup();
  const second = player(items, {}, { persistence: first.persistence });
  assert.equal(second.get('.ex-name').textContent, 'Alternative');
  assert.equal(second.get('.ex-value').textContent, '18″'); // +10% 22 sec - 4
  assert.match(second.get('.lvl-label').textContent, /\+1/);
  second.get('.resume-restart').click(); second.advance(3);
  assert.equal(second.get('.ex-name').textContent, 'Exercise');
  assert.equal(second.get('.ex-value').textContent, '× 10');
});

for (const mismatch of [{ userId: 2 }, { resumeKey: 'new-plan:0:1' }, { position: 1 }]) {
  test('different account, plan or cycle does not restore a stale workout: ' + JSON.stringify(mismatch), () => {
    const items = [item({ unit: 'seconds', value: 20 })];
    const first = player(items); first.advance(10); first.cleanup();
    const second = player(items, mismatch, { persistence: first.persistence });
    assert.ok(hidden(second.get('.resume-overlay')));
    second.advance(3);
    assert.equal(second.get('.ex-value').textContent, '20″');
  });
}

test('malformed snapshot is ignored without breaking a new workout', () => {
  const items = [item()];
  const first = player(items); first.advance(3); first.cleanup();
  first.persistence.data.checkpoint.state = { ...first.persistence.data.checkpoint.state, phase: 'complete', completion: { eventId: 'x' }, swaps: { constructor: 'x' } };
  const second = player(items, {}, { persistence: first.persistence });
  second.advance(3);
  assert.equal(second.get('.ex-name').textContent, 'Exercise');
  assert.equal(second.completed.length, 0);
});

test('offline completion is pending across reload, reuses its event and celebrates only after acknowledgment', async () => {
  const persistence = memoryState({ rejectCompletion: true });
  const items = [item()];
  const first = player(items, {}, { persistence });
  first.advance(3); first.get('.ctl-done').click();
  await settle();
  const event = first.completed[0];
  assert.equal(first.get('.save-status').textContent, 'Waiting for connection');
  assert.ok(!hidden(first.get('.save-retry')));
  assert.equal(first.get('.done-overlay').children.filter(n => n.className === 'confetti-piece').length, 0);
  first.cleanup();
  const second = player(items, {}, { persistence });
  await settle();
  assert.equal(second.completed.length, 1);
  assert.equal(second.completed[0].eventId, event.eventId);
  assert.equal(second.completed[0].completedAt, event.completedAt);
  persistence.options.rejectCompletion = false;
  second.get('.save-retry').click();
  await settle();
  assert.equal(second.get('.save-status').textContent, 'Saved');
  assert.equal(persistence.data.checkpoint, null);
  assert.equal(second.get('.done-overlay').children.filter(n => n.className === 'confetti-piece').length, 60);
});

test('unavailable storage allows exercise and explains that offline completion is not durable', async () => {
  const p = player([item()], {}, { storageUnavailable: true, rejectCompletion: true });
  p.advance(3); p.get('.ctl-done').click();
  await settle();
  assert.equal(p.get('.save-status').textContent, 'Keep this page open');
  assert.ok(!hidden(p.get('.save-retry')));
});

test('duplicate resume events do not create extra timers or wake locks; late locks are released after cleanup', async () => {
  let resolveWake, requests = 0, releases = 0;
  const navigator = { wakeLock: { request() { requests++; return new Promise(resolve => { resolveWake = resolve; }); } } };
  const p = player([item({ unit: 'seconds', value: 20 })], {}, { navigator });
  p.show(); p.show();
  assert.equal(requests, 1);
  p.advance(8); p.hide(); p.show(); p.show();
  assert.equal(requests, 1);
  p.get('.resume-continue').click(); p.show();
  p.advance(1);
  assert.equal(p.get('.ex-value').textContent, '14″');
  p.cleanup();
  resolveWake({ released: false, release() { releases++; return Promise.resolve(); }, addEventListener() {} });
  await settle();
  assert.equal(releases, 1);
  assert.equal(requests, 1);
});

test('real shared storage restores a paused timer in a fresh page and records the same finished event', async () => {
  const storageValues = new Map(), requests = [];
  const items = [item({ unit: 'seconds', value: 20 })];
  const first = player(items, {}, { realStore: true, storageValues });
  first.advance(8.25); first.hide(); first.cleanup();
  const second = player(items, {}, {
    realStore: true, storageValues,
    fetch: async (url, init) => {
      const event = JSON.parse(init.body);
      assert.equal(url, '/api/workouts/complete');
      const pending = second.stateStore.pending(1);
      assert.equal(pending[0].eventId, event.eventId); // persisted before send
      requests.push(event);
      return { ok: true, json: async () => ({ saved: true, eventId: event.eventId, position: 1, streak: 2 }) };
    },
  });
  assert.equal(second.get('.ex-value').textContent, '15″');
  second.advance(60);
  assert.equal(requests.length, 0);
  second.get('.resume-continue').click(); second.advance(15);
  await settle();
  assert.equal(requests.length, 1);
  assert.equal(second.get('.save-status').textContent, 'Saved');
  assert.equal(second.stateStore.pending(1).length, 0);
  assert.equal(second.stateStore.getResume(1, 0, 'plan'), null);
});

test('real module retries a pending finish after page eviction without a new event or false success', async () => {
  const storageValues = new Map(), requests = [];
  const first = player([item()], {}, { realStore: true, storageValues });
  first.advance(3); first.get('.ctl-done').click(); await settle();
  const originalEvent = first.stateStore.pending(1)[0];
  assert.equal(first.get('.save-status').textContent, 'Waiting for connection');
  assert.ok(hidden(first.get('.applause')));
  assert.ok(hidden(first.get('.done-streak')));
  first.cleanup();
  const second = player([item()], {}, { realStore: true, storageValues, fetch: async (url, init) => {
    const event = JSON.parse(init.body); requests.push(event);
    return { ok: true, json: async () => ({ saved: true, eventId: event.eventId, position: 1, streak: 3 }) };
  } });
  await settle();
  assert.equal(requests.length, 1);
  assert.equal(requests[0].eventId, originalEvent.eventId);
  assert.equal(requests[0].completedAt, originalEvent.completedAt);
  assert.equal(second.get('.save-status').textContent, 'Saved');
  assert.ok(!hidden(second.get('.applause')));
  assert.equal(second.stateStore.getResume(1, 0, 'plan'), null);
});

test('an expired scheduled restart resolves the current day before reading an old checkpoint', () => {
  const persistence = memoryState();
  const first = player([item({ unit: 'seconds', value: 20 })], {}, { persistence });
  first.advance(8); first.cleanup();
  const original = JSON.stringify(persistence.data.checkpoint);
  const second = player([item({ unit: 'seconds', value: 20 })], { restartAt: '2026-10-24T07:59:59Z' }, { persistence });
  assert.deepEqual(second.navigations, ['/start']);
  assert.equal(JSON.stringify(persistence.data.checkpoint), original);
  second.advance(30);
  assert.equal(second.completed.length, 0);
});

test('a scheduled restart redirects on return, but never navigates while hidden', () => {
  const p = player([item({ unit: 'seconds', value: 20 })], { restartAt: '2026-10-24T08:00:10Z' });
  p.advance(8); p.hide(); p.advance(100);
  assert.equal(p.navigations.length, 0);
  p.show(); p.show();
  assert.deepEqual(p.navigations, ['/start']);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, 15000);
  p.advance(20);
  assert.equal(p.completed.length, 0);
});

for (const control of ['.resume-continue', '.resume-restart']) {
  test(control + ' checks the restart deadline while already showing the paused page', () => {
    const p = player([item({ unit: 'seconds', value: 20 })], { restartAt: '2026-10-24T08:00:10Z' });
    p.advance(8); p.hide(); p.show(); p.advance(3);
    p.get(control).click();
    assert.deepEqual(p.navigations, ['/start']);
    p.advance(30);
    assert.equal(p.completed.length, 0);
  });
}

test('invalid restart deadlines are ignored and a finished pending event is preserved after the deadline', async () => {
  const invalid = player([item()], { restartAt: 'not-a-date' });
  invalid.advance(3);
  assert.equal(invalid.get('.ex-name').textContent, 'Exercise');
  assert.equal(invalid.navigations.length, 0);
  const p = player([item()], { restartAt: '2026-10-24T08:00:10Z' }, { rejectCompletion: true });
  p.advance(3); p.get('.ctl-done').click(); await settle();
  const event = p.completed[0];
  p.hide(); p.advance(100); p.show();
  assert.equal(p.navigations.length, 0);
  assert.equal(p.completed[0].eventId, event.eventId);
  assert.equal(p.get('.save-status').textContent, 'Waiting for connection');
});

const demo = (kind = 'alternating', fields = {}) => ({ src: '/poses/demo', frames: 3, kind, ...fields });
const demoPosition = p => p.get('.ex-visual').style.backgroundPosition;

test('alternating demonstration shows both legs through neutral without adding workout steps', () => {
  const p = player([item({ perSide: true, animation: demo() }), item({ id: 'next' })], { selfManagedSides: true });
  p.advance(3);
  const positions = [demoPosition(p)];
  for (let n = 0; n < 4; n++) { p.advance(1); positions.push(demoPosition(p)); }
  assert.deepEqual(positions, ['0% 50%', '50% 50%', '0% 50%', '100% 50%', '0% 50%']);
  assert.ok(hidden(p.get('.ex-svg')));
  assert.match(p.get('.ex-value').textContent, /6.*per side/);
  p.get('.ctl-done').click();
  assert.equal(p.get('.ex-index').textContent, '2 / 2');
  assert.equal(p.get('.ex-visual').style.backgroundImage, 'none');
  assert.ok(hidden(p.get('.ex-visual')));
});

test('movement returns through its middle and a second strip demonstrates the other side', () => {
  const p = player([item({ animation: demo('movement', { otherSrc: '/poses/other' }) })]);
  p.advance(3);
  const shown = [];
  for (let n = 0; n < 11; n++) {
    shown.push([p.get('.ex-visual').style.backgroundImage, demoPosition(p)]);
    p.advance(1);
  }
  assert.deepEqual(shown.map(x => x[1]), ['0% 50%', '50% 50%', '100% 50%', '50% 50%', '0% 50%', '0% 50%', '50% 50%', '100% 50%', '50% 50%', '0% 50%', '0% 50%']);
  assert.ok(shown.slice(0, 5).every(x => x[0] === 'url("/poses/demo")'));
  assert.ok(shown.slice(5, 10).every(x => x[0] === 'url("/poses/other")'));
});

test('a forward/back step sequence plays all phases in order instead of reversing the lead foot', () => {
  const p = player([item({ animation: demo('sequence', { otherSrc: '/poses/return' }) })]);
  p.advance(3);
  const shown = [];
  for (let n = 0; n < 7; n++) {
    shown.push([p.get('.ex-visual').style.backgroundImage, demoPosition(p)]);
    p.advance(1);
  }
  assert.deepEqual(shown, [
    ['url("/poses/demo")', '0% 50%'], ['url("/poses/demo")', '50% 50%'], ['url("/poses/demo")', '100% 50%'],
    ['url("/poses/return")', '0% 50%'], ['url("/poses/return")', '50% 50%'], ['url("/poses/return")', '100% 50%'],
    ['url("/poses/demo")', '0% 50%'],
  ]);
});

test('detail demonstration can be paused and stepped without consuming exercise time', () => {
  const p = player([item({ unit: 'seconds', value: 20, animation: demo() })]);
  p.advance(5);
  p.get('.ex-visual').click();
  const remaining = p.persistence.data.checkpoint.state.remainingMs;
  assert.ok(!hidden(p.get('.info-overlay')));
  p.advance(3);
  p.get('.demo-toggle').click();
  const position = demoPosition(p);
  p.advance(10);
  assert.equal(demoPosition(p), position);
  p.get('.demo-next').click();
  assert.notEqual(demoPosition(p), position);
  p.get('.demo-prev').click();
  assert.equal(demoPosition(p), position);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, remaining);
  p.get('.info-close').click();
  p.advance(1);
  assert.equal(p.persistence.data.checkpoint.state.remainingMs, remaining - 1000);
  p.cleanup();
  const stopped = demoPosition(p);
  p.advance(20);
  assert.equal(demoPosition(p), stopped);
});

test('hidden app pauses the demonstration and requires Continue before restarting', () => {
  const p = player([item({ animation: demo() })]);
  p.advance(4);
  p.get('.ex-visual').click();
  p.hide();
  const stopped = demoPosition(p);
  p.advance(10);
  p.show();
  p.advance(10);
  assert.equal(demoPosition(p), stopped);
  assert.ok(hidden(p.get('.info-overlay')));
  p.get('.resume-continue').click();
  p.advance(1);
  assert.equal(demoPosition(p), '50% 50%');
});

test('missing animation falls back to the still and only current exercise strips are loaded', () => {
  const p = player([item({ svg: '/poses/still', animation: demo() }), item({ id: 'later', animation: demo('movement', { src: '/poses/later' }) })]);
  p.advance(3);
  assert.deepEqual(p.images.map(image => image.src), ['/poses/demo']);
  p.images[0].onerror();
  assert.equal(p.get('.ex-svg').src, '/poses/still');
  assert.ok(!hidden(p.get('.ex-svg')));
  assert.equal(p.get('.ex-visual').style.backgroundImage, 'none');
  p.advance(4);
  assert.equal(p.images.length, 1);
});

test('reduced motion keeps a still frame but allows intentional playback in details', () => {
  const p = player([item({ animation: demo() })], {}, { reducedMotion: true });
  p.advance(10);
  assert.equal(demoPosition(p), '0% 50%');
  p.get('.ex-visual').click();
  p.advance(2);
  assert.equal(demoPosition(p), '0% 50%');
  p.get('.demo-toggle').click();
  p.advance(1);
  assert.equal(demoPosition(p), '50% 50%');
});

test('compass demonstration visits all four directions with a return to neutral', () => {
  const p = player([item({ animation: demo('alternating', { otherSrc: '/poses/back-left' }) })]);
  p.advance(3);
  const shown = [];
  for (let n = 0; n < 8; n++) {
    shown.push([p.get('.ex-visual').style.backgroundImage, demoPosition(p)]);
    p.advance(1);
  }
  assert.deepEqual(shown.map(x => x[1]), ['0% 50%', '50% 50%', '0% 50%', '100% 50%', '0% 50%', '50% 50%', '0% 50%', '100% 50%']);
  assert.ok(shown.slice(0, 4).every(x => x[0] === 'url("/poses/demo")'));
  assert.ok(shown.slice(4).every(x => x[0] === 'url("/poses/back-left")'));
});

test('separate timed sides restart the demonstration when the side changes', () => {
  const p = player([item({ unit: 'seconds', value: 5, perSide: true, animation: demo('movement') })]);
  p.advance(11);
  assert.match(p.get('.phase').textContent, /Side 2\/2/);
  assert.equal(demoPosition(p), '0% 50%');
  p.advance(1);
  assert.equal(demoPosition(p), '50% 50%');
});
