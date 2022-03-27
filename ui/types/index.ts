export type FeedT = {
  feedItems: Array<FeedItemT | null>
  totalPages: number
  currentPage: number
}

export type FeedItemT = {
  id: number
  title: string
  body: string
  __html: string
  domain: string
  url: string
  hackerNewsUrl: string
  thumbnailUrl: string
  totalPoints: number
  totalComments: number
}

export type ErrorT = {
  message: string
  reason?: string
}
