import { gql } from '@apollo/client'
import { COMMENT_FIELDS } from '../Comment'

const COMMENT_THREAD_FIELDS = gql`
  ${COMMENT_FIELDS}
  fragment CommentThreadFields on CommentConnection {
    edges {
      node {
        ...CommentFields
      }
    }
  }
`

export { COMMENT_THREAD_FIELDS }
