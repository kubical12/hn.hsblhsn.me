import { Button, KIND } from 'baseui/button'
import { ConnectionT, Job, Story, StyleProps } from '../../Types'
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
      <Button
        kind={KIND.secondary}
        onClick={loadNext}
        isLoading={loading}
        disabled={loading || items.pageInfo?.hasNextPage === false}
        overrides={{
          Root: {
            style: ({ $theme }: StyleProps) => ({
              width: '100%',
              marginTop: $theme.sizing.scale900,
            }),
          },
        }}
      >
        {items.pageInfo?.hasNextPage ? 'Load more' : 'No more items'}
      </Button>
    </div>
  )
}

export { PaginatedItemCardList }
export type { PaginatedItemCardListProps }
