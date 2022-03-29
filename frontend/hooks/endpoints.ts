import config from '../app.config'

export const ENDPOINTS = {
  feedList(kind: string, page: number): string {
    return `${config.apiBasePath}/lists/${kind}.json?page=${page}`
  },
  feedItem(id: number): string {
    return `${config.apiBasePath}/items/${id}.json`
  },
}
