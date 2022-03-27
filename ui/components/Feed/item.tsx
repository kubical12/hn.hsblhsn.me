import { Card, StyledAction, StyledBody, StyledThumbnail } from 'baseui/card'
import { styled } from 'baseui'
import { ImageT, FeedItemT } from '../../types'
import { Button, SIZE, SHAPE, KIND } from 'baseui/button'
import { Skeleton } from 'baseui/skeleton'
import { Link } from 'react-router-dom'
import { HeadingXSmall } from 'baseui/typography'
import { createComponent } from '../component'

// UIProps to pass to the ui component.
type UIProps = {
  title: string
  body: string
  thumbnailUrl?: string
  domain: string
  link: string
  hnLink: string
  readerViewLink: string
  totalComments: number
}

// containerProps is the props that the container component receives.
type ContainerProps = {
  item: FeedItemT
}

// ui of the component.
function ui(props: UIProps) {
  const cardOverrides = {
    Root: {
      style({ $theme }) {
        return {
          marginBottom: $theme.sizing.scale900,
        }
      },
    },
  }

  const showThumbnail = props.thumbnailUrl !== undefined
  const hideImageOnError = (e: { target: { style: { display: string } } }) => {
    e.target.style.display = 'none'
  }

  return (
    <Card overrides={cardOverrides}>
      <StyledItemSource>
        <a
          href={props.link}
          title={props.title}
          target="_blank"
          rel="noreferrer"
        >
          {props.domain}
        </a>
      </StyledItemSource>

      {/* show thumbnail only if there is one! */}
      {showThumbnail && (
        <StyledThumbnail src={props.thumbnailUrl} onError={hideImageOnError} />
      )}

      {/* news title */}
      <HeadingXSmall>
        <Link to={props.readerViewLink}>{props.title}</Link>
      </HeadingXSmall>

      {/* news summary */}
      <StyledBody>{props.body}</StyledBody>

      {/* interactions */}
      <StyledAction>
        <StyledItemInteraction>
          <a href={props.hnLink} target="_blank" rel="noreferrer">
            <Button size={SIZE.mini} shape={SHAPE.pill} kind={KIND.tertiary}>
              {props.totalComments} comments
            </Button>
          </a>
        </StyledItemInteraction>
      </StyledAction>
    </Card>
  )
}

// prelude prepares the container props for the ui component.
function prelude(props: ContainerProps): UIProps | undefined {
  const item = props.item
  if (!item || item.body.trim() === '') {
    return undefined
  }
  const url = new URL(item.link)
  const hostname = url.hostname
  const HNLink = `https://news.ycombinator.com/item?id=${item.id}`
  const linkToReaderView = `/read/${item.id}`
  let thumbnail: string | undefined = undefined
  if (item.images && item.images.length != 0) {
    const primaryImage = findBestImage(item.images)
    thumbnail = primaryImage.url
  }
  return {
    title: item.title,
    body: item.body.substring(0, 360),
    thumbnailUrl: thumbnail,
    domain: hostname,
    link: item.link,
    hnLink: HNLink,
    readerViewLink: linkToReaderView,
    totalComments: item.totalComments,
  }
}

// findBestImage returns the best image from the given array of images.
// it selects the image with the highest width.
function findBestImage(images: Array<ImageT>): ImageT {
  let best: ImageT = images[0]
  for (let i = 1; i < images.length; i++) {
    if (images[i].width > best.width) {
      best = images[i]
    }
  }
  return best
}

// StyledItemInteraction is the container for displaying HackerNews interactions.
const StyledItemInteraction = styled('div', ({ $theme }) => ({
  padding: '0',
  marginTop: $theme.sizing.scale500,
  display: 'flex',
  justifyContent: 'flex-end',
}))

// StyledItemSource is the container for displaying the source (domain) of the item.
const StyledItemSource = styled('div', ({ $theme }) => ({
  display: 'block',
  color: $theme.colors.contentTertiary,
  fontSize: $theme.typography.LabelXSmall.fontSize,
  fontWeight: $theme.typography.LabelXSmall.fontWeight,
  fontFamily: $theme.typography.LabelXSmall.fontFamily,
}))

const ItemSkeleton = () => (
  <Skeleton
    height="15rem"
    width="100%"
    animation
    overrides={{
      Root: {
        style: ({ $theme }) => ({
          borderRadius: $theme.sizing.scale300,
          marginBottom: $theme.sizing.scale900,
        }),
      },
    }}
  />
)
const ItemContainer = createComponent(ui, prelude)

export { ItemContainer, ItemSkeleton }
