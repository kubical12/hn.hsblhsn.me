import { gql } from '@apollo/client'
import { USER_POPOVER_FIELDS } from '../UserPopover'

const COMMENT_FIELDS = gql`
  ${USER_POPOVER_FIELDS}
  fragment CommentContentFields on Comment {
    id
    type
    by {
      ...UserPopoverFields
    }
    time
    text
    dead
    deleted
  }
  fragment CommentFields on Comment {
    ...CommentContentFields
    comments(first: 5) {
      pageInfo {
        hasNextPage
        endCursor
      }
      edges {
        node {
          ...CommentContentFields
          comments(first: 5) {
            pageInfo {
              hasNextPage
              endCursor
            }
            edges {
              node {
                ...CommentContentFields
              }
            }
          }
        }
      }
    }
  }
`

export { COMMENT_FIELDS }
