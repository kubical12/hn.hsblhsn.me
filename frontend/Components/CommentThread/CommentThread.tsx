import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { CommentConnection } from '../../Types'
import { Comment } from '../Comment'

interface CommentThreadProps {
  comments: CommentConnection | undefined
  isChild?: boolean
}

const CommentThread: React.FC<CommentThreadProps> = ({
  comments,
  isChild = false,
}: CommentThreadProps) => {
  const [css, theme] = useStyletron()
  if (!comments) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }
  const threadStyle = (child: boolean) => {
    return css({
      paddingTop: child ? theme.sizing.scale600 : theme.sizing.scale300,
      borderLeft: child ? `3px solid ${theme.colors.border}` : 'none',
      paddingLeft: child ? theme.sizing.scale300 : theme.sizing.scale0,
    })
  }
  return (
    <div>
      {comments.edges.map(({ node }, index) => (
        <Block className={threadStyle(isChild)} key={index}>
          <Comment comment={node} />
        </Block>
      ))}
    </div>
  )
}

export { CommentThread }
export type { CommentThreadProps }
