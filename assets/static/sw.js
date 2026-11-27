// Minimal service worker: enables installability and caches static assets
// (exercise illustrations, icons) for fast repeat loads and offline viewing.
// Dynamic doors pages always go to the network.

const CACHE = "morning-v1";

function retiredArtwork(request) {
  const path = new URL(request.url).pathname;
  return path.startsWith("/static/ex/") || /^\/static\/characters\/[mfbg]\.webp$/.test(path);
}

self.addEventListener("install", () => self.skipWaiting());

self.addEventListener("activate", (e) => {
  e.waitUntil((async () => {
    const keys = await caches.keys();
    await Promise.all(keys.filter((k) => k !== CACHE).map((k) => caches.delete(k)));
    const cache = await caches.open(CACHE);
    const requests = await cache.keys();
    await Promise.all(requests.filter(retiredArtwork).map((r) => cache.delete(r)));
    await self.clients.claim();
  })());
});

self.addEventListener("fetch", (e) => {
  const url = new URL(e.request.url);
  if (e.request.method !== "GET" || !url.pathname.startsWith("/static/")) return;
  e.respondWith(
    caches.open(CACHE).then((cache) =>
      cache.match(e.request).then((hit) => {
        if (hit) return hit;
        return fetch(e.request).then((resp) => {
          if (resp && resp.status === 200) cache.put(e.request, resp.clone());
          return resp;
        });
      })
    )
  );
});
