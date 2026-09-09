import { useId } from 'react'
import { useTranslation } from 'react-i18next'
import {
  Checkbox,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@green-ecolution/ui'
import { ACTIONS, type Permission } from '@/lib/auth/permissions'
import { permissionAreasFor } from '@/lib/auth/permissionAreas'

interface PluginPermissionMatrixProps {
  heading: string
  hint?: string
  permissions: ReadonlySet<string>
  disabled?: boolean
  onToggle: (permission: Permission) => void
}

/**
 * A `<fieldset>` with a `<legend>` gives the browser a role="group" whose
 * accessible name is the legend text — so the plugin matrix and the access
 * matrix are two distinctly-named groups a screen reader (and a test) can
 * address independently, without repeating a prefix on every cell.
 */
const PluginPermissionMatrix = ({
  heading,
  hint,
  permissions,
  disabled = false,
  onToggle,
}: PluginPermissionMatrixProps) => {
  const { t } = useTranslation('settings')
  const hintId = useId()
  const areas = permissionAreasFor(t)
  const actionShort: Record<(typeof ACTIONS)[number], string> = {
    read: t('plugin.permissionMatrix.actionShort.read'),
    create: t('plugin.permissionMatrix.actionShort.create'),
    update: t('plugin.permissionMatrix.actionShort.update'),
    delete: t('plugin.permissionMatrix.actionShort.delete'),
  }

  return (
    <fieldset aria-describedby={hint ? hintId : undefined} className="flex flex-col gap-2">
      <legend className="font-lato text-sm font-semibold text-dark">{heading}</legend>
      {hint && (
        <p id={hintId} className="-mt-1 text-sm text-dark-600">
          {hint}
        </p>
      )}
      <div className="overflow-x-auto rounded-xl border border-dark-100">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="min-w-40">
                {t('plugin.permissionMatrix.areaColumnLabel')}
              </TableHead>
              {ACTIONS.map((action) => (
                <TableHead key={action} className="text-center">
                  {actionShort[action]}
                </TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            {areas.map((area) => (
              <TableRow key={area.resource}>
                <TableCell className="font-medium text-dark">{area.label}</TableCell>
                {area.actions.map((action) => (
                  <TableCell key={action.permission} className="text-center">
                    <Checkbox
                      aria-label={`${area.label} ${actionShort[action.action]}`}
                      checked={permissions.has(action.permission)}
                      disabled={disabled}
                      onCheckedChange={() => onToggle(action.permission)}
                    />
                  </TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </fieldset>
  )
}

export default PluginPermissionMatrix
