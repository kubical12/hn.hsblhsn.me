import { Block } from 'baseui/block'
import { Container } from '../../components/Layout'
import { ErrorScreen } from '../../components/ErrorScreen'
import { useStyletron } from 'baseui'
import { Navbar } from '../../components/Navbar/navbar'

export function Page() {
  const [, theme] = useStyletron()
  const onBack = () => {
    window.history.back()
  }
  const error = {
    error: true,
    message: 'Page not found (404)',
  }
  const paddedContent = {
    paddingTop: theme.sizing.scale1600,
    paddingBottom: theme.sizing.scale1600,
    paddingRight: theme.sizing.scale700,
    paddingLeft: theme.sizing.scale700,
  }
  return (
    <Container
      top={<Navbar isLoading={false} onBack={onBack} />}
      left={<Block></Block>}
      center={
        <Block $style={paddedContent}>{<ErrorScreen error={error} />}</Block>
      }
      right={<Block></Block>}
    />
  )
}
