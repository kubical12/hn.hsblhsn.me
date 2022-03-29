import { Skeleton } from 'baseui/skeleton'

export function LoadingScreen() {
  const result = []
  for (let i = 0; i < 2; i++) {
    result.push(
      <Skeleton
        key={i}
        width="100%"
        height="18rem"
        animation
        overrides={{
          Root: {
            style: {
              borderRadius: '0.5rem',
              marginBottom: '1rem',
            },
          },
        }}
      />
    )
  }
  return <>{result}</>
}
