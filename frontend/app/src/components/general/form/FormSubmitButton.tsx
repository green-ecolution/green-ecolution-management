import { Button } from '@green-ecolution/ui'
import { MoveRight } from 'lucide-react'
import { useTranslation } from 'react-i18next'

interface FormSubmitButtonProps {
  /** Only for blockers the form explains elsewhere. Never gate this on form
   * validity: a greyed-out button leaves the user without a reason. */
  disabled?: boolean
  className?: string
}

/** Passing `className` replaces the default layout classes instead of merging with them. */
const FormSubmitButton = ({
  disabled = false,
  className = 'mt-10 lg:col-span-full lg:w-fit',
}: FormSubmitButtonProps) => {
  const { t } = useTranslation('common')

  return (
    <Button type="submit" className={className} disabled={disabled}>
      {t('actions.save')}
      <MoveRight className="icon-arrow-animate" />
    </Button>
  )
}

export default FormSubmitButton
