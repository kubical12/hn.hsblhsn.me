import { FeedItemT, FeedT } from '../../types'
import { Fragment } from 'react'
import { ItemContainer, ItemSkeleton } from './item'
import { Block } from 'baseui/block'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import { LeftRightButtons } from '../ButtonGroup'
import { createComponent } from '../component'

type ContainerProps = {
  feed: FeedT
  onBack: (() => void) | undefined
  onForward: (() => void) | undefined
}

type UIProps = {
  items: Array<FeedItemT | null>
  onBack: (() => void) | undefined
  onForward: (() => void) | undefined
}

// ui component.
function ui(props: UIProps) {
  const items = props.items.map((item, index) => (
    <ItemContainer key={index} item={item} />
  ))
  return (
    <Fragment>
      {/* feed list */}
      <Block>{items}</Block>
      {/* pagination buttons */}
      <LeftRightButtons
        onLeft={props.onBack}
        onRight={props.onForward}
        leftContent={
          <Fragment>
            <ChevronLeft />
            Prev.
          </Fragment>
        }
        rightContent={
          <Fragment>
            Next <ChevronRight />
          </Fragment>
        }
      />
    </Fragment>
  )
}

// prelude prepares the data for the ui component.
function prelude(props: ContainerProps): UIProps | undefined {
  const { feed, onBack, onForward } = props
  if (!feed) {
    return undefined
  }
  const items = feed.items
  if (!items || items.length === 0) {
    return undefined
  }
  return {
    items: items,
    onBack: onBack,
    onForward: onForward,
  }
}

// ListSkeleton is the skeleton for the feed list.
const ListSkeleton = () => (
  <Fragment>
    <ItemSkeleton />
    <ItemSkeleton />
  </Fragment>
)

const ListContainer = createComponent(ui, prelude)

export { ListContainer, ListSkeleton }
