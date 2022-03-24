import { useEffect, useState } from 'react'
import { useParams, useSearchParams } from 'react-router-dom'
import useAppNavigator, { PAGE, FEED_KIND } from './navigation'

type FeedStatusT = {
  kind: FEED_KIND | null;
  page: number | null;
};

const useFeedStatus = (): FeedStatusT | null => {
  const appNav = useAppNavigator()
  const [searchParams] = useSearchParams()
  const routeParams = useParams()
  const [feedKind, setFeedKind] = useState<FEED_KIND | null>(null)
  const [feedPage, setFeedPage] = useState<number | null>(null)

  useEffect(() => {
    if (routeParams.kind !== undefined) {
      const kind = parseFeedKind(routeParams.kind)
      setFeedKind(kind)
    } else {
      setFeedKind(null)
    }
  }, [routeParams.kind])

  useEffect(() => {
    const param = searchParams.get('page')
    if (param !== null) {
      const page: number = parseInt(param, 10)
      if (!isNaN(page)) {
        setFeedPage(Math.max(1, page))
      } else {
        setFeedPage(null)
      }
    } else {
      setFeedPage(null)
    }
  }, [searchParams])

  if (appNav.currentPage() !== PAGE.feed) {
    return null
  }
  return {
    kind: feedKind,
    page: feedPage,
  }
}

function parseFeedKind(str: string): FEED_KIND | null {
  switch (str) {
  case 'new':
    return FEED_KIND.new
  case 'top':
    return FEED_KIND.top
  default:
    return null
  }
}

export default useFeedStatus
