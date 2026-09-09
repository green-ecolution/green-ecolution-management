import { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from '@tanstack/react-router'
import { useQuery } from '@tanstack/react-query'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  Badge,
  Button,
  FormField,
  Label,
  Switch,
  TextareaField,
} from '@green-ecolution/ui'
import type { PluginResponse } from '@/api/backendApi'
import { organizationQueries } from '@/api/queries'
import { usePluginMutations } from '@/hooks/usePluginMutations'
import { Can } from '@/lib/auth/Can'
import { useHasPermission } from '@/lib/auth/useHasPermission'
import { useDateLocale } from '@/lib/i18n/useFormatters'
import { intlLocale } from '@/lib/i18n/format'
import PluginPermissionMatrix from './PluginPermissionMatrix'
import PluginKeyDialog from './PluginKeyDialog'
import { accessVisiblePermissions, usePluginPermissionDraft } from './usePluginPermissionDraft'
import { formatLastSeenAt, organizationNameOf } from './pluginList'
import { buildFrontendDto, type FrontendMode } from './pluginFrontend'

interface PluginDetailPageProps {
  plugin: PluginResponse
}

const PluginDetailPage = ({ plugin }: PluginDetailPageProps) => {
  const { t, i18n } = useTranslation(['settings', 'common'])
  const navigate = useNavigate()
  const dateLocale = useDateLocale()
  const canUpdate = useHasPermission(['plugin:update'])
  const { data: organizations } = useQuery(organizationQueries.list())
  const { updatePlugin, rotatePluginKey, uninstallPlugin } = usePluginMutations()

  const [name, setName] = useState(plugin.name)
  const [description, setDescription] = useState(plugin.description ?? '')
  const [frontendMode, setFrontendMode] = useState<FrontendMode>(
    plugin.frontendMode as FrontendMode,
  )
  const [target, setTarget] = useState(plugin.frontendTarget ?? '')
  const {
    permissions,
    accessPermissions,
    togglePermission,
    toggleAccessPermission,
    submittedAccessPermissions,
  } = usePluginPermissionDraft(plugin.permissions, plugin.requiredPermissions)

  const [issuedKey, setIssuedKey] = useState<string | null>(null)
  const [rotateConfirmOpen, setRotateConfirmOpen] = useState(false)
  const [uninstallConfirmOpen, setUninstallConfirmOpen] = useState(false)

  const organizationName =
    organizationNameOf(plugin.organizationId, organizations ?? []) ?? plugin.organizationId

  const handleSave = () => {
    updatePlugin.mutate({
      slug: plugin.slug,
      change: {
        name: name.trim(),
        description: description.trim() === '' ? null : description.trim(),
        frontend: buildFrontendDto(frontendMode, target),
        permissions: [...permissions],
        requiredPermissions: submittedAccessPermissions(),
      },
    })
  }

  const handleToggleEnabled = (enabled: boolean) => {
    updatePlugin.mutate({ slug: plugin.slug, change: { enabled } })
  }

  const handleRotate = () => {
    rotatePluginKey.mutate(plugin.slug, {
      onSuccess: (response) => {
        setRotateConfirmOpen(false)
        setIssuedKey(response.key)
      },
    })
  }

  const handleUninstall = () => {
    uninstallPlugin.mutate(plugin.slug, {
      onSuccess: () => {
        setUninstallConfirmOpen(false)
        void navigate({ to: '/settings/plugin' })
      },
    })
  }

  return (
    <div className="container mt-6 flex flex-col gap-6 2xl:w-4/5">
      <header className="flex flex-wrap items-start justify-between gap-4">
        <div>
          <h1 className="font-lato font-bold text-3xl mb-2 lg:text-4xl">{plugin.name}</h1>
          <p className="font-mono text-sm text-dark-500">
            {t('plugin.detail.slugLabel')}: {plugin.slug}
          </p>
          <p className="mt-1 text-sm text-dark-600">
            {plugin.createdAt &&
              t('plugin.detail.createdLabel', {
                date: new Intl.DateTimeFormat(intlLocale(i18n.language), {
                  dateStyle: 'medium',
                }).format(plugin.createdAt),
              })}
          </p>
        </div>
        <Badge variant={plugin.enabled ? 'success' : 'muted'}>
          {plugin.enabled ? t('plugin.status.enabled') : t('plugin.status.disabled')}
        </Badge>
      </header>

      <Can permission={['plugin:update']}>
        <div className="flex items-center gap-3">
          <Switch
            checked={plugin.enabled}
            disabled={updatePlugin.isPending}
            onCheckedChange={handleToggleEnabled}
            aria-label={t('plugin.detail.enableLabel')}
          />
          <div>
            <p className="font-medium text-dark">{t('plugin.detail.enableLabel')}</p>
            <p className="text-sm text-dark-600">{t('plugin.detail.enableHint')}</p>
          </div>
        </div>
      </Can>

      <div className="flex items-center gap-3">
        <p className="text-sm text-dark-600">
          {plugin.hasCredential
            ? t('plugin.detail.credentialSet')
            : t('plugin.detail.credentialMissing')}
        </p>
        <Can permission={['plugin:update']}>
          <Button
            type="button"
            variant="outline"
            size="sm"
            onClick={() => setRotateConfirmOpen(true)}
          >
            {t('plugin.detail.rotateButton')}
          </Button>
        </Can>
      </div>

      <p className="text-sm text-dark-600">
        {t('plugin.detail.lastSeenLabel', {
          date: formatLastSeenAt(plugin.lastSeenAt, dateLocale, t),
        })}
      </p>

      <p className="text-sm text-dark-600">
        {t('plugin.install.organizationLabel')}: {organizationName}
      </p>

      <FormField
        id="plugin-detail-name"
        label={t('plugin.install.nameLabel')}
        value={name}
        onChange={(event) => setName(event.target.value)}
        disabled={!canUpdate}
      />

      <TextareaField
        id="plugin-detail-description"
        label={t('plugin.install.descriptionLabel')}
        value={description}
        onChange={(event) => setDescription(event.target.value)}
        disabled={!canUpdate}
      />

      <div className="flex flex-col gap-y-2">
        <Label htmlFor="plugin-detail-frontend-mode">{t('plugin.install.frontendModeLabel')}</Label>
        <select
          id="plugin-detail-frontend-mode"
          value={frontendMode}
          disabled={!canUpdate}
          onChange={(event) => setFrontendMode(event.target.value as FrontendMode)}
          className="flex h-10 w-full max-w-sm rounded-lg border border-dark-200 bg-white px-3 py-2 text-base text-dark-800 shadow-xs outline-none focus-visible:border-green-dark focus-visible:ring-green-dark/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
        >
          <option value="none">{t('plugin.install.frontendModeOption.none')}</option>
          <option value="external">{t('plugin.install.frontendModeOption.external')}</option>
          <option value="proxied">{t('plugin.install.frontendModeOption.proxied')}</option>
        </select>
      </div>

      {frontendMode !== 'none' && (
        <FormField
          id="plugin-detail-target"
          label={t('plugin.install.targetLabel')}
          value={target}
          onChange={(event) => setTarget(event.target.value)}
          disabled={!canUpdate}
          className="max-w-sm"
        />
      )}

      <PluginPermissionMatrix
        heading={t('plugin.install.permissionsHeading')}
        hint={t('plugin.install.permissionsHint')}
        permissions={permissions}
        disabled={!canUpdate}
        onToggle={togglePermission}
      />

      <PluginPermissionMatrix
        heading={t('plugin.install.accessPermissionsHeading')}
        hint={t('plugin.install.accessPermissionsHint')}
        permissions={accessPermissions}
        accessPrefix={t('plugin.permissionMatrix.accessPrefix')}
        visiblePermissions={accessVisiblePermissions(permissions)}
        disabled={!canUpdate}
        onToggle={toggleAccessPermission}
      />

      <div className="flex flex-wrap items-center justify-between gap-3 border-t border-dark-200 pt-4">
        <Can permission={['plugin:delete']}>
          <Button type="button" variant="outline" onClick={() => setUninstallConfirmOpen(true)}>
            {t('plugin.detail.uninstallButton')}
          </Button>
        </Can>

        <Can permission={['plugin:update']}>
          <Button type="button" onClick={handleSave} disabled={updatePlugin.isPending}>
            {t('plugin.detail.saveButton')}
          </Button>
        </Can>
      </div>

      <AlertDialog open={rotateConfirmOpen} onOpenChange={setRotateConfirmOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{t('plugin.detail.rotateConfirmTitle')}</AlertDialogTitle>
            <AlertDialogDescription>
              {t('plugin.detail.rotateConfirmDescription')}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>{t('common:actions.cancel')}</AlertDialogCancel>
            <AlertDialogAction onClick={handleRotate}>
              {t('plugin.detail.rotateButton')}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={uninstallConfirmOpen} onOpenChange={setUninstallConfirmOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{t('plugin.detail.uninstallConfirmTitle')}</AlertDialogTitle>
            <AlertDialogDescription>
              {t('plugin.detail.uninstallConfirmDescription', { name: plugin.name })}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>{t('common:actions.cancel')}</AlertDialogCancel>
            <AlertDialogAction onClick={handleUninstall}>
              {t('plugin.detail.uninstallButton')}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <PluginKeyDialog
        open={issuedKey !== null}
        variant="rotated"
        apiKey={issuedKey ?? ''}
        onOpenChange={(open) => {
          if (!open) setIssuedKey(null)
        }}
      />
    </div>
  )
}

export default PluginDetailPage
