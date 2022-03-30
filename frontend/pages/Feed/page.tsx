import { Block } from 'baseui/block'
import { FeedListContainer, FeedListSkeleton } from '../../components/Feed'
import { Container } from '../../components/Layout'
import useFeed from '../../hooks/feed'
import { ErrorScreen } from '../../components/ErrorScreen'
import { useStyletron } from 'baseui'
import { Navbar } from '../../components/Navbar'
import useFeedStatus from '../../hooks/feedStatus'
import useAppNavigator from '../../hooks/navigation'
import { Fragment, useCallback } from 'react'
import { Meta } from '../../components/SEO'

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
      <Fragment>
        <Meta data={data} />
        <FeedListContainer
          feed={data}
          onBack={page && page > 1 ? onBack : undefined}
          onForward={onForward}
        />
      </Fragment>
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
