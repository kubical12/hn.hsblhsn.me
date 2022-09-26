import { ConnectionT, Story } from '../../Types'
import { useStyletron } from 'baseui'
import { PaddedBlock } from '../Layout'
import { ParagraphMedium } from 'baseui/typography'
import { PaginatedItemCardList } from '../ItemCardList'
import { Block } from 'baseui/block'

interface SearchResultsProps {
  query: string
  onLoadMore: () => void
  loading: boolean
  results: ConnectionT<Story>
}

const SearchResults = ({
  query,
  onLoadMore,
  loading,
  results,
}: SearchResultsProps) => {
  const [css, theme] = useStyletron()
  if (query.trim().length === 0) {
    return <Block />
  }

  if (results && results.edges.length === 0) {
    return (
      <PaddedBlock className="animate__animated animate__fadeIn animate__faster">
        <ParagraphMedium
          className={css({
            textAlign: 'center',
            color: theme.colors.contentSecondary,
          })}
        >
          :(
          <br />
          No results found.
          <br />
          Try a different search term.
        </ParagraphMedium>
      </PaddedBlock>
    )
  }
  return (
    <PaginatedItemCardList
      items={results}
      loading={loading}
      loadNext={onLoadMore}
      nextPageUrl={window.location.toString()}
    />
  )
}

export { SearchResults }
