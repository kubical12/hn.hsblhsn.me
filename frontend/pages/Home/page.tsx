import { Container } from '../../components/Layout'
import { useEffect } from 'react'
import { Empty } from '../../components/Empty'
import { Navbar } from '../../components/Navbar/navbar'
import useAppNavigator, { FEED_KIND } from '../../hooks/navigation'

export function Page() {
  const navigate = useAppNavigator()
  useEffect(() => {
    navigate.feed(FEED_KIND.top, 1)
  }, [])
  return (
    <Container
      top={<Navbar />}
      left={<Empty />}
      center={<Empty />}
      right={<Empty />}
    />
  )
}
