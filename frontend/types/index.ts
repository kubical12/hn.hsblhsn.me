export type FeedT = {
  items: Array<FeedItemT | null>
  seo: SEODataT | null
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
  seo: SEODataT | null
}

export type ErrorT = {
  message: string
  reason?: string
}

export type SEODataT = {
  title: string
  description: string
  imageUrl: string
  canonicalUrl: string
}
