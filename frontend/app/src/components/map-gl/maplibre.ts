import { Map as MaplibreMap, setWorkerUrl } from 'maplibre-gl'
import workerUrl from 'maplibre-gl/dist/maplibre-gl-worker.mjs?worker&url'

// maplibre-gl 6 locates its worker from a template literal against
// import.meta.url. Both operands are runtime values, so no bundler can see the
// reference and emit the file, and the url ends up pointing into a directory
// that only exists inside the package. Hand it one the bundler does emit;
// ?worker rather than ?url because the worker imports maplibre-gl-shared.mjs.
setWorkerUrl(workerUrl)

// Re-exported so the worker url cannot be configured after a Map already exists.
export { MaplibreMap }
