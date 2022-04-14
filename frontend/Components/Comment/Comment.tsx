import { Block } from 'baseui/block'
import { Label3, Paragraph2 } from 'baseui/typography'
import { useStyletron } from 'baseui'
import { CommentThread } from '../CommentThread'
import { useState } from 'react'
import './Comment.css'
import { Comment as CommentT } from '../../types'
import { StyleProps } from '../../types'
import { fromNow } from '../commonutils'

interface CommentProps {
  comment: CommentT
}

const Comment: React.FC<CommentProps> = ({ comment }: CommentProps) => {
  const [css, theme] = useStyletron()
  const [isExpanded, setIsExpanded] = useState(true)
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })

  if (comment.dead || comment.deleted || comment.type !== 'comment') {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  return (
    <Block>
      <Label3
        onClick={() => {
          setIsExpanded(!isExpanded)
        }}
        overrides={{
          Block: {
            style({ $theme }: StyleProps) {
              return {
                paddingTop: $theme.sizing.scale300,
                paddingBottom: $theme.sizing.scale300,
                paddingLeft: $theme.sizing.scale300,
                borderRadius: $theme.borders.radius300,
                cursor: 'pointer',
                backgroundColor: isExpanded
                  ? $theme.colors.backgroundSecondary
                  : $theme.colors.backgroundTertiary,
              }
            },
          },
        }}
      >
        <span className={color(theme.colors.accent)}>@{comment.by}</span>
        <span className={color(theme.colors.colorSecondary)}>
          &nbsp;commented&nbsp;
          {comment.time ? fromNow(comment.time * 1000) : 'unknown'}
        </span>
        <span className={color(theme.colors.colorSecondary)}>
          &nbsp;&nbsp;[{isExpanded ? 'close' : 'show'}]&nbsp;&nbsp;
        </span>
      </Label3>
      <Block display={isExpanded ? 'block' : 'none'}>
        <Paragraph2 as="div" className="comment-reader-view-content">
          <div dangerouslySetInnerHTML={{ __html: comment.text || '' }} />
        </Paragraph2>
        <Block className="pl-3">
          <CommentThread isChild={true} comments={comment.comments} />
        </Block>
      </Block>
    </Block>
  )
}

export { Comment }
export type { CommentProps }
