import { useStyletron } from 'baseui'
import { Button, SHAPE, KIND, SIZE } from 'baseui/button'
import { Card, StyledAction, StyledBody, StyledThumbnail } from 'baseui/card'
import { HeadingSmall, Label4 } from 'baseui/typography'
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
            backgroundColor: theme.colors.backgroundSecondary,
          },
        },
      }}
    >
      <Label4 color={theme.colors.contentTertiary}>
        <a href={getLink(item.id, item.url)} target="_blank" rel="noreferrer">
          {getHost(item.id, item.url)}
        </a>
        &nbsp;-&nbsp;&nbsp;{fromNow(item.time * 1000) || 'unknown time'}
      </Label4>

      {thumbnail && <StyledThumbnail src={thumbnail.url} alt={thumbnail.alt} />}

      <HeadingSmall paddingBottom={theme.sizing.scale600}>
        <Link to={`/item?id=${item.id}`}>
          {getTitle(item.title, item.openGraph?.title)}
        </Link>
      </HeadingSmall>

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
              {item.score || 0} points
            </Button>
          </a>
        )}
        {'descendants' in item && (
          <Link to={`/item?id=${item.id}#comments`}>
            <Button shape={SHAPE.pill} kind={KIND.tertiary} size={SIZE.mini}>
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
