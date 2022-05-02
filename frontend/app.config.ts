interface Config {
  host: string
  graphqlEndpoint: string
  ads: {
    enabled: boolean
    frequency: number
    google?: {
      adClient: string
      adSlot: string
      adLayout?: string
    }
  }
}

let config: Config = {
  host: 'https://hn.hsblhsn.me',
  graphqlEndpoint: '/graphql',
  ads: {
    enabled: false,
    frequency: 12,
  },
}

// get the config from local storage
const configStr = window?.localStorage?.getItem('hn_app_config')
try {
  if (configStr) {
    const localConfig = JSON.parse(configStr)
    config = { ...config, ...localConfig }
  }
} catch (e) {
  console.error('Failed to parse config from local storage', e)
}

export default config
