import { Block } from 'baseui/block'
import { Card, StyledBody, StyledThumbnail, StyledTitle } from 'baseui/card'
import { useStyletron } from 'baseui'
import { ImageT, FeedItemT } from '../../types'
import { Label4 } from 'baseui/typography'
import { useCallback } from 'react'
import useAppNavigator from '../../hooks/navigation'
import { Button, SIZE, SHAPE, KIND } from 'baseui/button'

export function FeedItem({ feedItem }: { feedItem: FeedItemT | null }) {
  const [css] = useStyletron()

  // read the feedItem on click.
  const appNav = useAppNavigator()
  const navigateToReaderView = useCallback(() => {
    if (feedItem) {
      appNav.read(feedItem.id)
    }
  }, [feedItem?.id])

  // do not render null feedItems or feedItems without a body or link.
  // these kind of feedItems does not add any value to the feed.
  if (!feedItem) {
    return null
  }
  if (feedItem.body === '' || feedItem.link === '') {
    return null
  }
  // linkToReaderView props.
  // add pointer cursor to the element.
  // go to the reader view on click.
  const linkToReaderView = {
    onClick: navigateToReaderView,
    className: css({
      cursor: 'pointer',
      display: 'flex',
    }),
  }

  // render the feedItem
  return (
    <Card
      overrides={{
        Root: {
          style: {
            marginBottom: '1rem',
          },
        },
      }}
    >
      <FeedItemSource feedItem={feedItem} />
      <FeedItemThumbnail images={feedItem.images} />
      <StyledTitle {...linkToReaderView}>
        <NoOpLink href={`/read/${feedItem.id}`}>{feedItem.title}</NoOpLink>
      </StyledTitle>
      <StyledBody>{feedItem.body.substring(0, 360)}</StyledBody>
      <StyledBody>
        <FeedItemInteractions feedItem={feedItem} />
      </StyledBody>
    </Card>
  )
}

// FeedItemInteractions returns the interactions of the feedItem.
function FeedItemInteractions({ feedItem }: { feedItem: FeedItemT }) {
  return (
    <Block
      $style={{
        display: 'flex',
        justifyContent: 'flex-end',
      }}
    >
      <a
        href={`https://news.ycombinator.com/item?id=${feedItem.id}`}
        target="_blank"
        rel="noreferrer"
      >
        <Button size={SIZE.mini} shape={SHAPE.pill} kind={KIND.tertiary}>
          {feedItem.totalComments} comments
        </Button>
      </a>
    </Block>
  )
}

// FeedItemThumbnail renders the thumbnail of the feedItem.
function FeedItemThumbnail({ images }: { images: Array<ImageT> | null }) {
  if (!images || images.length === 0) {
    return null
  }
  const primaryImage = findBestImage(images)
  return (
    <StyledThumbnail
      src={primaryImage.url}
      alt={primaryImage.alt}
      onError={(e: { target: { style: { display: string } } }) => {
        e.target.style.display = 'none'
      }}
    />
  )
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

// FeedItemSource renders the source of the feedItem.
export function FeedItemSource({ feedItem }: { feedItem: FeedItemT }) {
  // if the feedItem link is not parsable, do not render anything.
  let url: URL | undefined = undefined
  try {
    url = new URL(feedItem.link)
  } catch (e) {
    console.error(e)
  }
  if (url === undefined) {
    return null
  }

  // link to the original website.
  const linkToWebsite = (
    <a
      href={feedItem.link}
      title={feedItem.title}
      target="_blank"
      rel="noreferrer"
    >
      {url.hostname}
    </a>
  )

  return (
    <Label4
      overrides={{
        Block: {
          style: ({ $theme }) => ({
            marginBottom: $theme.sizing.scale300,
            color: $theme.colors.contentTertiary,
            textOverflow: 'ellipsis',
            whiteSpace: 'nowrap',
          }),
        },
      }}
    >
      {linkToWebsite}
    </Label4>
  )
}

// NoOpLink returns a link element that does nothing.
// It's just a way to help crawlers to get to the linked page.
function NoOpLink({
  href,
  children,
}: {
  href: string
  children: string | JSX.Element
}) {
  const noop = useCallback((e) => {
    e.preventDefault()
    return false
  }, [])
  return (
    <a href={href} onClick={noop}>
      {children}
    </a>
  )
}
