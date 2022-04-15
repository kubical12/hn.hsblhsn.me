import { styled } from 'baseui'
import { Block } from 'baseui/block'
import { ConnectionT, Job, Story } from '../../Types'
import { ItemCard } from '../ItemCard'

const StyledItemCardListItem = styled(Block, ({ $theme }) => ({
  marginTop: $theme.sizing.scale900,
}))

interface ItemCardListProps {
  items: ConnectionT<Story | Job>
}

const ItemCardList: React.FC<ItemCardListProps> = ({
  items,
}: ItemCardListProps) => {
  return (
    <Block>
      {items.edges.map((edge, index) => {
        const item = edge.node
        if (!item) {
          // eslint-disable-next-line unicorn/no-null
          return null
        }
        return (
          <StyledItemCardListItem key={index}>
            <ItemCard item={item} />
          </StyledItemCardListItem>
        )
      })}
    </Block>
  )
}

export { ItemCardList }
export type { ItemCardListProps }
