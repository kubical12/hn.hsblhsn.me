import { useEffect } from 'react'
import useAppNavigator, { FEED_KIND } from '../../hooks/navigation'

export function HomePage() {
  const navigate = useAppNavigator()
  useEffect(() => {
    navigate.feed(FEED_KIND.top, 1)
  }, [])
  return null
}
