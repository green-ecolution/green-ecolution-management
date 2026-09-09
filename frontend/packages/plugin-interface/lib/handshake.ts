const NAMESPACE = 'green-ecolution'
const VERSION = 1

export interface PluginContext {
  locale: 'de' | 'en'
  theme: 'light' | 'dark'
  user: { displayName: string }
  plugin: { slug: string }
}

export type Envelope<T> = {
  ns: typeof NAMESPACE
  v: typeof VERSION
  type: string
  payload: T
}

function envelope<T>(type: string, payload: T): Envelope<T> {
  return { ns: NAMESPACE, v: VERSION, type, payload }
}

function isEnvelope(data: unknown): data is Envelope<unknown> {
  if (typeof data !== 'object' || data === null) return false
  const candidate = data as Record<string, unknown>
  return candidate.ns === NAMESPACE && candidate.v === VERSION && typeof candidate.type === 'string'
}

function isResizePayload(payload: unknown): payload is { height: number } {
  return (
    typeof payload === 'object' &&
    payload !== null &&
    typeof (payload as { height?: unknown }).height === 'number'
  )
}

export interface CreatePluginHostOptions {
  origin: string
  context: PluginContext
  onResize?: (height: number) => void
}

/**
 * Both origin and source are checked on every message: origin alone would
 * accept a message forged by any other same-origin frame, and source alone
 * would accept a message replayed from a different, unrelated origin.
 */
export function createPluginHost(
  iframe: HTMLIFrameElement,
  { origin, context, onResize }: CreatePluginHostOptions,
): () => void {
  const handleMessage = (event: MessageEvent) => {
    if (event.origin !== origin) return
    if (event.source !== iframe.contentWindow) return
    if (!isEnvelope(event.data)) return

    if (event.data.type === 'ge:hello') {
      iframe.contentWindow?.postMessage(envelope('ge:init', context), origin)
      return
    }

    if (event.data.type === 'ge:resize' && isResizePayload(event.data.payload)) {
      onResize?.(event.data.payload.height)
    }
  }

  window.addEventListener('message', handleMessage)
  return () => window.removeEventListener('message', handleMessage)
}

/**
 * Runs inside the plugin's own document. The target is '*' because the
 * envelope carries no secret; the host itself only ever answers a hello that
 * came from its own iframe's contentWindow (see createPluginHost). The
 * plugin's counterpart is always window.parent, so that is checked here too
 * — nothing secret travels this way, but both sides verify each other.
 */
export function connectToHost(): Promise<PluginContext> {
  return new Promise((resolve) => {
    const handleMessage = (event: MessageEvent) => {
      if (event.source !== window.parent) return
      if (!isEnvelope(event.data) || event.data.type !== 'ge:init') return
      window.removeEventListener('message', handleMessage)
      resolve(event.data.payload as PluginContext)
    }
    window.addEventListener('message', handleMessage)
    window.parent.postMessage(envelope('ge:hello', {}), '*')
  })
}

export function notifyResize(height: number): void {
  window.parent.postMessage(envelope('ge:resize', { height }), '*')
}
