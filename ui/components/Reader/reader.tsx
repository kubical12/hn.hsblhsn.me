import { HeadingLarge, Label4 } from 'baseui/typography'
import { Block } from 'baseui/block'
import { FeedItemT } from '../../types'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import { LeftRightButtons } from '../ButtonGroup'
import { createComponent } from '../component'
import { Fragment } from 'react'
import { Skeleton } from 'baseui/skeleton'
import './style.postcss'

export type ReaderViewProps = {
  feedItem: FeedItemT | null
  font: string
}

type ReaderUIPropsT = {
  title: string
  host: string
  link: string
  hnLink: string
  content: string
  onBack: () => void
  onForward: () => void
}

type ReaderPropsT = {
  item: FeedItemT
}

function ui(props: ReaderUIPropsT) {
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
      <Label4>{props.host}</Label4>
      <HeadingLarge>{props.title}</HeadingLarge>
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
  const url = new URL(item.link)
  const hostname = url.hostname
  const HNLink = `https://news.ycombinator.com/item?id=${item.id}`
  return {
    title: item.title,
    content: item.__html,
    host: hostname,
    link: item.link,
    hnLink: HNLink,
    onBack: () => {
      window.history.back()
    },
    onForward: () => {
      window.open(HNLink, '_blank')
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
