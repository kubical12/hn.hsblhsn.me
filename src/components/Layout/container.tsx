import { Grid, Cell } from 'baseui/layout-grid'

type ContainerProps = {
  left: JSX.Element;
  center: JSX.Element;
  right: JSX.Element;
};

export function Container(props: ContainerProps) {
  return (
    <Grid
      gridMargins={0}
      gridGaps={0}
    >
      <Cell span={[0, 1, 3]}>{props.left}</Cell>
      <Cell span={[4, 6, 6]}>{props.center}</Cell>
      <Cell span={[0, 1, 3]}>{props.right}</Cell>
    </Grid>
  )
}
