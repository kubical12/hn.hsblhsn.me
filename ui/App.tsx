import { BaseProvider, DarkTheme, LightTheme, styled } from 'baseui'
import { Route, Routes } from 'react-router-dom'
import { useDarkMode } from './hooks/theme'
import FeedPage from './pages/Feed'
import HomePage from './pages/Home'
import ReadPage from './pages/Read'

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
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/read/:id" element={<ReadPage />} />
          <Route path="/feed/:kind" element={<FeedPage />} />
        </Routes>
      </StyledAppShell>
    </BaseProvider>
  )
}

export default App
