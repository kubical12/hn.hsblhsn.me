import { Button, KIND } from 'baseui/button'
import { Accordion, Panel } from 'baseui/accordion'
import { Block } from 'baseui/block'
import { ErrorT } from '../../types'
import { HeadingXXLarge, Paragraph1 } from 'baseui/typography'
import { useStyletron, styled } from 'baseui'
import { Cell, Grid } from 'baseui/layout-grid'
import { useCallback } from 'react'
import useAppNavigator from '../../hooks/navigation'

const MonoSpaced = styled('pre', {
  fontFamily: 'monospace',
  textAlign: 'left',
})

export function ErrorScreen({ error }: { error: ErrorT }) {
  const [css, theme] = useStyletron()
  const appNav = useAppNavigator()

  const reloadThisPage = useCallback(() => {
    window.location.reload()
  }, [])
  const goToHomePage = useCallback(() => {
    appNav.home()
  }, [])

  return (
    <Block
      className={css({
        textAlign: 'center',
      })}
    >
      <HeadingXXLarge>Oops!</HeadingXXLarge>
      <Block
        className={css({
          paddingTop: theme.sizing.scale600,
          paddingBottom: theme.sizing.scale1000,
        })}
      >
        <Paragraph1>An error occurred!</Paragraph1>
        <Paragraph1>
          We have received the details and will be looking into it.
        </Paragraph1>
        <Paragraph1>Sorry for the inconvenience!</Paragraph1>
        <Accordion>
          <Panel title="Error Details">
            <MonoSpaced>{JSON.stringify(error, undefined, 2)}</MonoSpaced>
          </Panel>
        </Accordion>
      </Block>
      <Grid gridColumns={[1, 1, 1]} gridGaps={5} gridMargins={5}>
        <Cell>
          <Button kind={KIND.primary} onClick={reloadThisPage}>
            Reload This Page
          </Button>
        </Cell>
        <Cell>
          <Button kind={KIND.minimal} onClick={goToHomePage}>
            Go To Homepage
          </Button>
        </Cell>
      </Grid>
    </Block>
  )
}
