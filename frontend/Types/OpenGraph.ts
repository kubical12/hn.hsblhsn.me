export interface Image {
  url: string
  alt: string
  width: number
}

export interface OpenGraph {
  title?: string
  description?: string
  url?: string
  image?: Array<Image>
}
