import React, { useEffect, useState } from 'react'
import { BaseProvider, createDarkTheme, createLightTheme, styled } from 'baseui'
import { Block } from 'baseui/block'
import { ConfigProvider } from '../Config'

const primitives = {
  accent: '#ff6600',
  accent100: '#ffc8a8',
}
const LIGHT_THEME = createLightTheme(primitives)
const DARK_THEME = createDarkTheme(primitives)
const StyledAppShell = styled(Block, ({ $theme }) => ({
  backgroundColor: $theme.colors.backgroundSecondary,
  color: $theme.colors.contentPrimary,
  overflow: 'hidden',
  minHeight: '100vh',
  minWidth: '100vw',
  maxWidth: '38rem',
}))

function useDarkMode() {
  const [isDark, setIsDark] = useState(
    window.matchMedia &&
      window.matchMedia('(prefers-color-scheme: dark)').matches
  )
  // setup theme change listener.
  useEffect(() => {
    const osTheme = window.matchMedia('(prefers-color-scheme: dark)')
    const changeTheme = (e: MediaQueryListEvent) => {
      setIsDark(e.matches)
    }
    osTheme.addEventListener('change', changeTheme)
    return () => {
      osTheme.removeEventListener('change', changeTheme)
    }
  }, [])
  return isDark
}

interface AppProps {
  children: React.ReactNode
}

const App: React.FC<AppProps> = ({ children }: AppProps) => {
  const isDark = useDarkMode()
  return (
    <BaseProvider theme={isDark ? DARK_THEME : LIGHT_THEME}>
      <ConfigProvider>
        <StyledAppShell>{children}</StyledAppShell>
      </ConfigProvider>
    </BaseProvider>
  )
}

export { App }
