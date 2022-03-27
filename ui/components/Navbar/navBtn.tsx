import { Button, SHAPE, SIZE, KIND } from 'baseui/button'

export type NavBtnPropsT = {
  onClick?: () => void
  children: JSX.Element
}

export function NavBtn(props: NavBtnPropsT) {
  return (
    <Button
      shape={SHAPE.circle}
      size={SIZE.compact}
      kind={KIND.secondary}
      disabled={props.onClick === undefined}
      onClick={props.onClick}
    >
      {props.children}
    </Button>
  )
}
