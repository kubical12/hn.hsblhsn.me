import { Fragment } from 'react'
import { Helmet } from 'react-helmet-async'
import config from '../../app.config'
import { getBestImage } from '../../Components/commonutils'
import { OpenGraph } from '../../types'

interface HeadProps {
  item: {
    id: string
    openGraph?: OpenGraph
  }
}

const Head: React.FC<HeadProps> = ({ item: { id, openGraph } }: HeadProps) => {
  if (!openGraph) {
    return <Fragment />
  }
  const image = getBestImage(openGraph?.image)
  const link = `${config.host}/items?id=${id}`
  return (
    <Helmet>
      <title>{openGraph.title}</title>
      <meta name="description" content={openGraph.description} />
      <meta property="og:title" content={openGraph.title} />
      <meta property="og:description" content={openGraph.description} />
      <meta property="og:image" content={image?.url} />
      <meta property="og:url" content={link} />
    </Helmet>
  )
}

export { Head }
