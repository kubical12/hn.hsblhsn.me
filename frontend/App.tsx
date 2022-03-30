import { BaseProvider, DarkTheme, LightTheme, styled } from 'baseui'
import { lazy, Suspense } from 'react'
import { Route, Routes } from 'react-router-dom'
import { Loader } from './components/Navbar'
import { useDarkMode } from './hooks/theme'

const FeedPage = lazy(() => import('./pages/Feed'))
const HomePage = lazy(() => import('./pages/Home'))
const ReadPage = lazy(() => import('./pages/Read'))
const NotFoundPage = lazy(() => import('./pages/NotFound'))

const StyledAppShell = styled('div', ({ $theme }) => ({
  backgroundColor: $theme.colors.backgroundPrimary,
  color: $theme.colors.contentPrimary,
  overflow: 'hidden',
  minHeight: '100vh',
  minWidth: '100vw',
}))

function App() {
  const isDark = useDarkMode()
  return (
    <BaseProvider theme={isDark ? DarkTheme : LightTheme}>
      <StyledAppShell>
        <Suspense fallback={<Loader isLoading={true} />}>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/read/:id" element={<ReadPage />} />
            <Route path="/feed/:kind" element={<FeedPage />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </Suspense>
      </StyledAppShell>
    </BaseProvider>
  )
}

export default App
