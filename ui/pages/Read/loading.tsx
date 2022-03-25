import { Skeleton } from 'baseui/skeleton'
import { Fragment } from 'react'

export function LoadingScreen() {
  return (
    <Fragment>
      <Skeleton
        width="100%"
        height="100px"
        animation
        overrides={{
          Root: {
            style: ({ $theme }) => ({
              marginBottom: $theme.sizing.scale1000,
            }),
          },
        }}
      />
      <Skeleton
        animation={true}
        rows={3}
        width="100%"
        overrides={{
          Root: {
            style: ({ $theme }) => ({
              marginBottom: $theme.sizing.scale1200,
            }),
          },
          Row: {
            style: {
              height: '20px',
              marginBottom: '15px',
            },
          },
        }}
      />
      <Skeleton
        animation={true}
        rows={3}
        width="100%"
        overrides={{
          Root: {
            style: ({ $theme }) => ({
              marginBottom: $theme.sizing.scale600,
            }),
          },
          Row: {
            style: {
              height: '20px',
              marginBottom: '15px',
            },
          },
        }}
      />
    </Fragment>
  )
}
