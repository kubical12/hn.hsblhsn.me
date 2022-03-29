export type FeedT = {
  items: Array<FeedItemT | null>
}

export type FeedItemT = {
  id: number
  title: string
  summary: string
  content: string
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
