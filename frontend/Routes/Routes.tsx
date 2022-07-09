import { Route, Routes as ReactRouterRoutes } from 'react-router-dom'
import HomePage from './HomePage'
import { lazy, Suspense } from 'react'
import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'

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
  const [css, theme] = useStyletron()
  return (
    <Block className="animate-pulse">
      <Block
        className={css({
          textAlign: 'center',
          height: theme.sizing.scale100,
          width: '100%',
          marginTop: '-1px',
          backgroundColor: theme.colors.backgroundAccent,
        })}
      ></Block>
    </Block>
  )
}

export { Routes }
