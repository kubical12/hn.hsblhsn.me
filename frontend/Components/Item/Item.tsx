import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { Button, KIND, SIZE } from 'baseui/button'
import { Cell, Grid } from 'baseui/layout-grid'
import { StyledLink } from 'baseui/link'
import {
  HeadingLarge,
  HeadingSmall,
  LabelXSmall,
  ParagraphXSmall,
} from 'baseui/typography'
import { useLocation } from 'react-router-dom'
import { Job, NodeT, Story } from '../../Types'
import { CommentThread } from '../CommentThread'
import { fromNow, getHost, getLink, getTitle } from '../commonutils'
import { TriangleDown, TriangleLeft, TriangleUp } from 'baseui/icon'
import './Item.css'
import { useCallback, useEffect, useMemo, useState } from 'react'
import { Popover } from 'baseui/popover'
import config from '../../app.config'
import { SnackbarProvider, DURATION, useSnackbar } from 'baseui/snackbar'

interface ItemProps {
  item: NodeT<Story | Job>
}

const Item: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  const location = useLocation()
  useEffect(() => {
    if (location.hash === '#comments') {
      document.getElementById('comments')?.scrollIntoView()
    }
  }, [location.key])

  if (!item || (item.type !== 'story' && item.type !== 'job')) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  const shouldShowJumpBtn = useMemo(() => {
    const content = item.text || item.html || item.openGraph?.description || ''
    return content && content.length > 2048
  }, [item])

  return (
    <Block paddingTop={theme.sizing.scale600}>
      <SnackbarProvider>
        <Header item={item} />
        {shouldShowJumpBtn && (
          <Button
            kind={KIND.tertiary}
            size={SIZE.compact}
            startEnhancer={<TriangleDown />}
            onClick={() => {
              window.location.hash = '#comments'
            }}
          >
            Jump to comments
          </Button>
        )}
        <Content item={item} />
        <ActionButtons item={item} />
        <Comments item={item} />
        <ContentLinks item={item} />
      </SnackbarProvider>
    </Block>
  )
}

const Header: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [css, theme] = useStyletron()
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })
  return (
    <Block>
      <LabelXSmall>
        <span className={color(theme.colors.accent)}>@{item.by.id}</span>&nbsp;
        <span className={color(theme.colors.contentSecondary)}>
          {item.time ? fromNow(item.time * 1000) : 'unknown'}
        </span>
      </LabelXSmall>
      <HeadingLarge as="h1">
        {getTitle(item.title, item.openGraph?.title)}
      </HeadingLarge>
      <LabelXSmall
        color={theme.colors.contentTertiary}
        paddingBottom={theme.sizing.scale800}
      >
        Read on&nbsp;
        <StyledLink
          href={getLink(item.id, item.url)}
          target="_blank"
          rel="noreferrer"
        >
          {getHost(item.id, item.url)}
        </StyledLink>
      </LabelXSmall>
    </Block>
  )
}

const Content: React.FC<ItemProps> = ({ item }: ItemProps) => {
  let val = ''
  if ('text' in item && item.text !== '') {
    val = item.text
  } else if ('html' in item && item.html !== '') {
    val = item.html || ''
  } else if (item.openGraph?.description) {
    val = item.openGraph.description
  }
  return (
    <section
      id="reader-view-content"
      dangerouslySetInnerHTML={{ __html: val }}
    />
  )
}

const ActionButtons: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  const [isPopoverOpen, setIsPopoverOpen] = useState(false)
  const btnOverrides = {
    BaseButton: {
      style: {
        width: '100%',
      },
    },
  }
  const back = () => {
    window.history.back()
  }
  const togglePopover = useCallback(() => {
    setIsPopoverOpen(!isPopoverOpen)
  }, [isPopoverOpen])
  return (
    <Block
      paddingTop={theme.sizing.scale1200}
      paddingBottom={theme.sizing.scale600}
    >
      <Grid gridColumns={12} gridGaps={0} gridMargins={0}>
        <Cell span={4}>
          <Button
            onClick={back}
            kind="secondary"
            overrides={btnOverrides}
            startEnhancer={<TriangleLeft />}
          >
            Back
          </Button>
        </Cell>
        <Cell span={8}>
          <Popover
            overrides={{
              Body: {
                style: {
                  backgroundColor: 'rgba(0, 0, 0, 0)',
                },
              },
              Inner: {
                style: {
                  backgroundColor: 'rgba(0, 0, 0, 0)',
                },
              },
            }}
            isOpen={isPopoverOpen}
            onClickOutside={togglePopover}
            onClick={togglePopover}
            content={<MoreBtnPopOver item={item} closeFunc={togglePopover} />}
          >
            <Button
              kind="secondary"
              overrides={btnOverrides}
              startEnhancer={isPopoverOpen ? <TriangleUp /> : <TriangleDown />}
            >
              {isPopoverOpen ? 'Close' : 'More'}
            </Button>
          </Popover>
        </Cell>
      </Grid>
    </Block>
  )
}

interface MoreBtnPopOverProps extends ItemProps {
  closeFunc: () => void
}

const MoreBtnPopOver: React.FC<MoreBtnPopOverProps> = ({
  item,
  closeFunc,
}: MoreBtnPopOverProps) => {
  const [css, theme] = useStyletron()
  const { enqueue } = useSnackbar()
  const popoverCss = css({
    padding: theme.sizing.scale600,
    minWidth: '320px',
    backgroundColor: theme.colors.backgroundTertiary,
    border: `2px solid ${theme.colors.borderTransparent}`,
    borderRadius: theme.sizing.scale300,
  })
  const popoverItemCss = css({
    display: 'flex',
    width: '100%',
    marginBottom: theme.sizing.scale300,
    paddingTop: theme.sizing.scale300,
    paddingBottom: theme.sizing.scale300,
    paddingLeft: theme.sizing.scale400,
    paddingRight: theme.sizing.scale400,
    fontWeight: theme.typography.font750.fontWeight,
    cursor: 'pointer',
    borderRadius: theme.sizing.scale300,
    userSelect: 'none',
    ':hover': {
      backgroundColor: theme.colors.backgroundPrimary,
    },
  })

  const canShare = useMemo(() => {
    return navigator.share !== undefined
  }, [])

  const canCopy = useMemo(() => {
    return navigator.clipboard !== undefined
  }, [])

  const currentPageLink = useMemo(() => {
    return `${config.host}/item?id=${item.id}`
  }, [item.id])

  const copyLink = () => {
    navigator.clipboard.writeText(currentPageLink)
    enqueue(
      {
        message: 'Copied link to clipboard',
      },
      DURATION.short
    )
    closeFunc()
  }

  const openNativeShare = () => {
    if (canShare) {
      navigator.share({
        title: item.title,
        url: currentPageLink,
      })
    }
    closeFunc()
  }
  const openLinkInNewTab = () => {
    window.open(getLink(item.id, item.url), '_blank')
    closeFunc()
  }
  const openInHackerNews = () => {
    window.open(getLink(item.id, undefined), '_blank')
    closeFunc()
  }

  return (
    <Block className={popoverCss}>
      {canCopy && (
        <Block className={popoverItemCss} onClick={copyLink}>
          <CopyIcon /> Copy link
        </Block>
      )}
      {canShare && (
        <Block className={popoverItemCss} onClick={openNativeShare}>
          <ShareIcon /> Share
        </Block>
      )}
      <Block className={popoverItemCss} onClick={openInHackerNews}>
        <HackerNewsIcon />
        Read on HackerNews
      </Block>
      {item.url && (
        <Block className={popoverItemCss} onClick={openLinkInNewTab}>
          <ExternalLinkIcon />
          Open link in new tab
        </Block>
      )}
    </Block>
  )
}

const Comments: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block id="comments">
      {'descendants' in item && (
        <HeadingSmall paddingTop={theme.sizing.scale600}>
          {item.descendants} comments
          <LabelXSmall color={theme.colors.contentTertiary}>
            Posted on&nbsp;
            <StyledLink
              href={getLink(item.id, undefined)}
              target="_blank"
              rel="noreferrer"
            >
              {getHost(item.id, undefined)}
            </StyledLink>
          </LabelXSmall>
        </HeadingSmall>
      )}
      {'comments' in item && (
        <Block>
          <CommentThread parentId={item.id} comments={item.comments} />
        </Block>
      )}
    </Block>
  )
}

const ContentLinks: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block>
      <ParagraphXSmall as="div" paddingTop={theme.sizing.scale2400}>
        <LabelXSmall>
          Contents:
          <br />
          <StyledLink href={getLink(item.id, item.url)}>
            {getLink(item.id, item.url)}
          </StyledLink>
        </LabelXSmall>
        <br />
        <LabelXSmall>
          Comments:
          <br />
          <StyledLink href={getLink(item.id, undefined)}>
            {getLink(item.id, undefined)}
          </StyledLink>
        </LabelXSmall>
      </ParagraphXSmall>
    </Block>
  )
}

const CopyIcon = () => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    className="h-5 w-5 mr-2"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
    strokeWidth={2}
  >
    <path
      strokeLinecap="round"
      strokeLinejoin="round"
      d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
    />
  </svg>
)

const ExternalLinkIcon = () => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    className="h-5 w-5 mr-2"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
    strokeWidth={2}
  >
    <path
      strokeLinecap="round"
      strokeLinejoin="round"
      d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
    />
  </svg>
)

const ShareIcon = () => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    className="h-5 w-5 mr-2"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
    strokeWidth={2}
  >
    <path
      strokeLinecap="round"
      strokeLinejoin="round"
      d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z"
    />
  </svg>
)

const HackerNewsIcon = () => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    className="h-5 w-5 mr-2"
    fill="none"
    viewBox="0 0 24 24"
    stroke="currentColor"
    strokeWidth={2}
  >
    <path
      strokeLinecap="round"
      strokeLinejoin="round"
      d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"
    />
  </svg>
)

export { Item }
export type { ItemProps }
