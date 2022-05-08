import { ReactNode, createContext, useEffect, useState } from 'react'
import { defaultConfig, Config } from '../../app.config'

const ConfigContext = createContext<Config>(defaultConfig)

interface ConfiguredWindow extends Window {
  hnAppConfig?: Config
}

let configLoadTicker = 0
const configLoadInterval = 200
const maxConfigLoadTicks = 30

const ConfigProvider = ({ children }: { children: ReactNode }) => {
  const [config, setConfig] = useState<Config>(defaultConfig)

  useEffect(() => {
    const cfgWindow: ConfiguredWindow = window
    if (cfgWindow.hnAppConfig) {
      console.log('remote config loaded (#1001)')
      setConfig({ ...defaultConfig, ...cfgWindow.hnAppConfig })
      return
    }

    const interval = setInterval(() => {
      if (configLoadTicker > maxConfigLoadTicks) {
        clearInterval(interval)
      }
      configLoadTicker++
      if (cfgWindow.hnAppConfig) {
        console.log('remote config loaded (#1002)')
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
