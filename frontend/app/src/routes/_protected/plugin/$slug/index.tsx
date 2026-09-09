import { createFileRoute, redirect } from '@tanstack/react-router'
import { useEffect, useRef } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'
import { createPluginHost, type PluginContext } from '@green-ecolution/plugin-interface'
import { pluginQuery, userQueries } from '@/api/queries'
import {
  permissionsOf,
  satisfies,
  UNRESTRICTED,
  type PermissionRequirement,
} from '@/lib/auth/permissions'
import { readAuthBypass } from '@/lib/auth/runtimeConfig'
import { useCurrentUser } from '@/lib/auth/useCurrentUser'
import { languageOf } from '@/lib/i18n/languages'
import { entityNotFound, pendingLoading, prefetch } from '@/lib/router'

export const Route = createFileRoute('/_protected/plugin/$slug/')({
  component: PluginViewPage,
  pendingComponent: pendingLoading({ key: 'settings:plugin.view.loading' }),
  beforeLoad: async ({ context: { queryClient }, params: { slug } }) => {
    const plugin = await queryClient.ensureQueryData(pluginQuery(slug))

    // Only an externally hosted frontend can be embedded today; a proxied
    // target is served through the backend and has no browser-reachable URL
    // yet (see the follow-up proxy plan), so it redirects the same as a
    // missing one.
    if (plugin.frontendMode !== 'external' || !plugin.frontendTarget) {
      throw redirect({ to: '/settings/plugin' })
    }

    const perms = readAuthBypass()
      ? UNRESTRICTED
      : permissionsOf(await queryClient.ensureQueryData(userQueries.me()))
    if (!satisfies(perms, plugin.requiredPermissions as PermissionRequirement)) {
      throw redirect({ to: '/settings/plugin' })
    }
  },
  loader: ({ context: { queryClient }, params: { slug } }) =>
    prefetch(queryClient, pluginQuery(slug), 'pluginQuery'),
  errorComponent: entityNotFound({
    entityName: { key: 'settings:plugin.entityName' },
    backTo: '/settings/plugin',
    backLabel: { key: 'settings:plugin.notFoundBackLabel' },
  }),
})

function PluginViewPage() {
  const { slug } = Route.useParams()
  const { data: plugin } = useSuspenseQuery(pluginQuery(slug))
  const { i18n } = useTranslation()
  const { firstName, lastName, username } = useCurrentUser()
  const frameRef = useRef<HTMLIFrameElement>(null)

  const displayName = `${firstName} ${lastName}`.trim() || username
  const target = plugin.frontendTarget

  useEffect(() => {
    const iframe = frameRef.current
    if (!iframe || !target) return

    const context: PluginContext = {
      locale: languageOf(i18n.language),
      // The app has no theme switch yet, so every plugin view starts in light mode.
      theme: 'light',
      user: { displayName },
      plugin: { slug: plugin.slug },
    }

    return createPluginHost(iframe, { origin: new URL(target).origin, context })
  }, [target, plugin.slug, displayName, i18n.language])

  return (
    <iframe
      ref={frameRef}
      src={target ?? undefined}
      title={plugin.name}
      // allow-same-origin is safe here only because frontend_target is always a
      // foreign origin (enforced on install/update); serving a plugin from the
      // app's own origin would give it an unrestricted same-origin document.
      // eslint-disable-next-line react-dom/no-unsafe-iframe-sandbox -- see comment above
      sandbox="allow-scripts allow-forms allow-popups allow-same-origin"
      referrerPolicy="no-referrer"
      allow=""
      className="w-full h-[calc(100vh-12rem)] border-0"
    />
  )
}
