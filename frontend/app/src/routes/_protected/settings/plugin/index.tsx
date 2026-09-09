import { createFileRoute, Link } from '@tanstack/react-router'
import { Suspense, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'
import {
  Badge,
  Button,
  Loading,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@green-ecolution/ui'
import { pluginsQuery, organizationQueries } from '@/api/queries'
import { usePluginMutations } from '@/hooks/usePluginMutations'
import { Can } from '@/lib/auth/Can'
import { useDateLocale } from '@/lib/i18n/useFormatters'
import PluginInstallDialog, {
  type PluginInstallPayload,
} from '@/components/settings/plugin/PluginInstallDialog'
import PluginKeyDialog from '@/components/settings/plugin/PluginKeyDialog'
import { formatLastSeenAt, organizationNameOf } from '@/components/settings/plugin/pluginList'

export const Route = createFileRoute('/_protected/settings/plugin/')({
  component: PluginView,
})

function PluginView() {
  const { t } = useTranslation('settings')

  return (
    <div className="container mt-6">
      <article className="mb-10 flex flex-wrap items-start justify-between gap-4 2xl:w-4/5">
        <div>
          <h1 className="font-lato font-bold text-3xl mb-4 lg:text-4xl xl:text-5xl">
            {t('plugin.overviewTitle')}
          </h1>
          <p>{t('plugin.overviewIntro')}</p>
        </div>
      </article>

      <Suspense fallback={<Loading className="mt-10 justify-center" label={t('plugin.loading')} />}>
        <PluginOverview />
      </Suspense>
    </div>
  )
}

const PluginOverview = () => {
  const { t } = useTranslation(['settings', 'common'])
  const dateLocale = useDateLocale()
  const { data: pluginList } = useQuery(pluginsQuery())
  const { data: organizations } = useQuery(organizationQueries.list())
  const { installPlugin } = usePluginMutations()

  const [installOpen, setInstallOpen] = useState(false)
  const [issuedKey, setIssuedKey] = useState<string | null>(null)

  const handleInstall = (payload: PluginInstallPayload) => {
    installPlugin.mutate(
      {
        slug: payload.slug,
        name: payload.name,
        description: payload.description,
        organizationId: payload.organizationId,
        permissions: payload.permissions,
        requiredPermissions: payload.requiredPermissions,
        frontend: payload.frontend,
      },
      {
        onSuccess: (response) => {
          setInstallOpen(false)
          setIssuedKey(response.key)
        },
      },
    )
  }

  return (
    <>
      <div className="mb-4 flex justify-end">
        <Can permission={['plugin:create']}>
          <Button type="button" onClick={() => setInstallOpen(true)}>
            {t('plugin.installButton')}
          </Button>
        </Can>
      </div>

      {pluginList && pluginList.length > 0 ? (
        <div className="overflow-x-auto rounded-xl border border-dark-100">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>{t('plugin.columns.status')}</TableHead>
                <TableHead>{t('plugin.columns.slug')}</TableHead>
                <TableHead>{t('plugin.columns.organization')}</TableHead>
                <TableHead>{t('plugin.columns.lastSeen')}</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {pluginList.map((plugin) => (
                <TableRow key={plugin.slug} className="cursor-pointer">
                  <TableCell colSpan={4} className="p-0">
                    <Link
                      to="/settings/plugin/$slug"
                      params={{ slug: plugin.slug }}
                      aria-label={t('plugin.detailAriaLabel', { name: plugin.name })}
                      className="grid grid-cols-4 gap-4 px-4 py-3 hover:bg-dark-50 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
                    >
                      <span>
                        <Badge variant={plugin.enabled ? 'success' : 'muted'}>
                          {plugin.enabled
                            ? t('plugin.status.enabled')
                            : t('plugin.status.disabled')}
                        </Badge>
                      </span>
                      <span>
                        <span className="block font-medium text-dark">{plugin.name}</span>
                        <span className="block font-mono text-sm text-dark-500">{plugin.slug}</span>
                      </span>
                      <span>
                        {organizationNameOf(plugin.organizationId, organizations ?? []) ??
                          plugin.organizationId}
                      </span>
                      <span>{formatLastSeenAt(plugin.lastSeenAt, dateLocale, t)}</span>
                    </Link>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      ) : (
        <div className="text-center mt-6">
          <p className="text-dark-500">{t('plugin.empty')}</p>
        </div>
      )}

      <PluginInstallDialog
        open={installOpen}
        organizations={organizations ?? []}
        saving={installPlugin.isPending}
        onOpenChange={setInstallOpen}
        onSubmit={handleInstall}
      />

      <PluginKeyDialog
        open={issuedKey !== null}
        variant="installed"
        apiKey={issuedKey ?? ''}
        onOpenChange={(open) => {
          if (!open) setIssuedKey(null)
        }}
      />
    </>
  )
}
