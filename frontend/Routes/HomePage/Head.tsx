import { Helmet } from 'react-helmet-async'
import config from '../../app.config'

interface HeadProps {
  path: string
}

const Head: React.FC<HeadProps> = ({ path }: HeadProps) => {
  const typ = getStoryType(path)
  const og = {
    title: `${typ} | Hacker News`,
    description:
      'Hacker News is a social news website ' +
      'focusing on computer science and entrepreneurship.',
    url: config.host,
    imageUrl: '',
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
