import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'
import type { PluginCreateRequest, PluginUpdateRequest } from '@/api/backendApi'
import { pluginApi } from '@/api/backendApi'
import createToast from '@/hooks/createToast'

export interface UpdatePluginVariables {
  slug: string
  change: PluginUpdateRequest
}

export const usePluginMutations = () => {
  const queryClient = useQueryClient()
  const showToast = createToast()
  const { t } = useTranslation('settings')

  const invalidate = () => {
    void queryClient.invalidateQueries({ queryKey: ['plugins'] })
  }

  const installPlugin = useMutation({
    mutationFn: (pluginCreateRequest: PluginCreateRequest) =>
      pluginApi.installPlugin({ pluginCreateRequest }),
    onSuccess: () => {
      invalidate()
      showToast(t('plugin.detail.toast.installed'))
    },
    onError: () => showToast(t('plugin.detail.toast.installFailed'), 'error'),
  })

  const updatePlugin = useMutation({
    mutationFn: ({ slug, change }: UpdatePluginVariables) =>
      pluginApi.updatePlugin({ pluginSlug: slug, pluginUpdateRequest: change }),
    onSuccess: () => {
      invalidate()
      showToast(t('plugin.detail.toast.saved'))
    },
    onError: () => showToast(t('plugin.detail.toast.saveFailed'), 'error'),
  })

  const rotatePluginKey = useMutation({
    mutationFn: (slug: string) => pluginApi.rotatePluginKey({ pluginSlug: slug }),
    onSuccess: () => {
      invalidate()
      showToast(t('plugin.detail.toast.rotated'))
    },
    onError: () => showToast(t('plugin.detail.toast.rotateFailed'), 'error'),
  })

  const uninstallPlugin = useMutation({
    mutationFn: (slug: string) => pluginApi.uninstallPlugin({ pluginSlug: slug }),
    onSuccess: () => {
      invalidate()
      showToast(t('plugin.detail.toast.uninstalled'))
    },
    onError: () => showToast(t('plugin.detail.toast.uninstallFailed'), 'error'),
  })

  return { installPlugin, updatePlugin, rotatePluginKey, uninstallPlugin }
}
