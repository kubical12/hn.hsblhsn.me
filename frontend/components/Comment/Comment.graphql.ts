import { gql } from '@apollo/client'

const COMMENT_FIELDS = gql`
  fragment CommentContentFields on Comment {
    id
    type
    by
    time
    text
    dead
    deleted
  }
  fragment CommentFields on Comment {
    ...CommentContentFields
    comments {
      edges {
        node {
          ...CommentContentFields
          comments {
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
