import { FeedT } from '../../types'
import { Fragment, useCallback } from 'react'
import { FeedItem } from './feedItem'
import { FlexGrid, FlexGridItem } from 'baseui/flex-grid'
import { Button, KIND } from 'baseui/button'
import { useStyletron } from 'baseui'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import useAppNavigator, { FEED_KIND } from '../../hooks/navigation'
import useFeedStatus from '../../hooks/feedStatus'
import { ErrorScreen } from '../ErrorScreen'

type FeedProps = {
  feed: FeedT
}

// FeedList renders a list of feedItems.
export function FeedList({ feed }: FeedProps) {
  if (!feed) {
    return (
      <ErrorScreen
        error={{
          message: 'No feed item found on this page.',
          reason: 'feed is null',
        }}
      />
    )
  }
  if (!feed.feedItems || feed.feedItems.length === 0) {
    return (
      <ErrorScreen
        error={{
          message: 'No feed item found on this page.',
          reason: 'feed is empty',
        }}
      />
    )
  }
  return (
    <Fragment>
      {feed.feedItems.map((feedItem, index) => (
        <FeedItem key={index} feedItem={feedItem} />
      ))}
      <FeedListPaginator />
    </Fragment>
  )
}

// FeedListPaginator returns two buttons to move between pages.
function FeedListPaginator() {
  const feedStatus = useFeedStatus()
  const appNav = useAppNavigator()
  const [css] = useStyletron()

  const paginate = useCallback(
    (index: number) => {
      return () => {
        const kind = feedStatus?.kind as FEED_KIND
        const page = feedStatus?.page as number
        appNav.feed(kind, page + index)
      }
    },
    [feedStatus?.kind, feedStatus?.page]
  )

  const paginatorBtn = {
    kind: KIND.secondary,
    className: css({
      width: '100%',
    }),
  }

  return (
    <FlexGrid
      flexGridColumnCount={[1, 1, 2, 2]}
      flexGridColumnGap="scale1000"
      flexGridRowGap="scale500"
    >
      <FlexGridItem>
        <Button
          startEnhancer={<ChevronLeft />}
          onClick={paginate(-1)}
          {...paginatorBtn}
          disabled={(feedStatus?.page as number) <= 1}
        >
          Prev.
        </Button>
      </FlexGridItem>
      <FlexGridItem>
        <Button
          endEnhancer={<ChevronRight />}
          onClick={paginate(1)}
          {...paginatorBtn}
        >
          Next
        </Button>
      </FlexGridItem>
    </FlexGrid>
  )
}
