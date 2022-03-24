import { Block } from 'baseui/block'
import { Route, Routes } from 'react-router-dom'
import Navbar from './components/Navbar'
import FeedPage from './pages/Feed'
import HomePage from './pages/Home'
import ReadPage from './pages/Read'

function App() {
  return (
    <Block>
      <Navbar />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/read/:id" element={<ReadPage />} />
        <Route path="/feed/:kind" element={<FeedPage />} />
      </Routes>
    </Block>
  )
}

export default App
