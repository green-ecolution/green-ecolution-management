import { describe, it, expect } from 'vitest'
import { servedByNetwork } from './pwaNavigateFallback'

describe('navigateFallbackDenylist', () => {
  it('leaves the handbook PDF to the network', () => {
    expect(servedByNetwork('/handbook/green-ecolution-handbuch.pdf')).toBe(true)
  })

  it('leaves the backend routes to the network', () => {
    expect(servedByNetwork('/api/config.js')).toBe(true)
    expect(servedByNetwork('/api/v1/tree')).toBe(true)
  })

  it('keeps the SPA routes on the navigation fallback', () => {
    expect(servedByNetwork('/')).toBe(false)
    expect(servedByNetwork('/help')).toBe(false)
    expect(servedByNetwork('/help/dashboard')).toBe(false)
    expect(servedByNetwork('/map')).toBe(false)
  })
})
