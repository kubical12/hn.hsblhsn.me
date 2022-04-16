import { Button, KIND } from 'baseui/button'
import { Accordion, Panel } from 'baseui/accordion'
import { Block } from 'baseui/block'
import { HeadingXXLarge, ParagraphLarge } from 'baseui/typography'
import { useStyletron, styled } from 'baseui'
import { Cell, Grid } from 'baseui/layout-grid'
import { useCallback } from 'react'
import { ApolloError } from '@apollo/client'

const MonoSpaced = styled('pre', {
  fontFamily: 'monospace',
  textAlign: 'left',
})

export function ErrorScreen({ error }: { error: ApolloError | unknown }) {
  const [css, theme] = useStyletron()
  const reloadThisPage = useCallback(() => {
    window.location.reload()
  }, [])
  const goToHomePage = useCallback(() => {
    window.location.href = '/'
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
          marginTop: theme.sizing.scale600,
          marginBottom: theme.sizing.scale1000,
        })}
      >
        <ParagraphLarge>An error occurred!</ParagraphLarge>
        <ParagraphLarge>
          We have received the details and will be looking into it.
        </ParagraphLarge>
        <ParagraphLarge>Sorry for the inconvenience!</ParagraphLarge>
        <br />
        <br />
        <br />
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
          <Button kind={KIND.tertiary} onClick={goToHomePage}>
            Go To Homepage
          </Button>
        </Cell>
      </Grid>
    </Block>
  )
}
