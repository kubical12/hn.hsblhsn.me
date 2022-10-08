import { gql } from '@apollo/client'

const USER_POPOVER_FIELDS = gql`
  fragment UserPopoverFields on User {
    id
    created
    karma
    about
    submitted {
      totalCount
    }
  }
`

export { USER_POPOVER_FIELDS }
