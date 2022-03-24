import { Block } from 'baseui/block'
import { FeedList } from '../../components/Feed'
import { Container } from '../../components/Layout'
import useFeed from '../../hooks/feed'
import { ErrorScreen } from '../../components/ErrorScreen'
import { LoadingScreen } from './loading'

export function FeedPage() {
  const { loading, data, error } = useFeed()
  let content: JSX.Element | null = null

  if (loading) {
    content = <LoadingScreen />
  } else if (error) {
    content = <ErrorScreen error={error} />
  } else if (data) {
    content = <FeedList feed={data} />
  }

  return (
    <Container
      left={<Block></Block>}
      center={
        <Block
          $style={{
            paddingTop: '1.5rem',
            paddingBottom: '1.5rem',
          }}
        >
          {content}
        </Block>
      }
      right={<Block></Block>}
    />
  )
}
