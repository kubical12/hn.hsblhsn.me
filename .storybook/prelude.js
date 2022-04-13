import { BaseProvider, LightTheme } from 'baseui'
import { BrowserRouter } from 'react-router-dom'
import { Client as Styletron } from 'styletron-engine-atomic'
import { Provider as StyletronProvider } from 'styletron-react'
import '../frontend/index.css'
const engine = new Styletron()

export default function withProviders(story) {
  return (
    <StyletronProvider value={engine}>
      <BaseProvider theme={LightTheme}>
        <BrowserRouter>
          {typeof story === 'function' ? story() : story.children}
        </BrowserRouter>
      </BaseProvider>
    </StyletronProvider>
  )
}
