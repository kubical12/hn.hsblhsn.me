import { Block } from 'baseui/block'
import { useStyletron } from 'baseui'
import { BEHAVIOR, Cell, Grid } from 'baseui/layout-grid'
import { LabelXSmall } from 'baseui/typography'

const Footer: React.FC = () => {
  const [, theme] = useStyletron()
  return (
    <Grid
      behavior={BEHAVIOR.fixed}
      gridColumns={[4, 8, 16]}
      gridMargins={[0, 0, 0]}
      gridGaps={0}
      gridMaxWidth={1200}
      overrides={{
        Grid: {
          style: {
            backgroundColor: theme.colors.backgroundTertiary,
            position: 'relative',
            marginLeft: 'auto',
            marginRight: 'auto',
            width: '100%',
            bottom: 0,
            marginTop: theme.sizing.scale1200,
            paddingTop: theme.sizing.scale1200,
            paddingBottom: theme.sizing.scale1200,
          },
        },
      }}
    >
      <Cell span={[0, 1, 3]}></Cell>
      <Cell span={[4, 6, 10]}>
        <Block
          $style={{
            paddingLeft: theme.sizing.scale600,
            paddingRight: theme.sizing.scale600,
          }}
        >
          <Logo />
          <LabelXSmall
            $style={{
              fontSize: theme.sizing.scale500,
              color: theme.colors.contentTertiary,
            }}
          >
            Source code available on{' '}
            <a
              rel="noreferrer"
              target="_blank"
              href="https://github.com/hsblhsn/hn.hsblhsn.me"
              style={{
                textDecoration: 'underline',
              }}
            >
              GitHub
            </a>
          </LabelXSmall>
        </Block>
      </Cell>
      <Cell span={[0, 1, 3]}></Cell>
    </Grid>
  )
}

const Logo: React.FC = () => {
  const [, theme] = useStyletron()
  return (
    <Block
      $style={{
        display: 'flex',
        alignItems: 'items-left',
        justifyContent: 'justify-left',
        color: theme.colors.contentTertiary,
        fontSize: theme.sizing.scale600,
      }}
    >
      <Block
        $style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <span
          className="flex-1 h-6 w-6 text-center mr-2 animate__animated animate__fadeIn"
          style={{
            backgroundColor: theme.colors.mono500,
            color: theme.colors.contentTertiary,
            borderWidth: 1,
            borderStyle: 'solid',
            borderColor: theme.colors.borderSelected,
          }}
        >
          H
        </span>
        <span>HackerNews</span>
      </Block>
    </Block>
  )
}

export { Footer }
