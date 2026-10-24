import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import vm from 'node:vm';

// Drive the shipped player with a small DOM and a deterministic clock: these
// tests check actual transitions, Done/Skip buttons and visible side labels.
function player(items, extra = {}) {
  const nodes = new Map();
  const timers = new Map();
  const completed = [];
  const spoken = [];
  let now = 0, nextTimer = 1;
  function element() {
    const classes = new Set();
    const events = new Map();
    return {
      textContent: '', style: {}, children: [], disabled: false,
      classList: {
        add(...xs) { xs.forEach(x => classes.add(x)); },
        remove(...xs) { xs.forEach(x => classes.delete(x)); },
        contains(x) { return classes.has(x); },
        toggle(x, yes) { if (yes ?? !classes.has(x)) classes.add(x); else classes.delete(x); },
      },
      addEventListener(name, fn) { events.set(name, fn); },
      appendChild(child) { this.children.push(child); },
      removeAttribute(name) { delete this[name]; },
      click() { if (!this.disabled) events.get('click')?.({ preventDefault() {} }); },
    };
  }
  const root = element();
  const get = selector => {
    if (!nodes.has(selector)) {
      const node = element();
      node.parentElement = element();
      nodes.set(selector, node);
    }
    return nodes.get(selector);
  };
  root.querySelector = get;
  const document = {
    querySelector: () => root, createElement: element,
    addEventListener() {}, removeEventListener() {}, visibilityState: 'visible',
  };
  const schedule = (fn, delay, interval) => {
    const id = nextTimer++;
    timers.set(id, { fn, due: now + delay, interval });
    return id;
  };
  const context = {
    document, navigator: {}, window: {
      addEventListener() {}, removeEventListener() {},
      speechSynthesis: { getVoices: () => [], cancel() {}, speak: utterance => spoken.push(utterance.text) },
    },
    SpeechSynthesisUtterance: function (text) { this.text = text; },
    $data: () => ({ items, voiceMode: 'off', cues: { switch_side: 'Change sides', go: 'Go' }, labels: { side: 'Side', per_side: 'per side' }, ...extra }),
    $hook: (...args) => { completed.push(args); return Promise.resolve(1); },
    $sys: { clean() {} },
    setInterval: (fn, delay) => schedule(fn, delay, delay), clearInterval: id => timers.delete(id),
    setTimeout: (fn, delay) => schedule(fn, delay, 0), clearTimeout: id => timers.delete(id),
  };
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
  return { get, advance, completed, spoken };
}
const item = (fields = {}) => ({ id: 'a', name: 'Exercise', unit: 'reps', value: 6, slot: 'main', round: 1, rest: 0, ...fields });
const hidden = node => node.classList.contains('hidden');

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

for (const unit of ['reps', 'seconds']) {
  test(`${unit}: both sides use the prescribed 10-second switch, then the 90-second rest`, () => {
    const p = player([item({ unit, value: 5, perSide: true, sideRest: 10, rest: 90 }), item({ id: 'b', name: 'Next' })]);
    p.advance(3);
    assert.match(p.get('.phase').textContent, /Side 1\/2/);
    if (unit === 'seconds') p.advance(5); else p.get('.ctl-done').click();
    assert.equal(p.get('.rest-title').textContent, 'Change sides');
    assert.equal(p.get('.rest-count').textContent, 10);
    assert.equal(p.get('.rest-index').textContent, '1 / 2');
    p.advance(9);
    assert.equal(p.get('.rest-count').textContent, 1);
    p.advance(1);
    assert.ok(hidden(p.get('.rest-overlay')));
    assert.match(p.get('.phase').textContent, /Side 2\/2/);
    if (unit === 'seconds') p.advance(5); else p.get('.ctl-done').click();
    assert.equal(p.get('.rest-count').textContent, 90);
    assert.equal(p.get('.rest-index').textContent, '2 / 2');
    p.advance(90);
    assert.equal(p.get('.ex-name').textContent, 'Next');
  });
}

test('switch pause Skip starts the second side, without skipping the exercise', () => {
  const p = player([item({ perSide: true, sideRest: 10 }), item({ name: 'Next' })]);
  p.advance(3);
  p.get('.ctl-done').click();
  p.get('.rest-skip').click();
  assert.match(p.get('.phase').textContent, /Side 2\/2/);
  assert.equal(p.get('.ex-name').textContent, 'Exercise');
  p.get('.ctl-done').click();
  assert.equal(p.get('.ex-name').textContent, 'Next');
  assert.ok(hidden(p.get('.rest-overlay')));
});

test('legacy reps use one Done tap for both sides, legacy timers switch immediately', () => {
  const p = player([item({ perSide: true }), item({ unit: 'seconds', value: 5, perSide: true })]);
  p.advance(3);
  assert.equal(p.get('.ex-value').textContent, '× 6 per side');
  p.get('.ctl-done').click();
  p.advance(5);
  assert.match(p.get('.phase').textContent, /Side 2\/2/);
  assert.ok(hidden(p.get('.rest-overlay')));
  p.advance(5);
  assert.equal(p.completed.length, 1);
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

test('narration introduces mistakes explicitly, including descriptive and legacy negative wording', () => {
  const reminder = 'Напоминание о частых ошибках';
  for (const mistake of ['Сводить колени или отрывать пятки.', 'Не своди колени.']) {
    const p = player([item({ wrong: [mistake] })], {
      voiceMode: 'normal', labels: { mistake_reminder: reminder },
    });
    p.advance(6);
    assert.ok(p.spoken.includes(reminder + ': ' + mistake));
    assert.ok(!p.spoken.includes(mistake));
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
    p.get('.info-btn').click();
    assert.ok(!hidden(p.get('.info-overlay')));
    assert.ok(hidden(p.get('.ex-svg')));
    assert.equal(p.get('.info-warn').textContent, expected);
    assert.ok(!hidden(p.get('.info-warn-sec')));
  }
});
