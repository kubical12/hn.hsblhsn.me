import { gql } from '@apollo/client'
import { COMMENT_THREAD_FIELDS } from '../CommentThread'

const ITEM_FIELDS = gql`
  ${COMMENT_THREAD_FIELDS}
  fragment ItemFields on Story {
    id
    by
    time
    url
    title
    text
    html
    type
    descendants
    comments {
      ...CommentThreadFields
    }
  }
  fragment JobFields on Job {
    id
    by
    time
    url
    title
    text
    html
    type
    openGraph {
      title
      description
    }
  }
`

export { ITEM_FIELDS }
