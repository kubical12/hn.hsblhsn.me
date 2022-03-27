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
    <>
      <ChevronLeft />
      Back
    </>
  )
  const forwardBtnContent = (
    <>
      HackerNews
      <ChevronRight />
    </>
  )

  return (
    <Block>
      <Label4 color={theme.colors.contentTertiary}>
        <a href={props.url} target="_blank" rel="noreferrer">
          {props.domain}
        </a>
      </Label4>
      <HeadingLarge>
        <a href={props.url} target="_blank" rel="noreferrer">
          {props.title}
        </a>
      </HeadingLarge>
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
    title: item.title,
    content: item.__html,
    domain: item.domain,
    url: item.url,
    hackerNewsUrl: item.hackerNewsUrl,
    onBack: () => {
      window.history.back()
    },
    onForward: () => {
      window.open(item.hackerNewsUrl, '_blank')
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
            paddingBottom: $theme.sizing.scale1000,
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
            paddingBottom: $theme.sizing.scale1200,
          }),
        },
        Row: {
          style: {
            height: '20px',
            paddingBottom: '15px',
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
            paddingBottom: $theme.sizing.scale600,
          }),
        },
        Row: {
          style: {
            height: '20px',
            paddingBottom: '15px',
          },
        },
      }}
    />
  </Fragment>
)

const ReaderContainer = createComponent(ui, prelude)

export { ReaderContainer, ReaderSkeleton }
