/* eslint-disable unicorn/no-null */
import { memo } from 'react'

export function createComponent<CProps, UIProps>(
  ui: (props: UIProps) => JSX.Element,
  prelude: (props: CProps) => UIProps | undefined
) {
  const component = (props: CProps) => {
    let uiProps = undefined
    try {
      uiProps = prelude(props)
      if (uiProps) {
        return ui(uiProps)
      }
      return null
    } catch (e) {
      console.error(e)
      return null
    }
  }
  return memo(component)
}
