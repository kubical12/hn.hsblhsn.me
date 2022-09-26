import { Grid, Cell, BEHAVIOR } from 'baseui/layout-grid'
import { Block } from 'baseui/block'

type ContainerProps = {
  left: JSX.Element
  center: JSX.Element
  right: JSX.Element
}

export function Container(props: ContainerProps) {
  return (
    <Block className="animate__animated animate__fadeIn animate__faster">
      <Grid
        behavior={BEHAVIOR.fixed}
        gridColumns={[4, 8, 16]}
        gridMargins={[0, 0, 0]}
        gridGaps={0}
        gridMaxWidth={1200}
      >
        <Cell span={[0, 1, 3]}>{props.left}</Cell>
        <Cell span={[4, 6, 10]}>{props.center}</Cell>
        <Cell span={[0, 1, 3]}>{props.right}</Cell>
      </Grid>
    </Block>
  )
}
