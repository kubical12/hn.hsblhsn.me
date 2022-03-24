const HOST = 'http://localhost:8080'
const BASE_PATH = `${HOST}/api/v1`

export const ENDPOINTS = {
  feedList(kind: string, page: number): string {
    return `${BASE_PATH}/feeds/${kind}/${page}`
  },
  feedItem(id: number): string {
    return `${BASE_PATH}/feed_items/${id}`
  },
}
