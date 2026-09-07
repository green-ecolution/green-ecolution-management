import { useTranslation } from 'react-i18next'
import { useMediaQuery } from '@/hooks/useMediaQuery'
import ChapterToc from './ChapterToc'
import HandbookSearch from './HandbookSearch'
import type { ChapterSection } from '@/lib/handbook/types'

function ChapterAside({ sections }: { sections: ChapterSection[] }) {
  const { t } = useTranslation('help')
  const beside = useMediaQuery('(min-width: 1024px)')

  if (beside) {
    return (
      <aside className="col-start-2 row-start-1 sticky top-24 self-start space-y-6">
        <HandbookSearch variant="sidebar" />
        <ChapterToc sections={sections} />
      </aside>
    )
  }

  // Searching is the reason most readers come back, so it stays visible while
  // only the section list folds away.
  return (
    <div className="mb-8 space-y-4">
      <HandbookSearch variant="sidebar" />
      {sections.length > 0 && (
        <details className="rounded-xl border border-dark-100">
          <summary className="cursor-pointer px-4 py-3 font-lato font-semibold">
            {t('chapter.onThisPage')}
          </summary>
          <div className="px-4 pb-4">
            <ChapterToc sections={sections} withHeading={false} />
          </div>
        </details>
      )}
    </div>
  )
}

export default ChapterAside
