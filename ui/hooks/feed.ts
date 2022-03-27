import { useEffect, useState } from 'react'
import { ErrorT, FeedT } from '../types'
import { ENDPOINTS } from './endpoints'
import useFeedStatus from './feedStatus'

function useFeed() {
  const feedStatus = useFeedStatus()
  const [data, setData] = useState<FeedT | undefined>(undefined)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<ErrorT | undefined>(undefined)

  const fetchData = async (url: string) => {
    setLoading(true)
    try {
      const response = await fetch(url)
      const payload = await response.json()
      if (payload && payload.error) {
        setError(payload)
      } else {
        setData(payload)
      }
    } catch (error) {
      setError({
        message: String(error),
      })
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (!feedStatus || !feedStatus.kind || !feedStatus.page) {
      return
    }
    const url = ENDPOINTS.feedList(feedStatus.kind, feedStatus.page)
    fetchData(url)
  }, [feedStatus?.kind, feedStatus?.page])

  return { data, loading, error }
}

export default useFeed
