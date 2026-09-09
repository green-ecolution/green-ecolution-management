import type { TFunction } from 'i18next'
import type { PluginFrontendDto } from '@/api/backendApi'

export type FrontendMode = 'none' | 'external' | 'proxied'

/**
 * Mirrors the backend's own check (`PluginFrontendDto::into_domain` in
 * `http/v1/dto/plugin.rs`) so an invalid external URL is caught before the
 * request round-trips, not just reported back as a 400.
 */
export const validateTarget = (
  mode: FrontendMode,
  value: string,
  t: TFunction<'settings'>,
): string | null => {
  if (mode === 'none') return null
  const trimmed = value.trim()
  if (trimmed === '') return t('plugin.install.targetRequired')

  if (mode === 'external') {
    let url: URL
    try {
      url = new URL(trimmed)
    } catch {
      return t('plugin.install.targetInvalidUrl')
    }
    const isLocalhost = url.hostname === 'localhost'
    if (url.protocol !== 'https:' && !isLocalhost) {
      return t('plugin.install.targetRequiresHttps')
    }
    return null
  }

  const match = /^([^\s:]+):(\d+)$/.exec(trimmed)
  if (!match) return t('plugin.install.targetInvalidHostPort')
  const port = Number(match[2])
  if (port < 1 || port > 65535) return t('plugin.install.targetInvalidPort')
  return null
}

export const buildFrontendDto = (mode: FrontendMode, target: string): PluginFrontendDto => {
  if (mode === 'none') return { mode: 'none' }
  return { mode, target: target.trim() }
}
