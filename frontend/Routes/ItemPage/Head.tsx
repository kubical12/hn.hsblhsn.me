import { Helmet } from 'react-helmet-async'
import config from '../../app.config'
import { getBestImage } from '../../Components/commonutils'
import { OpenGraph } from '../../types'

interface HeadProps {
  item: {
    id: string
    title: string
    text: string
    openGraph?: OpenGraph
  }
}

const Head: React.FC<HeadProps> = ({
  item: { id, title, text, openGraph },
}: HeadProps) => {
  const og = {
    title: `${openGraph?.title || title} | Hacker News`,
    description: openGraph?.description || text,
    imageUrl: getBestImage(openGraph?.image)?.url || '',
    url: `${config.host}/items?id=${id}`,
  }
  return (
    <Helmet>
      <title>{og.title}</title>
      <meta property="og:title" content={og.title} />
      <meta name="description" content={og.description} />
      <meta property="og:description" content={og.description} />
      <meta property="og:image" content={og.imageUrl} />
      <meta property="og:url" content={og.url} />
    </Helmet>
  )
}

export { Head }
