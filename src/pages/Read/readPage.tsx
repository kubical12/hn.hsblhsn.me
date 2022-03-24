import { Block } from 'baseui/block'
import { Container } from '../../components/Layout'
import { ErrorScreen } from '../../components/ErrorScreen'
import { LoadingScreen } from './loading'
import { ReaderView } from '../../components/ReaderVIew'
import useRead from '../../hooks/read'
import { useStyletron } from 'baseui'

export function FeedPage() {
  const [,theme] = useStyletron()
  const { loading, data, error } = useRead()
  let content: JSX.Element | null = null
  if (loading) {
    content = <LoadingScreen />
  } else if (error) {
    content = <ErrorScreen error={error} />
  } else if (data) {
    content = <ReaderView feedItem={data} font="font-sans" />
  }

  return (
    <Container
      left={<Block></Block>}
      center={
        <Block
          $style={{
            paddingTop: '1.5rem',
            paddingBottom: '1.5rem',
            paddingRight: theme.sizing.scale700,
            paddingLeft: theme.sizing.scale700,
          }}
        >
          {content}
        </Block>
      }
      right={<Block></Block>}
    />
  )
}
