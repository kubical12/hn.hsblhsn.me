import { ReactNode, createContext, useEffect, useState } from 'react'
import { defaultConfig, Config } from '../../app.config'

const ConfigContext = createContext<Config>(defaultConfig)

interface ConfiguredWindow extends Window {
  hnAppConfig?: Config
}

let configLoadTicker = 0
const configLoadInterval = 500
const maxConfigLoadTicks = 15

const ConfigProvider = ({ children }: { children: ReactNode }) => {
  const [config, setConfig] = useState<Config>(defaultConfig)
  useEffect(() => {
    const interval = setInterval(() => {
      if (configLoadTicker > maxConfigLoadTicks) {
        clearInterval(interval)
      }
      configLoadTicker++
      const cfgWindow: ConfiguredWindow = window
      // Check if config is loaded
      if (cfgWindow.hnAppConfig) {
        console.log('remote config loaded')
        setConfig({ ...defaultConfig, ...cfgWindow.hnAppConfig })
        clearInterval(interval)
      }
    }, configLoadInterval)

    return () => clearInterval(interval)
  }, [])
  return (
    <ConfigContext.Provider value={config}>{children}</ConfigContext.Provider>
  )
}

export { ConfigContext, ConfigProvider }
