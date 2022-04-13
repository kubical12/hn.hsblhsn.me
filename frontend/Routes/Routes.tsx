import { Route, Routes as ReactRouterRoutes } from 'react-router-dom'
import { HomePage } from './HomePage'
import { ItemPage } from './ItemPage'

const Routes = () => (
  <ReactRouterRoutes>
    <Route path="/" element={<HomePage />} />
    <Route path="/newest" element={<HomePage />} />
    <Route path="/ask" element={<HomePage />} />
    <Route path="/jobs" element={<HomePage />} />
    <Route path="/items" element={<ItemPage />} />
    <Route path="*" element={<HomePage />} />
  </ReactRouterRoutes>
)

export { Routes }
