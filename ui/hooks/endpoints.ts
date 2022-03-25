import config from '../app.config'

export const ENDPOINTS = {
  feedList(kind: string, page: number): string {
    return `${config.apiBasePath}/feeds/${kind}/${page}`
  },
  feedItem(id: number): string {
    return `${config.apiBasePath}/feed_items/${id}`
  },
}
