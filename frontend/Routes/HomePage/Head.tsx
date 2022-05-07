import { SEO } from '../../Components/SEO'
import { useContext } from 'react'
import { ConfigContext } from '../../Components/Config'

interface HeadProps {
  path: string
}

const Head: React.FC<HeadProps> = ({ path }: HeadProps) => {
  const config = useContext(ConfigContext)
  const typ = getStoryType(path)
  return (
    <SEO
      title={`${typ} | Hacker News`}
      description="Hacker News is a social news website focusing on computer science and entrepreneurship."
      imageUrl={undefined}
      url={config.host}
    />
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
