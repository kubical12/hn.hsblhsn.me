import { useStyletron } from 'baseui'
import { useEffect, useRef } from 'react'
import LoadingBar from 'react-top-loading-bar'

interface LoadingBarRefI {
  staticStart(): void
  continuousStart(): void
  complete(): void
}

export function Loader({ isLoading }: { isLoading: boolean | undefined }) {
  const ref = useRef<LoadingBarRefI>(null)
  const [, theme] = useStyletron()
  useEffect(() => {
    if (isLoading && isLoading === true) {
      ref.current?.continuousStart()
    } else {
      ref.current?.complete()
    }
  }, [isLoading])
  return (
    <LoadingBar
      color={theme.colors.accent}
      height={3}
      ref={ref}
      shadow={true}
    />
  )
}
