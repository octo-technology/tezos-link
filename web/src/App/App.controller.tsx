import { ConnectedRouter } from 'connected-react-router'
import { History } from 'history'
import * as PropTypes from 'prop-types'
import * as React from 'react'

import { Drawer } from './App.components/Drawer/Drawer.controller'
import { Header } from './App.components/Header/Header.controller'
import { ProgressBar } from './App.components/ProgressBar/ProgressBar.controller'
import { Toaster } from './App.components/Toaster/Toaster.controller'
import { AppView } from './App.view'
import { Hamburger } from './App.components/Hamburger/Hamburger.controller'

type AppProps = { history: History }

export const App = ({ history }: AppProps) => {
  return (
    <>
      <ConnectedRouter history={history}>
        <Header />
        <Drawer />
        <Hamburger />
        <AppView />
        <Toaster />
        <ProgressBar />
      </ConnectedRouter>
    </>
  )
}

App.propTypes = {
  history: PropTypes.object
}
