import React from 'react'
import ReactDOM from 'react-dom'
import { Client as Styletron } from 'styletron-engine-atomic'
import { Provider as StyletronProvider } from 'styletron-react'
import './index.css'
import App from './App'
import { BrowserRouter } from 'react-router-dom'

const engine = new Styletron()

ReactDOM.render(
  <React.StrictMode>
    <StyletronProvider value={engine}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </StyletronProvider>
  </React.StrictMode>,
  document.getElementById('root')
)
