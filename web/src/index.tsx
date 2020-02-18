import * as Sentry from '@sentry/browser'
import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { BrowserRouter as Router } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'

import { configureStore, history } from './App/App.config/configureStore'
import { App } from './App/App.controller'
import { register } from './serviceWorker'
import { GlobalStyle } from './styles'

import './styles/index.scss'

Sentry.init({ dsn: process.env.REACT_APP_SENTRY_DSN, release: 'sms-client@' + process.env.npm_pacakage_version })

const store = configureStore({})
const Root = () => {
  return (
    <ThemeProvider theme={{ color: 'red' }}>
      <GlobalStyle />
      <Provider store={store}>
        <Router>
            <App history={history} />
        </Router>
      </Provider>
    </ThemeProvider>
  )
}

const rootElement = document.getElementById('root')
ReactDOM.render(<Root />, rootElement)

register()
