import { useStyletron } from 'baseui'
import { Button, KIND, SHAPE, SIZE } from 'baseui/button'
import { Card, StyledAction, StyledBody, StyledThumbnail } from 'baseui/card'
import { HeadingXSmall, LabelXSmall } from 'baseui/typography'
import { Link } from 'react-router-dom'
import { Job, Story } from '../../Types'
import {
  fromNow,
  getBestImage,
  getHost,
  getLink,
  getTitle,
} from '../commonutils'

interface ItemCardProps {
  item: Story | Job
}

const ItemCard: React.FC<ItemCardProps> = ({ item }: ItemCardProps) => {
  const [, theme] = useStyletron()
  const thumbnail = getBestImage(item.openGraph?.image)
  return (
    <Card
      overrides={{
        Root: {
          style: {
            backgroundColor: 'inherit',
          },
        },
      }}
    >
      <LabelXSmall color={theme.colors.contentTertiary}>
        <a href={getLink(item.id, item.url)} target="_blank" rel="noreferrer">
          {getHost(item.id, item.url).toUpperCase()}
        </a>
        &nbsp;&middot;&nbsp;&nbsp;{fromNow(item.time * 1000) || 'unknown time'}
      </LabelXSmall>

      {thumbnail && <StyledThumbnail src={thumbnail.url} alt={thumbnail.alt} />}

      <Link to={`/item?id=${item.id}`}>
        <HeadingXSmall
          paddingBottom={theme.sizing.scale600}
          paddingTop={theme.sizing.scale200}
        >
          {getTitle(item.title, item.openGraph?.title)}
        </HeadingXSmall>
      </Link>

      <StyledBody>
        {'text' in item && item.text !== '' ? (
          <section
            dangerouslySetInnerHTML={{ __html: ellipsis(item.text, 360) }}
          />
        ) : (
          <section>{ellipsis(item.openGraph?.description, 360)}</section>
        )}
      </StyledBody>

      <StyledAction>
        {'score' in item && (
          <a href={getLink(item.id)} target="_blank" rel="noreferrer">
            <Button shape={SHAPE.pill} kind={KIND.tertiary} size={SIZE.mini}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="w-4 h-4 mr-1"
                viewBox="0 0 24 24"
                fill="currentColor"
              >
                <path d="M12.781 2.375c-.381-.475-1.181-.475-1.562 0l-8 10A1.001 1.001 0 0 0 4 14h4v7a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-7h4a1.001 1.001 0 0 0 .781-1.625l-8-10zM15 12h-1v8h-4v-8H6.081L12 4.601 17.919 12H15z" />
              </svg>
              {item.score || 0} points
            </Button>
          </a>
        )}
        {'descendants' in item && (
          <Link to={`/item?id=${item.id}#comments`}>
            <Button shape={SHAPE.pill} kind={KIND.tertiary} size={SIZE.mini}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="w-4 h-4 mr-1"
                viewBox="0 0 24 24"
                fill="currentColor"
              >
                <path d="M8 21a1 1 0 0 1-1-1v-3H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6.586l-3.707 3.707A1 1 0 0 1 8 21zM19 5H5v10h3a1 1 0 0 1 1 1v1.586l2.293-2.293A1 1 0 0 1 12 15h7V5z" />
                <circle cx="16" cy="10" r="1" />
                <circle cx="12" cy="10" r="1" />
                <circle cx="8" cy="10" r="1" />
              </svg>
              {item.descendants || 0} comments
            </Button>
          </Link>
        )}
      </StyledAction>
    </Card>
  )
}

const ellipsis = (str: string | undefined, max: number) => {
  if (!str) return ''
  if (str.length <= max) {
    return str
  }
  return str.substring(0, max) + '...'
}

export { ItemCard }
export type { ItemCardProps }
