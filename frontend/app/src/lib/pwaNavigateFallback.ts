// Paths the service worker must leave to the network. Workbox answers every
// request with `mode: 'navigate'` from the precached index.html, which turns a
// PDF download into a 25 KB index.html under a .pdf name and swallows the
// backend's own routes.
export const navigateFallbackDenylist = [/^\/handbook\//, /^\/api\//]

export function servedByNetwork(path: string): boolean {
  return navigateFallbackDenylist.some((pattern) => pattern.test(path))
}
