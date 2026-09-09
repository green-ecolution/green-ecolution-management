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
  /**
   * Set for the access-permissions matrix so every checkbox label reads
   * "{accessPrefix}: {area} {action}" — distinguishing it from the plugin's
   * own permission matrix, which shares the same resource/action grid.
   */
  accessPrefix?: string
  /**
   * Restricts which cells render a checkbox at all (the rest show a dash).
   * The access matrix uses this to only offer a write-action toggle for a
   * permission the plugin itself holds — an action button in its view can
   * only exist for a right the plugin actually has — while "read" stays
   * offered everywhere, since viewing data doesn't depend on that. Omitted
   * for the plugin's own matrix, which always offers every cell.
   */
  visiblePermissions?: ReadonlySet<string>
  disabled?: boolean
  onToggle: (permission: Permission) => void
}

const PluginPermissionMatrix = ({
  heading,
  hint,
  permissions,
  accessPrefix,
  visiblePermissions,
  disabled = false,
  onToggle,
}: PluginPermissionMatrixProps) => {
  const { t } = useTranslation('settings')
  const areas = permissionAreasFor(t)
  const actionShort: Record<(typeof ACTIONS)[number], string> = {
    read: t('plugin.permissionMatrix.actionShort.read'),
    create: t('plugin.permissionMatrix.actionShort.create'),
    update: t('plugin.permissionMatrix.actionShort.update'),
    delete: t('plugin.permissionMatrix.actionShort.delete'),
  }

  return (
    <div className="flex flex-col gap-2">
      <div>
        <p className="font-lato text-sm font-semibold text-dark">{heading}</p>
        {hint && <p className="mt-0.5 text-sm text-dark-600">{hint}</p>}
      </div>
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
                {area.actions.map((action) => {
                  if (visiblePermissions && !visiblePermissions.has(action.permission)) {
                    return (
                      <TableCell key={action.permission} className="text-center text-dark-300">
                        &ndash;
                      </TableCell>
                    )
                  }
                  const label = accessPrefix
                    ? `${accessPrefix}: ${area.label} ${actionShort[action.action]}`
                    : `${area.label} ${actionShort[action.action]}`
                  return (
                    <TableCell key={action.permission} className="text-center">
                      <Checkbox
                        aria-label={label}
                        checked={permissions.has(action.permission)}
                        disabled={disabled}
                        onCheckedChange={() => onToggle(action.permission)}
                      />
                    </TableCell>
                  )
                })}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}

export default PluginPermissionMatrix
