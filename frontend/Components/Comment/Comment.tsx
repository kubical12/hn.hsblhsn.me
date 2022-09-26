import { Block } from 'baseui/block'
import { LabelXSmall } from 'baseui/typography'
import { useStyletron } from 'baseui'
import { CommentThread } from '../CommentThread'
import React, { useState } from 'react'
import './Comment.css'
import { Comment as CommentT } from '../../Types'
import { StyleProps } from '../../Types'
import { fromNow } from '../commonutils'

interface CommentProps {
  comment: CommentT
}

const Comment: React.FC<CommentProps> = ({ comment }: CommentProps) => {
  const [css, theme] = useStyletron()
  const shouldBeClosed = comment.dead || comment.deleted
  const [isExpanded, setIsExpanded] = useState(!shouldBeClosed)
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })

  if (comment.type !== 'comment') {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  return (
    <Block>
      <LabelXSmall
        onClick={() => {
          setIsExpanded(!isExpanded)
        }}
        overrides={{
          Block: {
            style({ $theme }: StyleProps) {
              return {
                marginTop: $theme.sizing.scale600,
                paddingTop: $theme.sizing.scale300,
                paddingBottom: $theme.sizing.scale100,
                paddingLeft: $theme.sizing.scale300,
                borderRadius: $theme.borders.radius300,
                cursor: 'pointer',
                backgroundColor: 'inherit',
                border: isExpanded
                  ? `2px solid transparent`
                  : `2px solid ${$theme.colors.borderOpaque}`,
                opacity: isExpanded ? 1 : 0.5,
                transition: 'border-color 0.5s ease-in-out',
              }
            },
          },
        }}
      >
        <span className={color(theme.colors.accent)}>
          {comment.deleted
            ? '[deleted]'
            : comment.dead
            ? '[dead]'
            : `@${comment.by.id}`}
        </span>
        <span className={color(theme.colors.contentSecondary)}>
          &nbsp;&middot;&nbsp;
          {comment.time ? fromNow(comment.time * 1000) : 'unknown'}
        </span>
        <span className={color(theme.colors.contentSecondary)}>
          &nbsp;&nbsp;[{isExpanded ? 'close' : 'show'}]&nbsp;&nbsp;
        </span>
      </LabelXSmall>
      <Block display={isExpanded ? 'block' : 'none'}>
        <div
          className="comment-reader-view-content"
          dangerouslySetInnerHTML={{ __html: comment.text || '' }}
        />
        <Block className="pl-3">
          <CommentThread
            parentId={comment.id}
            isChild={true}
            comments={comment.comments}
          />
        </Block>
      </Block>
    </Block>
  )
}

export { Comment }
export type { CommentProps }
