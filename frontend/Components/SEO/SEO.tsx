import { Helmet } from 'react-helmet-async'
import config from '../../app.config'
import defaultOGImg from '../../og-banner.jpg'

interface SEOProps {
  title: string
  description: string
  url: string
  imageUrl?: string
}

const SEO: React.FC<SEOProps> = (props: SEOProps) => {
  const imageUrl = props.imageUrl || `${config.host}${defaultOGImg}`
  return (
    <Helmet>
      <title>{props.title}</title>
      <meta property="og:title" content={props.title} />
      <meta name="description" content={props.description} />
      <meta property="og:description" content={props.description} />
      <meta property="og:image" content={imageUrl} />
      <meta property="og:url" content={props.url} />
    </Helmet>
  )
}

export { SEO }
