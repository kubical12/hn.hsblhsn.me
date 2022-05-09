import { styled } from 'baseui'
import { Block } from 'baseui/block'
import { ConnectionT, Job, Story } from '../../Types'
import { ItemCard } from '../ItemCard'
import { AdWindow, FeedAd } from '../GoogleAds'
import { useContext } from 'react'
import { ConfigContext } from '../Config'

const StyledItemCardListItem = styled(Block, ({ $theme }) => ({
  marginTop: $theme.sizing.scale900,
}))

interface ItemCardListProps {
  items: ConnectionT<Story | Job>
}

const ItemCardList: React.FC<ItemCardListProps> = ({
  items,
}: ItemCardListProps) => {
  const config = useContext(ConfigContext)
  return (
    <Block>
      {items.edges.map((edge, index) => {
        const item = edge.node
        if (!item) {
          // eslint-disable-next-line unicorn/no-null
          return null
        }

        // eslint-disable-next-line unicorn/no-null
        let ad = null
        const shouldShowAd =
          config.ads.enabled &&
          index != 0 &&
          index % config.ads.frequency === 0 &&
          (window as AdWindow)?.adsbygoogle
        if (shouldShowAd && config.ads.google) {
          ad = (
            <FeedAd
              layoutKey={config.ads.google?.adLayout}
              client={config.ads.google.adClient}
              slot={config.ads.google.feedAdSlot}
              key={index}
            />
          )
        }
        return (
          <StyledItemCardListItem key={index}>
            {ad}
            <ItemCard item={item} />
          </StyledItemCardListItem>
        )
      })}
    </Block>
  )
}

export { ItemCardList }
export type { ItemCardListProps }
