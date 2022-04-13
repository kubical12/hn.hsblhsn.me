import { Block } from 'baseui/block'
import { gql, useQuery } from '@apollo/client'
import { Container } from '../../Components/Layout'
import { NavBar } from '../../Components/Navbar'
import { Item, ITEM_FIELDS } from '../../Components/Item'
import { useSearchParams } from 'react-router-dom'
import { PaddedBlock } from '../../Components/Layout'
import { LoadingScreen } from './LoadingScreen'
import { ErrorScreen } from './ErrorScreen'
import { Item as ItemT } from '../../types'

const GET_ITEM_QUERY = gql`
  ${ITEM_FIELDS}
  query GetItem($id: ID!) {
    item: node(id: $id) {
      ...ItemFields
      ...JobFields
    }
  }
`

interface GetItemQueryData {
  item: ItemT
}

interface GetItemQueryVars {
  id: string
}

const ItemPage: React.FC = () => {
  const [searchParams] = useSearchParams()
  const { data, loading, error } = useQuery<GetItemQueryData, GetItemQueryVars>(
    GET_ITEM_QUERY,
    {
      variables: {
        id: searchParams.get('id') || '',
      },
    }
  )
  return (
    <Container
      top={<NavBar />}
      left={<Block />}
      center={
        <PaddedBlock>
          {loading ? (
            <LoadingScreen />
          ) : error ? (
            <ErrorScreen error={error} />
          ) : data ? (
            <Item item={data.item} />
          ) : // eslint-disable-next-line unicorn/no-null
          null}
        </PaddedBlock>
      }
      right={<Block />}
    />
  )
}

export { ItemPage }
