import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { CommentConnection, StyleProps } from '../../Types'
import { Comment } from '../Comment'
import { ApolloError, useApolloClient } from '@apollo/client'
import { LOAD_MORE_COMMENTS_QUERY } from './CommentThread.graphql'
import React, { useCallback, useState } from 'react'
import { Button, KIND, SIZE } from 'baseui/button'
import { ChevronDown } from 'baseui/icon'

interface CommentThreadProps {
  parentId: string
  comments: CommentConnection | undefined
  isChild?: boolean
}

const CommentThread: React.FC<CommentThreadProps> = ({
  parentId,
  comments,
  isChild = false,
}: CommentThreadProps) => {
  const [css, theme] = useStyletron()
  const [isExpanded, setIsExpanded] = useState(true)
  const toggleIsExpanded = useCallback(() => {
    setIsExpanded(!isExpanded)
  }, [isExpanded])
  if (!comments) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  const threadStyle = (child: boolean) => {
    return css({
      paddingLeft: child ? theme.sizing.scale500 : theme.sizing.scale0,
    })
  }
  return (
    <Block
      className={css({
        display: 'flex',
      })}
    >
      {isChild && (
        // show the side border to toggle comment threads.
        <Block
          className={css({
            minWidth: theme.sizing.scale100,
            maxWidth: theme.sizing.scale100,
            width: theme.sizing.scale100,
            marginTop: theme.sizing.scale100,
            backgroundColor: theme.colors.backgroundTertiary,
            ':hover': {
              backgroundColor: theme.colors.backgroundAccent,
            },
            cursor: 'pointer',
          })}
          onClick={toggleIsExpanded}
        ></Block>
      )}
      {!isExpanded && (
        <Button
          kind={KIND.tertiary}
          className={css({
            marginTop: theme.sizing.scale100,
            padding: theme.sizing.scale300,
            borderRadius: theme.borders.radius300,
            paddingLeft: theme.sizing.scale600,
            opacity: '0.5',
            cursor: 'pointer',
            ':hover': {
              backgroundColor: 'transparent',
            },
          })}
          onClick={toggleIsExpanded}
          startEnhancer={<ChevronDown color={theme.colors.accent} />}
        >
          Expand {comments.edges.length} comment
          {comments.edges.length > 1 ? 's' : ''}
        </Button>
      )}
      {isExpanded && (
        <Block className={threadStyle(isChild)}>
          {comments.edges.map(({ node }, index) => {
            if (!node) {
              // eslint-disable-next-line unicorn/no-null
              return null
            }
            return <Comment comment={node} key={index} />
          })}
          {comments.pageInfo.hasNextPage && (
            <LoadMoreComments
              isChild={isChild}
              parentId={parentId}
              after={comments.pageInfo.endCursor}
            />
          )}
        </Block>
      )}
    </Block>
  )
}

interface LoadMoreCommentsProps {
  parentId: string
  after: string
  isChild: boolean
}

const LoadMoreComments: React.FC<LoadMoreCommentsProps> = ({
  isChild,
  parentId,
  after,
}: LoadMoreCommentsProps) => {
  const client = useApolloClient()
  const [loading, setLoading] = useState(false)
  const [comments, setComments] = useState<CommentConnection | undefined>()
  const [error, setError] = useState<ApolloError | undefined>()
  const [, theme] = useStyletron()
  const onLoadAfter = (parentId: string, after: string) => {
    setLoading(true)
    client
      .query({
        query: LOAD_MORE_COMMENTS_QUERY,
        variables: {
          parentId: parentId,
          after: after,
        },
        fetchPolicy: 'cache-first',
        errorPolicy: 'ignore',
      })
      .then(({ data, loading, error }) => {
        setLoading(loading)
        setComments(data.item.comments)
        setError(error)
      })
      .catch((error) => {
        setError(error)
      })
      .finally(() => {
        setLoading(false)
      })
  }
  if (comments) {
    return (
      <Block paddingLeft={isChild ? theme.sizing.scale500 : 0}>
        <CommentThread
          parentId={parentId}
          comments={comments}
          isChild={isChild}
        />
      </Block>
    )
  }
  return (
    <Button
      isLoading={loading}
      disabled={loading || error !== undefined}
      kind={KIND.tertiary}
      size={SIZE.compact}
      startEnhancer={<ChevronDown size={24} />}
      overrides={{
        BaseButton: {
          style: ({ $theme }: StyleProps) => {
            return {
              marginTop: $theme.sizing.scale600,
              color: $theme.colors.accent,
            }
          },
        },
      }}
      onClick={() => {
        onLoadAfter(parentId, after)
      }}
    >
      {error ? 'Could not load' : 'Load more comments'}
    </Button>
  )
}

export { CommentThread }
export type { CommentThreadProps }
