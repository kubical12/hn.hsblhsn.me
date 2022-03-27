import { useEffect, useState } from 'react'
import { useParams, useSearchParams } from 'react-router-dom'
import useAppNavigator, { PAGE, FEED_KIND } from './navigation'

type FeedStatusT = {
  kind: FEED_KIND | undefined
  page: number | undefined
}

const useFeedStatus = (): FeedStatusT | undefined => {
  const appNav = useAppNavigator()
  const [searchParams] = useSearchParams()
  const routeParams = useParams()
  const [feedKind, setFeedKind] = useState<FEED_KIND | undefined>(undefined)
  const [feedPage, setFeedPage] = useState<number | undefined>(undefined)

  useEffect(() => {
    if (routeParams.kind !== undefined) {
      const kind = parseFeedKind(routeParams.kind)
      setFeedKind(kind)
    } else {
      setFeedKind(undefined)
    }
  }, [routeParams.kind])

  useEffect(() => {
    const param = searchParams.get('page')
    if (param) {
      const page: number = parseInt(param, 10)
      if (!isNaN(page)) {
        setFeedPage(Math.max(1, page))
      } else {
        setFeedPage(undefined)
      }
    } else {
      setFeedPage(undefined)
    }
  }, [searchParams])

  if (appNav.currentPage() !== PAGE.feed) {
    return undefined
  }
  return {
    kind: feedKind,
    page: feedPage,
  }
}

function parseFeedKind(str: string): FEED_KIND | undefined {
  switch (str) {
    case 'new':
      return FEED_KIND.new
    case 'top':
      return FEED_KIND.top
    default:
      return undefined
  }
}

export default useFeedStatus
