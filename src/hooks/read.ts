import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { ErrorT, FeedItemT } from '../types'
import { ENDPOINTS } from './endpoints'

function useRead() {
  const routeParams = useParams()
  const [data, setData] = useState<FeedItemT | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<ErrorT | null>(null)

  const fetchData = async (url: string) => {
    setLoading(true)
    try {
      const response = await fetch(url)
      const payload = await response.json()
      setData(payload)
    } catch (error) {
      setError({
        message: String(error),
      })
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (routeParams.id === undefined) {
      return
    }
    const url = ENDPOINTS.feedItem(Number(routeParams.id))
    fetchData(url)
  }, [routeParams.id])

  return { data, loading, error }
}

export default useRead
