import { HeadingLarge, Label4 } from 'baseui/typography'
import { Block } from 'baseui/block'
import { FeedItemT } from '../../types'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import { LeftRightButtons } from '../ButtonGroup'
import { createComponent } from '../component'
import { Fragment } from 'react'
import { Skeleton } from 'baseui/skeleton'
import './style.postcss'
import { useStyletron } from 'baseui'

export type ReaderViewProps = {
  feedItem: FeedItemT | null
  font: string
}

type ReaderUIPropsT = {
  id: number
  title: string
  domain: string
  url: string
  hackerNewsUrl: string
  content: string
  onBack: () => void
  onForward: () => void
}

type ReaderPropsT = {
  item: FeedItemT
}

function ui(props: ReaderUIPropsT) {
  const [, theme] = useStyletron()
  const backBtnContent = (
    <Fragment>
      <ChevronLeft />
      Back
    </Fragment>
  )
  const forwardBtnContent = (
    <Fragment>
      Open
      <ChevronRight />
    </Fragment>
  )

  return (
    <Block>
      <Label4 color={theme.colors.contentTertiary}>
        <a href={props.url} target="_blank" rel="noreferrer">
          {props.domain}
        </a>
      </Label4>
      <a href={props.hackerNewsUrl} target="_blank" rel="noreferrer">
        <HeadingLarge>{props.title}</HeadingLarge>
        <Label4
          $style={{
            'text-decoration-line': 'underline',
            color: theme.colors.contentTertiary,
          }}
        >
          Open #{props.id} in HackerNews.
        </Label4>
      </a>
      <Block>
        <div
          id="reader-view-content"
          dangerouslySetInnerHTML={{ __html: props.content }}
        />
      </Block>
      <LeftRightButtons
        onLeft={props.onBack}
        onRight={props.onForward}
        leftContent={backBtnContent}
        rightContent={forwardBtnContent}
      />
    </Block>
  )
}

function prelude(props: ReaderPropsT): ReaderUIPropsT | undefined {
  if (!props || !props.item) {
    return undefined
  }
  const item = props.item
  return {
    id: item.id,
    title: item.title,
    content: item.content,
    domain: item.domain,
    url: item.url,
    hackerNewsUrl: item.hackerNewsUrl,
    onBack: () => {
      window.history.back()
    },
    onForward: () => {
      window.open(item.url, '_blank')
    },
  }
}

const ReaderSkeleton = () => (
  <Fragment>
    <Skeleton
      width="100%"
      height="100px"
      animation
      overrides={{
        Root: {
          style: ({ $theme }) => ({
            marginBottom: $theme.sizing.scale1000,
          }),
        },
      }}
    />
    <Skeleton
      animation={true}
      rows={3}
      width="100%"
      overrides={{
        Root: {
          style: ({ $theme }) => ({
            marginBottom: $theme.sizing.scale1200,
          }),
        },
        Row: {
          style: {
            height: '20px',
            marginBottom: '15px',
          },
        },
      }}
    />
    <Skeleton
      animation={true}
      rows={3}
      width="100%"
      overrides={{
        Root: {
          style: ({ $theme }) => ({
            marginBottom: $theme.sizing.scale600,
          }),
        },
        Row: {
          style: {
            height: '20px',
            marginBottom: '15px',
          },
        },
      }}
    />
  </Fragment>
)

const ReaderContainer = createComponent(ui, prelude)

export { ReaderContainer, ReaderSkeleton }
