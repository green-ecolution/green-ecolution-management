import { createFileRoute, redirect } from '@tanstack/react-router'
import { useEffect, useRef } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'
import { Alert, AlertContent, AlertDescription, AlertIcon } from '@green-ecolution/ui'
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
import { pluginViewKind } from '@/components/plugin/pluginView'

export const Route = createFileRoute('/_protected/plugin/$slug/')({
  component: PluginViewPage,
  pendingComponent: pendingLoading({ key: 'settings:plugin.view.loading' }),
  beforeLoad: async ({ context: { queryClient }, params: { slug } }) => {
    const plugin = await queryClient.ensureQueryData(pluginQuery(slug))

    // A missing or insufficient permission redirects rather than explaining,
    // same as every other route guard — but whether the view exists at all
    // (frontendMode) is not an access question, so it is rendered as a
    // message by the component instead of redirecting here, which would be
    // indistinguishable from "you may not".
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

function PluginViewNotice({ children }: { children: React.ReactNode }) {
  return (
    <div className="container mt-6">
      <Alert variant="info" className="flex items-start gap-3">
        <AlertIcon variant="info" />
        <AlertContent>
          <AlertDescription>{children}</AlertDescription>
        </AlertContent>
      </Alert>
    </div>
  )
}

function PluginViewPage() {
  const { slug } = Route.useParams()
  const { data: plugin } = useSuspenseQuery(pluginQuery(slug))
  const { t, i18n } = useTranslation('settings')
  const { firstName, lastName, username } = useCurrentUser()
  const frameRef = useRef<HTMLIFrameElement>(null)

  const displayName = `${firstName} ${lastName}`.trim() || username
  const view = pluginViewKind(plugin)
  const target = view.kind === 'iframe' ? view.target : null

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

  if (view.kind === 'proxied') {
    return <PluginViewNotice>{t('plugin.view.proxiedNotice')}</PluginViewNotice>
  }
  if (view.kind === 'unsafeOrigin') {
    return <PluginViewNotice>{t('plugin.view.unsafeOriginNotice')}</PluginViewNotice>
  }
  if (view.kind === 'unavailable') {
    return <PluginViewNotice>{t('plugin.view.noFrontendNotice')}</PluginViewNotice>
  }

  return (
    <iframe
      ref={frameRef}
      src={view.target}
      title={plugin.name}
      // allow-same-origin is safe here only because frontend_target is always a
      // foreign origin (enforced on install/update, and re-checked against
      // window.location.origin by pluginViewKind before this branch is reached);
      // serving a plugin from the app's own origin would give it an
      // unrestricted same-origin document.
      // eslint-disable-next-line react-dom/no-unsafe-iframe-sandbox -- see comment above
      sandbox="allow-scripts allow-forms allow-popups allow-same-origin"
      referrerPolicy="no-referrer"
      allow=""
      className="w-full h-[calc(100vh-12rem)] border-0"
    />
  )
}
