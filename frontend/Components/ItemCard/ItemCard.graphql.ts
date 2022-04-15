import { gql } from '@apollo/client'

const ITEM_CARD_STORY_FIELDS = gql`
  fragment StoryCardFields on Story {
    id
    url
    title
    text
    descendants
    score
    time
    openGraph {
      id
      title
      description
      image {
        url
        alt
        width
      }
    }
  }
`
const ITEM_CARD_JOB_FIELDS = gql`
  fragment JobCardFields on Job {
    id
    url
    title
    time
    openGraph {
      title
      description
      image {
        url
        alt
        width
      }
    }
  }
`

export { ITEM_CARD_STORY_FIELDS, ITEM_CARD_JOB_FIELDS }
