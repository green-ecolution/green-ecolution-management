import { describe, expect, it } from 'vitest'
import { isSearchableQuery, searchHandbook, splitOnMatch } from './search'
import type { SearchEntry } from './types'

const entries: SearchEntry[] = [
  {
    slug: 'watering-plans',
    anchor: 'route-festlegen',
    sectionTitle: 'Route festlegen',
    text: 'Erst die Gruppen wählen, dann die Route berechnen lassen.',
  },
  {
    slug: 'sensors',
    anchor: 'zustande',
    sectionTitle: 'Zustände',
    text: 'Ein Sensor ist vorbereitet, online oder offline.',
  },
]

describe('searchHandbook', () => {
  it('returns nothing for a query shorter than two characters', () => {
    expect(searchHandbook(entries, 'r')).toEqual([])
  })

  it('finds a match in the section text regardless of case', () => {
    const hits = searchHandbook(entries, 'ROUTE')

    expect(hits).toHaveLength(1)
    expect(hits[0]).toMatchObject({ slug: 'watering-plans', anchor: 'route-festlegen' })
  })

  it('finds a match in the section title', () => {
    expect(searchHandbook(entries, 'zustände')[0].slug).toBe('sensors')
  })

  it('cuts an excerpt around the match', () => {
    expect(searchHandbook(entries, 'berechnen')[0].excerpt).toContain('berechnen')
  })

  it('ranks a title match above a body match', () => {
    const hits = searchHandbook(entries, 'route')

    expect(hits[0].sectionTitle).toBe('Route festlegen')
  })

  it('leaves the excerpt empty for a title match on a section without paragraph text', () => {
    const withoutText: SearchEntry[] = [
      { slug: 'glossary', anchor: 'begriffe', sectionTitle: 'Begriffe', text: '' },
    ]

    const hits = searchHandbook(withoutText, 'begriffe')

    expect(hits).toHaveLength(1)
    expect(hits[0].excerpt).toBe('')
  })
})

describe('isSearchableQuery', () => {
  it('rejects a whitespace-only query', () => {
    expect(isSearchableQuery('   ')).toBe(false)
  })

  it('rejects a one-character query', () => {
    expect(isSearchableQuery('r')).toBe(false)
  })

  it('accepts a two-character query', () => {
    expect(isSearchableQuery('ro')).toBe(true)
  })

  it('trims surrounding whitespace before checking the length', () => {
    expect(isSearchableQuery('  a  ')).toBe(false)
    expect(isSearchableQuery('  ro  ')).toBe(true)
  })
})

describe('splitOnMatch', () => {
  it('splits around the match and keeps the original casing', () => {
    expect(splitOnMatch('Erst die Route berechnen', 'route')).toEqual({
      before: 'Erst die ',
      match: 'Route',
      after: ' berechnen',
    })
  })

  it('leaves nothing before a match at the start', () => {
    expect(splitOnMatch('Route festlegen', 'Route')).toEqual({
      before: '',
      match: 'Route',
      after: ' festlegen',
    })
  })

  it('splits on the first occurrence only', () => {
    expect(splitOnMatch('Route und Route', 'route')?.after).toBe(' und Route')
  })

  it('returns nothing when the text does not contain the query', () => {
    expect(splitOnMatch('Route festlegen', 'sensor')).toBeNull()
  })

  it('returns nothing for a query too short to search with', () => {
    expect(splitOnMatch('Route festlegen', 'r')).toBeNull()
  })

  it('ignores whitespace around the query', () => {
    expect(splitOnMatch('Route festlegen', '  route  ')?.match).toBe('Route')
  })
})
