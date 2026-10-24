// Shared by the player and home page. Every queued finish has its own storage
// key, so another tab saving a checkpoint cannot overwrite an unsent finish.
(function () {
  "use strict";
  if (window.MorningWorkoutState) return;

  var prefix = "morning.workout.v1:", activeUser = 0;
  var memory = Object.create(null), inFlight = Object.create(null), listeners = [];
  var retryTimer = null;

  function copy(value) { return JSON.parse(JSON.stringify(value)); }
  function own(object, key) { return Object.prototype.hasOwnProperty.call(object, key); }
  function userOK(id) { return Number.isSafeInteger(id) && id > 0; }
  function eventOK(event) {
    return event && userOK(event.userId) && /^[A-Za-z0-9_-]{32,128}$/.test(event.eventId) &&
      Number.isInteger(event.day) && event.day >= 1 && event.day <= 30 &&
      Number.isSafeInteger(event.position) && event.position >= 0 &&
      typeof event.completedAt === "string" && typeof event.completedOn === "string";
  }
  function key(userId, kind, eventId) { return prefix + userId + ":" + kind + (eventId ? ":" + eventId : ""); }
  function read(storageKey) {
    var raw;
    if (own(memory, storageKey)) raw = memory[storageKey];
    else {
      try { raw = window.localStorage.getItem(storageKey); } catch (_) { raw = null; }
    }
    if (!raw) return null;
    try { return JSON.parse(raw); } catch (_) { return null; }
  }
  function write(storageKey, value) {
    var raw;
    try { raw = JSON.stringify(value); } catch (_) { return false; }
    memory[storageKey] = raw;
    try { window.localStorage.setItem(storageKey, raw); return true; } catch (_) { return false; }
  }
  function remove(storageKey) {
    memory[storageKey] = null;
    try { window.localStorage.removeItem(storageKey); } catch (_) {}
  }
  function keys() {
    var found = Object.create(null);
    Object.keys(memory).forEach(function (storageKey) { found[storageKey] = true; });
    try {
      for (var i = 0; i < window.localStorage.length; i++) {
        var storageKey = window.localStorage.key(i);
        if (storageKey && storageKey.indexOf(prefix) === 0) found[storageKey] = true;
      }
    } catch (_) {}
    return Object.keys(found);
  }
  function notify(event) {
    listeners.slice().forEach(function (listener) {
      try { listener(copy(event)); } catch (_) {}
    });
  }
  function identity(payload) {
    return { userId: payload.userId, position: payload.position, day: payload.day,
      programKey: payload.programKey, resumeKey: payload.resumeKey };
  }
  function identityOK(payload) {
    return payload && userOK(payload.userId) && Number.isSafeInteger(payload.position) && payload.position >= 0 &&
      Number.isInteger(payload.day) && payload.day >= 1 && payload.day <= 30 &&
      typeof payload.programKey === "string" && payload.programKey !== "" &&
      typeof payload.resumeKey === "string" && payload.resumeKey !== "";
  }
  function checkpoint(userId) {
    var saved = read(key(userId, "checkpoint"));
    if (!saved || !identityOK(saved.identity) || saved.identity.userId !== userId ||
        !saved.state || typeof saved.state !== "object" || Array.isArray(saved.state)) return null;
    return saved;
  }
  function saveCheckpoint(payload, state) {
    if (!identityOK(payload) || !state || typeof state !== "object") return false;
    return write(key(payload.userId, "checkpoint"), { identity: identity(payload), state: state });
  }
  function loadCheckpoint(payload) {
    if (!identityOK(payload)) return null;
    var saved = checkpoint(payload.userId);
    if (!saved) return null;
    var matches = Object.keys(identity(payload)).every(function (field) {
      return saved.identity[field] === payload[field];
    });
    if (!matches) { remove(key(payload.userId, "checkpoint")); return null; }
    return saved.state;
  }
  function clearCheckpoint(payload) {
    if (!identityOK(payload)) return;
    var saved = checkpoint(payload.userId);
    if (saved && saved.identity.resumeKey === payload.resumeKey) remove(key(payload.userId, "checkpoint"));
  }
  function getResume(userId, position, programKey) {
    if (!userOK(userId)) return null;
    var saved = checkpoint(userId);
    if (!saved) return null;
    if (saved.identity.position !== position || saved.identity.programKey !== programKey) {
      remove(key(userId, "checkpoint"));
      return null;
    }
    return { day: saved.identity.day, url: "/day/" + saved.identity.day };
  }
  function ack(eventId, userId) {
    var receipt = read(key(userId, "ack", eventId));
    return receipt && receipt.saved === true && receipt.eventId === eventId &&
      Number.isSafeInteger(receipt.position) && receipt.position >= 0 &&
      Number.isSafeInteger(receipt.streak) && receipt.streak >= 0 ? receipt : null;
  }
  function pending(userId) {
    if (!userOK(userId)) return [];
    var pendingPrefix = key(userId, "pending") + ":";
    return keys().filter(function (storageKey) { return storageKey.indexOf(pendingPrefix) === 0; })
      .map(read).filter(function (event) { return eventOK(event) && event.userId === userId; });
  }
  function queueCompletion(event) {
    if (!eventOK(event)) return { durable: false };
    var pendingKey = key(event.userId, "pending", event.eventId);
    if (ack(event.eventId, event.userId)) return { durable: true };
    var first = read(pendingKey);
    // A retry must retain the original time/date/position, even if its caller
    // rebuilt the event after midnight or after another tab advanced the cycle.
    if (!eventOK(first) || first.userId !== event.userId || first.eventId !== event.eventId) first = event;
    return { durable: write(pendingKey, first) };
  }
  function acceptReceipt(event, receipt) {
    var durable = write(key(event.userId, "ack", event.eventId), receipt);
    remove(key(event.userId, "pending", event.eventId));
    var savedCheckpoint = checkpoint(event.userId);
    if (savedCheckpoint && savedCheckpoint.state.eventId === event.eventId) remove(key(event.userId, "checkpoint"));
    notify({ eventId: event.eventId, userId: event.userId, status: "saved", durable: durable, receipt: receipt });
    return receipt;
  }
  function sendCompletion(event) {
    if (!eventOK(event)) return Promise.reject(new Error("invalid_completion"));
    if (activeUser !== event.userId) return Promise.reject(new Error("account_mismatch"));
    var queued = queueCompletion(event);
    var saved = ack(event.eventId, event.userId);
    if (saved) return Promise.resolve(acceptReceipt(event, saved));
    var pendingKey = key(event.userId, "pending", event.eventId);
    event = read(pendingKey) || copy(event);
    if (inFlight[pendingKey]) return inFlight[pendingKey];
    var controller = typeof AbortController !== "undefined" ? new AbortController() : null;
    var timeout = controller ? setTimeout(function () { controller.abort(); }, 20000) : null;
    var request = Promise.resolve().then(function () {
      return window.fetch("/api/workouts/complete", {
        method: "POST", credentials: "same-origin", keepalive: true,
        headers: { "Content-Type": "application/json" }, body: JSON.stringify(event),
        signal: controller ? controller.signal : undefined,
      });
    }).then(function (response) {
      if (!response.ok) throw new Error("save_failed_" + response.status);
      return response.json();
    }).then(function (receipt) {
      if (!receipt || receipt.saved !== true || receipt.eventId !== event.eventId ||
          !Number.isSafeInteger(receipt.position) || receipt.position < 0 ||
          !Number.isSafeInteger(receipt.streak) || receipt.streak < 0) throw new Error("invalid_receipt");
      return acceptReceipt(event, receipt);
    }).catch(function (error) {
      notify({ eventId: event.eventId, userId: event.userId, status: "pending", durable: queued.durable });
      throw error;
    });
    inFlight[pendingKey] = request;
    notify({ eventId: event.eventId, userId: event.userId, status: "saving", durable: queued.durable });
    var clean = function () { if (timeout !== null) clearTimeout(timeout); delete inFlight[pendingKey]; };
    request.then(clean, clean);
    return request;
  }
  function flush(userId) {
    if (!userOK(userId) || userId !== activeUser) return Promise.resolve([]);
    return Promise.allSettled(pending(userId).map(sendCompletion));
  }
  function start(userId) {
    activeUser = userOK(userId) ? userId : 0;
    if (retryTimer === null) retryTimer = setInterval(function () {
      if (activeUser && (!window.navigator || window.navigator.onLine !== false)) flush(activeUser);
    }, 30000);
    return flush(activeUser);
  }
  function newEventId() {
    if (window.crypto && typeof window.crypto.randomUUID === "function") return window.crypto.randomUUID();
    var bytes = new Uint8Array(16);
    if (window.crypto && typeof window.crypto.getRandomValues === "function") window.crypto.getRandomValues(bytes);
    else for (var i = 0; i < bytes.length; i++) bytes[i] = Math.floor(Math.random() * 256);
    return Array.from(bytes).map(function (byte) { return byte.toString(16).padStart(2, "0"); }).join("");
  }
  window.MorningWorkoutState = {
    start: start, saveCheckpoint: saveCheckpoint, loadCheckpoint: loadCheckpoint,
    clearCheckpoint: clearCheckpoint, getResume: getResume, newEventId: newEventId,
    queueCompletion: queueCompletion, sendCompletion: sendCompletion, ack: ack,
    pending: pending, flush: flush,
    subscribe: function (listener) {
      listeners.push(listener);
      return function () { listeners = listeners.filter(function (item) { return item !== listener; }); };
    },
  };
  window.addEventListener("online", function () { flush(activeUser); });
  window.addEventListener("pageshow", function () { flush(activeUser); });
  document.addEventListener("visibilitychange", function () {
    if (document.visibilityState === "visible") flush(activeUser);
  });
  window.addEventListener("storage", function (event) {
    if (event.key && event.key.indexOf(prefix) === 0) {
      delete memory[event.key];
      var ackPrefix = key(activeUser, "ack") + ":";
      if (event.key.indexOf(ackPrefix) === 0) {
        var eventId = event.key.slice(ackPrefix.length);
        var receipt = ack(eventId, activeUser);
        if (receipt) notify({ eventId: eventId, userId: activeUser, status: "saved", durable: true, receipt: receipt });
      }
      if (event.key.indexOf(key(activeUser, "pending") + ":") === 0) flush(activeUser);
    }
  });
})();
