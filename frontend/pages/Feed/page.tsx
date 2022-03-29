import { Block } from 'baseui/block'
import { FeedListContainer, FeedListSkeleton } from '../../components/Feed'
import { Container } from '../../components/Layout'
import useFeed from '../../hooks/feed'
import { ErrorScreen } from '../../components/ErrorScreen'
import { useStyletron } from 'baseui'
import { Navbar } from '../../components/Navbar/navbar'
import useFeedStatus from '../../hooks/feedStatus'
import useAppNavigator from '../../hooks/navigation'
import { useCallback } from 'react'

export function Page() {
  const [, theme] = useStyletron()
  const { loading, data, error } = useFeed()
  const { kind, page } = useFeedStatus()
  const { feed } = useAppNavigator()

  const paginate = useCallback(
    (inc: number) => {
      if (kind && page) {
        feed(kind, page + inc)
      }
    },
    [feed, kind, page]
  )

  const onBack = () => {
    paginate(-1)
  }
  const onForward = () => {
    paginate(1)
  }

  let content: JSX.Element | undefined = undefined

  if (loading) {
    content = <FeedListSkeleton />
  } else if (error) {
    content = <ErrorScreen error={error} />
  } else if (data) {
    content = (
      <FeedListContainer
        feed={data}
        onBack={page && page > 1 ? onBack : undefined}
        onForward={onForward}
      />
    )
  }

  return (
    <Container
      top={
        <Navbar
          isLoading={loading}
          onBack={page && page > 1 ? onBack : undefined}
          onForward={onForward}
        />
      }
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
