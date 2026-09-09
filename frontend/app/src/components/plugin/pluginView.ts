import type { PluginResponse } from '@/api/backendApi'

/**
 * What the plugin viewer route renders for a given plugin. Only `external`
 * mode has a browser-reachable url today; `proxied` and `none` each get a
 * distinct message rather than being folded into one generic fallback, since
 * "in the cluster, not embeddable yet" and "this plugin has no view" are
 * different facts to report.
 */
export type PluginViewKind =
  { kind: 'iframe'; target: string } | { kind: 'proxied' } | { kind: 'unavailable' }

export const pluginViewKind = (
  plugin: Pick<PluginResponse, 'frontendMode' | 'frontendTarget'>,
): PluginViewKind => {
  if (plugin.frontendMode === 'proxied') return { kind: 'proxied' }
  if (plugin.frontendMode === 'external' && plugin.frontendTarget) {
    return { kind: 'iframe', target: plugin.frontendTarget }
  }
  return { kind: 'unavailable' }
}
