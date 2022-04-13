export interface Image {
  url: string
  alt: string
  width: number
}

export interface OpenGraph {
  title?: string
  description?: string
  image?: Array<Image>
}
