import { useEffect, useState } from 'react'

export function useDarkMode() {
  const [isDark, setIsDark] = useState(
    window.matchMedia &&
      window.matchMedia('(prefers-color-scheme: dark)').matches
      ? true
      : false
  )
  // setup theme change listener.
  useEffect(() => {
    const osTheme = window.matchMedia('(prefers-color-scheme: dark)')
    const changeTheme = (e: MediaQueryListEvent) => {
      setIsDark(e.matches ? true : false)
    }
    osTheme.addEventListener('change', changeTheme)
    return () => {
      osTheme.removeEventListener('change', changeTheme)
    }
  }, [])
  return isDark
}
