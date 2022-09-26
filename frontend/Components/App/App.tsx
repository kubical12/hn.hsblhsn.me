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

function useTheme() {
  const [theme, setTheme] = useState(
    window.matchMedia &&
      window.matchMedia('(prefers-color-scheme: dark)').matches
      ? DARK_THEME
      : LIGHT_THEME
  )
  // setup theme change listener.
  useEffect(() => {
    const osTheme = window.matchMedia('(prefers-color-scheme: dark)')
    const changeTheme = (e: MediaQueryListEvent) => {
      const shouldUseDarkTheme = e.matches
      setTheme(shouldUseDarkTheme ? DARK_THEME : LIGHT_THEME)
    }
    osTheme.addEventListener('change', changeTheme)
    return () => {
      osTheme.removeEventListener('change', changeTheme)
    }
  }, [])
  return theme
}

interface AppProps {
  children: React.ReactNode
}

const App: React.FC<AppProps> = ({ children }: AppProps) => {
  const theme = useTheme()
  useEffect(() => {
    console.log('setting theme')
    document.body.style.backgroundColor = theme.colors.backgroundSecondary
  }, [theme])
  return (
    <BaseProvider theme={theme}>
      <ConfigProvider>
        <StyledAppShell>{children}</StyledAppShell>
      </ConfigProvider>
    </BaseProvider>
  )
}

export { App }
