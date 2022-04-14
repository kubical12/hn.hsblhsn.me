import { Helmet } from 'react-helmet-async'
import config from '../../app.config'
import { getBestImage } from '../../Components/commonutils'
import { OpenGraph } from '../../types'

interface HeadProps {
  path: string
}

const Head: React.FC<HeadProps> = ({ path }: HeadProps) => {
  const typ = getStoryType(path)
  const openGraph: OpenGraph = {
    title: `${typ} | Hacker News`,
    description:
      'Hacker News is a social news website' +
      'focusing on computer science and entrepreneurship.',
    url: config.host,
  }
  const image = getBestImage(openGraph?.image)
  return (
    <Helmet>
      <title>{openGraph.title}</title>
      <meta name="description" content={openGraph.description} />
      <meta property="og:title" content={openGraph.title} />
      <meta property="og:description" content={openGraph.description} />
      <meta property="og:image" content={image?.url} />
      <meta property="og:url" content={openGraph.url} />
    </Helmet>
  )
}

const getStoryType = (path: string): string => {
  if (path.includes('/newest')) {
    return 'New links'
  } else if (path.includes('/ask')) {
    return 'Ask'
  } else if (path.includes('/show')) {
    return 'Show '
  } else if (path.includes('/job')) {
    return 'Job'
  } else {
    return 'Top'
  }
}

export { Head }
