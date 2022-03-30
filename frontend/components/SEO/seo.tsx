import { Helmet } from 'react-helmet-async'
import { Empty } from '../Empty'

type SEOFriendly = {
  seo:
    | {
        title: string
        description: string
        imageUrl: string
        canonicalUrl: string
      }
    | undefined
    | null
}

export function Meta(props: { data: SEOFriendly }) {
  if (!props || !props.data || !props.data.seo) {
    return <Empty />
  }
  const seo = props.data.seo
  return (
    <Helmet>
      <title>{seo.title}</title>
      <link rel="canonical" href={seo.canonicalUrl} />
      <meta property="og:title" content={seo.title} />
      <meta property="og:type" content="website" />
      <meta property="og:description" content={seo.description} />
      <meta property="og:image" content={seo.imageUrl} />
      <meta property="og:image:alt" content={seo.title} />
      <meta name="twitter:card" content="summary" />
      <meta name="twitter:title" content={seo.title} />
      <meta name="twitter:description" content={seo.description} />
      <meta name="twitter:image" content={seo.imageUrl} />
      <meta name="twitter:image:alt" content={seo.title} />
      <meta name="description" content={seo.description} />
      <meta name="robots" content="index,follow" />
    </Helmet>
  )
}
