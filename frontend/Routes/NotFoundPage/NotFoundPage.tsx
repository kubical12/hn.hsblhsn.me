import { ErrorScreen } from '../../Components/ErrorScreen'
import { Container, PaddedBlock } from '../../Components/Layout'
import { Fragment } from 'react'
import { NavBar } from '../../Components/NavBar'

const NotFoundPage: React.FC = () => {
  return (
    <Container
      top={<NavBar />}
      left={<Fragment />}
      center={
        <PaddedBlock>
          <ErrorScreen
            error={{
              title: '404',
              message: 'Page not found',
            }}
          />
        </PaddedBlock>
      }
      right={<Fragment />}
    />
  )
}

export { NotFoundPage }
