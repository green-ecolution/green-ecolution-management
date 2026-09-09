import { useTranslation } from 'react-i18next'
import {
  Button,
  CopyableText,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@green-ecolution/ui'

interface PluginKeyDialogProps {
  open: boolean
  /** Install-time and rotate-time differ in title, and install adds the activation hint. */
  variant: 'installed' | 'rotated'
  apiKey: string
  onOpenChange: (open: boolean) => void
}

const PluginKeyDialog = ({ open, variant, apiKey, onOpenChange }: PluginKeyDialogProps) => {
  const { t } = useTranslation('settings')

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {variant === 'installed'
              ? t('plugin.keyDialog.installedTitle')
              : t('plugin.keyDialog.rotatedTitle')}
          </DialogTitle>
          <DialogDescription>
            {t('plugin.keyDialog.description')}
            {/* Only on install: save_new writes enabled = false, so the adapter
                is answered with 403 plugin.disabled until an admin enables it.
                Rotation leaves the flag alone, so the hint would be wrong there. */}
            {variant === 'installed' && ` ${t('plugin.keyDialog.activationHint')}`}
          </DialogDescription>
        </DialogHeader>

        <CopyableText label={t('plugin.keyDialog.keyLabel')} value={apiKey} />

        <DialogFooter>
          <Button type="button" onClick={() => onOpenChange(false)}>
            {t('plugin.keyDialog.close')}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export default PluginKeyDialog
