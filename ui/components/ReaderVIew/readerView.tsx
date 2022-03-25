import { HeadingLarge } from 'baseui/typography'
import { Block } from 'baseui/block'
import { Fragment, useCallback } from 'react'
import { FeedItemT } from '../../types'
import { ErrorScreen } from '../ErrorScreen'
import { useStyletron } from 'baseui'
import { FeedItemSource } from '../Feed'
import { Button, KIND, SHAPE } from 'baseui/button'
import { FlexGrid, FlexGridItem } from 'baseui/flex-grid'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import './style.css'

export type ReaderViewProps = {
  feedItem: FeedItemT | null
  font: string
}

export function ReaderView({ feedItem, font }: ReaderViewProps) {
  const [css, theme] = useStyletron()
  const goBack = useCallback(() => {
    window.history.back()
  }, [])

  const openInHN = useCallback(
    (e) => {
      e.preventDefault()
      const hnLink = `https://news.ycombinator.com/item?id=${feedItem?.id}`
      window.open(hnLink, '_blank')
      return false
    },
    [feedItem?.id]
  )

  const fullWidthBtn = {
    BaseButton: {
      style: {
        width: '100%',
        overflow: 'hidden',
      },
    },
  }

  if (!feedItem) {
    return (
      <ErrorScreen
        error={{
          message: 'Could not open the feedItem in reader view.',
          reason: 'FeedItem is null.',
        }}
      />
    )
  }
  return (
    <Fragment>
      <Block
        className={css({
          marginBottom: theme.sizing.scale1000,
        })}
      >
        <HeadingLarge>
          <a href={feedItem.link} target="_blank" rel="noreferrer">
            {feedItem.title}
          </a>
        </HeadingLarge>
        <FeedItemSource feedItem={feedItem} />
      </Block>
      <div
        className={`reader-view ${font}`}
        dangerouslySetInnerHTML={{ __html: feedItem.__html }}
      />
      <FlexGrid
        flexGridColumnCount={[2, 2, 2, 2]}
        flexGridColumnGap={theme.sizing.scale1000}
        marginTop={theme.sizing.scale1000}
        marginBottom={theme.sizing.scale1000}
      >
        <FlexGridItem>
          <Button
            kind={KIND.secondary}
            startEnhancer={<ChevronLeft />}
            onClick={goBack}
            shape={SHAPE.pill}
            overrides={fullWidthBtn}
          >
            Back
          </Button>
        </FlexGridItem>
        <FlexGridItem>
          <Button
            $as={'a'}
            href={feedItem.link}
            kind={KIND.secondary}
            shape={SHAPE.pill}
            endEnhancer={<ChevronRight />}
            onClick={openInHN}
            overrides={fullWidthBtn}
          >
            HackerNews
          </Button>
        </FlexGridItem>
      </FlexGrid>
    </Fragment>
  )
}
