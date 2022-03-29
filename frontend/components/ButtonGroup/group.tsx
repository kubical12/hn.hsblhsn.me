import { Button, SIZE, KIND } from 'baseui/button'
import { FlexGrid, FlexGridItem } from 'baseui/flex-grid'
import { useStyletron } from 'baseui'

type LeftRightButtonsPropsT = {
  onLeft: (() => void) | undefined
  onRight: (() => void) | undefined
  leftContent: React.ReactNode
  rightContent: React.ReactNode
}

export function LeftRightButtons(props: LeftRightButtonsPropsT) {
  const [, theme] = useStyletron()
  const btnProps = {
    kind: KIND.secondary,
    size: SIZE.default,
    overrides: {
      BaseButton: {
        style: {
          width: '100%',
        },
      },
    },
  }
  return (
    <FlexGrid
      flexGridColumnCount={[2, 2, 2, 2]}
      flexGridColumnGap={theme.sizing.scale1000}
    >
      <FlexGridItem>
        <Button onClick={props.onLeft} disabled={!props.onLeft} {...btnProps}>
          {props.leftContent}
        </Button>
      </FlexGridItem>
      <FlexGridItem>
        <Button onClick={props.onRight} disabled={!props.onRight} {...btnProps}>
          {props.rightContent}
        </Button>
      </FlexGridItem>
    </FlexGrid>
  )
}
