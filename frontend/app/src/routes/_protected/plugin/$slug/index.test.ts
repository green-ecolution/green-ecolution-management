import { describe, it, expect, vi, beforeEach } from 'vitest'
import { isRedirect } from '@tanstack/react-router'
import { QueryClient } from '@tanstack/react-query'
import type { PluginResponse, UserResponse } from '@green-ecolution/backend-client'

const readAuthBypass = vi.fn(() => false)

vi.mock('@/lib/auth/runtimeConfig', () => ({
  readAuthBypass: () => readAuthBypass(),
}))

const { Route } = await import('./index')

const me = (permissions: string[]): UserResponse =>
  ({
    roles: [
      {
        id: 'role-1',
        name: 'Rolle',
        description: '',
        organizationId: null,
        permissions,
        createdAt: '2026-07-27T00:00:00Z',
      },
    ],
  }) as unknown as UserResponse

const plugin = (overrides: Partial<PluginResponse> = {}): PluginResponse => ({
  id: 'p1',
  slug: 'acme',
  name: 'Acme',
  organizationId: 'org-1',
  permissions: [],
  requiredPermissions: ['tree:read'],
  frontendMode: 'external',
  frontendTarget: 'https://plugin.example.com',
  enabled: true,
  hasCredential: true,
  ...overrides,
})

interface BeforeLoadOpts {
  context: { queryClient: QueryClient }
  params: { slug: string }
}
const beforeLoad = (opts: BeforeLoadOpts) =>
  (Route.options.beforeLoad as (o: BeforeLoadOpts) => Promise<void>)(opts)

const contextFor = (pluginResponse: PluginResponse, permissions: string[]): BeforeLoadOpts => {
  const queryClient = new QueryClient({ defaultOptions: { queries: { retry: false } } })
  queryClient.setQueryData(['plugins', pluginResponse.slug], pluginResponse)
  queryClient.setQueryData(['users', 'me'], me(permissions))
  return { context: { queryClient }, params: { slug: pluginResponse.slug } }
}

describe('/plugin/$slug beforeLoad', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    readAuthBypass.mockReturnValue(false)
  })

  it('resolves when the user holds the required permissions', async () => {
    await expect(beforeLoad(contextFor(plugin(), ['tree:read']))).resolves.toBeUndefined()
  })

  it('redirects to the plugin list when the user lacks the required permissions', async () => {
    const thrown = await beforeLoad(contextFor(plugin(), ['vehicle:read'])).catch(
      (error: unknown) => error,
    )
    expect(isRedirect(thrown)).toBe(true)
    expect((thrown as { options: { to?: string } }).options.to).toBe('/settings/plugin')
  })

  it('redirects a proxied plugin the same way when permissions are missing', async () => {
    const thrown = await beforeLoad(
      contextFor(plugin({ frontendMode: 'proxied', frontendTarget: 'plugin-backend:8080' }), [
        'vehicle:read',
      ]),
    ).catch((error: unknown) => error)
    expect(isRedirect(thrown)).toBe(true)
  })

  it('resolves for a proxied plugin when permissions are held, leaving the view kind to the component', async () => {
    await expect(
      beforeLoad(
        contextFor(plugin({ frontendMode: 'proxied', frontendTarget: 'plugin-backend:8080' }), [
          'tree:read',
        ]),
      ),
    ).resolves.toBeUndefined()
  })

  it('resolves without consulting the user cache when auth is bypassed', async () => {
    readAuthBypass.mockReturnValue(true)
    const queryClient = new QueryClient()
    queryClient.setQueryData(['plugins', 'acme'], plugin())
    const spy = vi.spyOn(queryClient, 'ensureQueryData')

    await expect(
      beforeLoad({ context: { queryClient }, params: { slug: 'acme' } }),
    ).resolves.toBeUndefined()

    expect(spy).toHaveBeenCalledTimes(1)
    expect(spy.mock.calls[0][0]).toMatchObject({ queryKey: ['plugins', 'acme'] })
  })
})
