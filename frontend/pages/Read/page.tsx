import { Block } from 'baseui/block'
import { Container } from '../../components/Layout'
import { ErrorScreen } from '../../components/ErrorScreen'
import { ReaderContainer, ReaderSkeleton } from '../../components/Reader'
import useRead from '../../hooks/read'
import { useStyletron } from 'baseui'
import { Navbar } from '../../components/Navbar/navbar'
import { Meta } from '../../components/SEO'
import { Fragment } from 'react'

export function Page() {
  const [, theme] = useStyletron()
  const { loading, data, error } = useRead()
  const onBack = () => {
    window.history.back()
  }

  let content: JSX.Element | undefined = undefined
  if (loading) {
    content = <ReaderSkeleton />
  } else if (error) {
    content = <ErrorScreen error={error} />
  } else if (data) {
    content = (
      <Fragment>
        <Meta data={data} /> <ReaderContainer item={data} />
      </Fragment>
    )
  }

  const paddedContent = {
    paddingTop: theme.sizing.scale1600,
    paddingBottom: theme.sizing.scale1600,
    paddingRight: theme.sizing.scale700,
    paddingLeft: theme.sizing.scale700,
  }

  return (
    <Container
      top={<Navbar isLoading={loading} onBack={onBack} />}
      left={<Block></Block>}
      center={<Block $style={paddedContent}>{content}</Block>}
      right={<Block></Block>}
    />
  )
}
