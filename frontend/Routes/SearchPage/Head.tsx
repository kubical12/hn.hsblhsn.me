import { SEO } from '../../Components/SEO'

const Head: React.FC = () => {
  return (
    <SEO
      title="Search | Hacker News"
      description="Hacker News Search, millions articles and comments at your fingertips."
      imageUrl={undefined}
      url={`/search`}
    />
  )
}

export { Head }
