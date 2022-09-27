import { Route, Routes as ReactRouterRoutes } from 'react-router-dom'
import HomePage from './HomePage'
import { lazy, Suspense } from 'react'
import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { Spinner } from 'baseui/spinner'

const ItemPage = lazy(() => import('./ItemPage'))
const SearchPage = lazy(() => import('./SearchPage'))
const NotFoundPage = lazy(() => import('./NotFoundPage'))

const Routes = () => (
  <Suspense fallback={<Fallback />}>
    <ReactRouterRoutes>
      <Route path="/" element={<HomePage />} />
      <Route path="/newest" element={<HomePage />} />
      <Route path="/ask" element={<HomePage />} />
      <Route path="/show" element={<HomePage />} />
      <Route path="/jobs" element={<HomePage />} />
      <Route path="/item" element={<ItemPage />} />
      <Route path="/search" element={<SearchPage />} />
      <Route path="*" element={<NotFoundPage />} />
    </ReactRouterRoutes>
  </Suspense>
)

const Fallback = () => {
  const [, theme] = useStyletron()
  return (
    <Block
      $style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        marginBottom: '100vh',
      }}
    >
      <Block>
        <Spinner
          $style={{
            width: '2rem',
            height: '2rem',
            marginTop: '4rem',
          }}
          $size="small"
          $color={theme.colors.contentTertiary}
        />
      </Block>
    </Block>
  )
}

export { Routes }
