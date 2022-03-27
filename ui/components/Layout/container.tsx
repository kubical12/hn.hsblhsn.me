import { Grid, Cell, BEHAVIOR } from 'baseui/layout-grid'
import { Fragment } from 'react'

type ContainerProps = {
  top: JSX.Element
  left: JSX.Element
  center: JSX.Element
  right: JSX.Element
}

export function Container(props: ContainerProps) {
  return (
    <Fragment>
      {props.top}
      <Grid
        behavior={BEHAVIOR.fixed}
        gridMargins={0}
        gridGaps={0}
        gridMaxWidth={1500}
      >
        <Cell span={[0, 1, 3]}>{props.left}</Cell>
        <Cell span={[4, 6, 6]}>{props.center}</Cell>
        <Cell span={[0, 1, 3]}>{props.right}</Cell>
      </Grid>
    </Fragment>
  )
}
