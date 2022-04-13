import { Image } from '../../types'

export const getHost = (id: string, url?: string) => {
  if (!url) {
    url = `https://news.ycombinator.com/item?id=${id}`
  }
  const domain = url.split('/')[2]
  return domain ? `${domain.toLowerCase()}` : ''
}

export const getLink = (id: string, url?: string) => {
  if (!url) {
    url = `https://news.ycombinator.com/item?id=${id}`
  }
  return url
}

export const getBestImage = (images?: Image[]) => {
  if (!images || images.length == 0) {
    return undefined
  }
  let bestImage = images[0]
  images.forEach((val) => {
    if (val.width > bestImage.width) {
      bestImage = val
    }
  })
  return bestImage
}
