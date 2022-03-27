import { Card, StyledAction, StyledBody, StyledThumbnail } from 'baseui/card'
import { styled } from 'baseui'
import { FeedItemT } from '../../types'
import { Button, SIZE, SHAPE, KIND } from 'baseui/button'
import { Skeleton } from 'baseui/skeleton'
import { Link } from 'react-router-dom'
import { HeadingXSmall } from 'baseui/typography'
import { createComponent } from '../component'

// UIProps to pass to the ui component.
type UIProps = FeedItemT

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
          href={props.url}
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
        <Link to={`/read/${props.id}`}>{props.title}</Link>
      </HeadingXSmall>

      {/* news summary */}
      <StyledBody>{props.summary}</StyledBody>

      {/* interactions */}
      <StyledAction>
        <StyledItemInteraction>
          <a href={props.hackerNewsUrl} target="_blank" rel="noreferrer">
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
  if (!item) {
    return undefined
  }
  return item
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
