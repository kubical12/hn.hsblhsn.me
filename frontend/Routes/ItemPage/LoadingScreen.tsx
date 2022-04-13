import { Fragment } from 'react'
import { Skeleton } from 'baseui/skeleton'
import { StyleProps } from '../../types'

const LoadingScreen: React.FC = () => (
  <Fragment>
    <Skeleton
      width="100%"
      height="100px"
      animation
      overrides={{
        Root: {
          style: ({ $theme }: StyleProps) => ({
            marginTop: $theme.sizing.scale600,
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
          style: ({ $theme }: StyleProps) => ({
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
          style: ({ $theme }: StyleProps) => ({
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

export { LoadingScreen }
