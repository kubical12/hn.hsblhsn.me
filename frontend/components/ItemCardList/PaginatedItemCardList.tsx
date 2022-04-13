import { Button, KIND } from 'baseui/button'
import { ConnectionT, Job, Story, StyleProps } from '../../types'
import { ItemCardList } from './ItemCardList'

interface PaginatedItemCardListProps {
  items: ConnectionT<Story | Job>
  loading: boolean
  loadNext: () => void
}

const PaginatedItemCardList: React.FC<PaginatedItemCardListProps> = ({
  items,
  loading,
  loadNext,
}: PaginatedItemCardListProps) => {
  return (
    <div>
      <ItemCardList items={items} />
      {items.pageInfo?.hasNextPage && (
        <Button
          kind={KIND.secondary}
          onClick={loadNext}
          isLoading={loading}
          disabled={loading}
          overrides={{
            Root: {
              style: ({ $theme }: StyleProps) => ({
                width: '100%',
                marginTop: $theme.sizing.scale900,
              }),
            },
          }}
        >
          {loading ? 'Loading...' : 'Load more'}
        </Button>
      )}
    </div>
  )
}

export { PaginatedItemCardList }
export type { PaginatedItemCardListProps }
