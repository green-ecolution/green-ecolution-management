import { useState } from 'react'
import { toggleAction } from '@/lib/auth/permissionAreas'
import { RESOURCES, type Permission } from '@/lib/auth/permissions'

/**
 * Which permission cells the access matrix offers a toggle for: "read" is
 * always offered (viewing the plugin's data doesn't depend on what the
 * plugin itself may write), everything else only for a permission the
 * plugin currently holds — an action button in its view can only exist for
 * a right the plugin actually has.
 */
export const accessVisiblePermissions = (permissions: ReadonlySet<string>): ReadonlySet<string> => {
  const visible = new Set(permissions)
  RESOURCES.forEach((resource) => visible.add(`${resource}:read`))
  return visible
}

const samePermissionSet = (a: readonly string[], b: readonly string[]): boolean => {
  const left = new Set(a)
  const right = new Set(b)
  return left.size === right.size && [...left].every((entry) => right.has(entry))
}

/**
 * Tracks the plugin's own permission set and the access (required) permission
 * set together: the access set mirrors the plugin set until it is edited
 * directly, at which point mirroring stops for good. `initialAccess`
 * diverging from `initialPermissions` is read as evidence that a previous
 * edit already ended the mirroring — there is no persisted "touched" flag,
 * so an existing divergence is the only signal available.
 */
export const usePluginPermissionDraft = (
  initialPermissions: readonly string[] = [],
  initialAccess: readonly string[] = [],
) => {
  const [permissions, setPermissions] = useState<ReadonlySet<string>>(
    () => new Set(initialPermissions),
  )
  const [accessPermissions, setAccessPermissions] = useState<ReadonlySet<string>>(
    () => new Set(initialAccess),
  )
  const [accessTouched, setAccessTouched] = useState(
    () => !samePermissionSet(initialPermissions, initialAccess),
  )

  const togglePermission = (permission: Permission) => {
    setPermissions((prev) => toggleAction(prev, permission))
    if (!accessTouched) {
      setAccessPermissions((prev) => toggleAction(prev, permission))
    }
  }

  const toggleAccessPermission = (permission: Permission) => {
    setAccessTouched(true)
    setAccessPermissions((prev) => toggleAction(prev, permission))
  }

  const reset = (nextPermissions: readonly string[] = [], nextAccess: readonly string[] = []) => {
    setPermissions(new Set(nextPermissions))
    setAccessPermissions(new Set(nextAccess))
    setAccessTouched(!samePermissionSet(nextPermissions, nextAccess))
  }

  /** Required permissions to submit: the access set, pruned to currently visible cells. */
  const submittedAccessPermissions = (): string[] => {
    const visible = accessVisiblePermissions(permissions)
    return [...accessPermissions].filter((permission) => visible.has(permission))
  }

  return {
    permissions,
    accessPermissions,
    accessTouched,
    togglePermission,
    toggleAccessPermission,
    reset,
    submittedAccessPermissions,
  }
}
