import type { GeoJSONSource } from 'maplibre-gl'
import type { GeoJSON } from 'geojson'

// maplibre-gl 6 made setData async (it resolves once the worker has parsed the
// data). Layer updates are fire-and-forget, so swallow the promise here instead
// of letting a rejection escape as an unhandled one.
export const setSourceData = (source: GeoJSONSource | undefined, data: GeoJSON): void => {
  source?.setData(data).catch((error: unknown) => console.error('setData failed:', error))
}
