import config from '../../app.config'
import { getBestImage } from '../../Components/commonutils'
import { SEO } from '../../Components/SEO'
import { OpenGraph } from '../../types'

interface HeadProps {
  item: {
    id: string
    title: string
    text?: string
    openGraph?: OpenGraph
  }
}

const Head: React.FC<HeadProps> = ({
  item: { id, title, text, openGraph },
}: HeadProps) => {
  const description = openGraph?.description || text || ''
  return (
    <SEO
      title={`${openGraph?.title || title} | Hacker News`}
      description={description.replace(/<\/?[^>]+(>|$)/g, '')}
      imageUrl={
        getBestImage(openGraph?.image)?.url.replace(
          '&size=thumbnail',
          '&size=full'
        ) || ''
      }
      url={`${config.host}/items?id=${id}`}
    />
  )
}

export { Head }
