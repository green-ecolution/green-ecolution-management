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
  /** Distinguishes the install-time title from the rotate-time one; the body is identical. */
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
          <DialogDescription>{t('plugin.keyDialog.description')}</DialogDescription>
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
