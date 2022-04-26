import { SEO } from '../../Components/SEO'

interface HeadProps {
  query: string
}

const Head: React.FC<HeadProps> = ({ query }: HeadProps) => {
  return (
    <SEO
      title={`Search ${query || ''} on Hacker News`}
      description="Hacker News Search, millions articles and comments at your fingertips."
      imageUrl={undefined}
      url={`/search`}
    />
  )
}

export { Head }
