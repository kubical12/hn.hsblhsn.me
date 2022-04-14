import { useStyletron } from 'baseui'
import { Button, SHAPE, KIND, SIZE } from 'baseui/button'
import { Card, StyledAction, StyledBody, StyledThumbnail } from 'baseui/card'
import { HeadingSmall, Label4 } from 'baseui/typography'
import { Link } from 'react-router-dom'
import { Job, Story } from '../../types'
import { getBestImage, getHost, getLink } from '../commonutils'

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
            backgroundColor: theme.colors.backgroundSecondary,
          },
        },
      }}
    >
      <Label4 color={theme.colors.contentTertiary}>
        <a href={getLink(item.id, item.url)} target="_blank" rel="noreferrer">
          {getHost(item.id, item.url)}
        </a>
      </Label4>

      {thumbnail && <StyledThumbnail src={thumbnail.url} alt={thumbnail.alt} />}

      <HeadingSmall paddingBottom={theme.sizing.scale600}>
        <Link to={`/items?id=${item.id}`}>
          {item.openGraph?.title || item.title}
        </Link>
      </HeadingSmall>

      <StyledBody>
        {'text' in item && item.text !== '' ? (
          <section dangerouslySetInnerHTML={{ __html: item.text }} />
        ) : (
          <section>{item.openGraph?.description}</section>
        )}
      </StyledBody>

      <StyledAction>
        {'score' in item && (
          <a href={getLink(item.id)} target="_blank" rel="noreferrer">
            <Button shape={SHAPE.pill} kind={KIND.tertiary} size={SIZE.mini}>
              {item.score || 0} points
            </Button>
          </a>
        )}
        {'descendants' in item && (
          <Link to={`/items?id=${item.id}#comments`}>
            <Button shape={SHAPE.pill} kind={KIND.tertiary} size={SIZE.mini}>
              {item.descendants || 0} comments
            </Button>
          </Link>
        )}
      </StyledAction>
    </Card>
  )
}

export { ItemCard }
export type { ItemCardProps }
