import { gql } from '@apollo/client'
import { COMMENT_THREAD_FIELDS } from '../CommentThread'
import { USER_POPOVER_FIELDS } from '../UserPopover'

const ITEM_FIELDS = gql`
  ${COMMENT_THREAD_FIELDS}
  ${USER_POPOVER_FIELDS}
  fragment ItemFields on Story {
    id
    by {
      ...UserPopoverFields
    }
    time
    url
    title
    text
    html
    type
    score
    descendants
    comments(first: 5) {
      ...CommentThreadFields
    }
    openGraph {
      id
      title
      description
      url
      image {
        url
      }
    }
  }
  fragment JobFields on Job {
    id
    by {
      ...UserPopoverFields
    }
    time
    url
    title
    text
    html
    type
    openGraph {
      id
      title
      description
      url
      image {
        url
      }
    }
  }
`

export { ITEM_FIELDS }
