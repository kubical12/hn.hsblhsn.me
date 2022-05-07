export interface Config {
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

const defaultConfig: Config = {
  host: 'https://hn.hsblhsn.me',
  graphqlEndpoint: '/graphql',
  ads: {
    enabled: false,
    frequency: 4,
  },
}

export { defaultConfig }
