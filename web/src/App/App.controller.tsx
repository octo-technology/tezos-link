import { ConnectedRouter } from 'connected-react-router'
import { History } from 'history'
import * as PropTypes from 'prop-types'
import * as React from 'react'
import { useEffect } from 'react'
import { useDispatch } from 'react-redux'

import { Drawer } from './App.components/Drawer/Drawer.controller'
import { Footer } from './App.components/Footer/Footer.controller'
//import { Hamburger } from './App.components/Hamburger/Hamburger.controller'
import { Header } from './App.components/Header/Header.controller'
import { ProgressBar } from './App.components/ProgressBar/ProgressBar.controller'
import { Toaster } from './App.components/Toaster/Toaster.controller'
import { AppView } from './App.view'

type AppProps = { history: History }

export const App = ({ history }: AppProps) => {
  const dispatch = useDispatch()

  useEffect(() => {
    //dispatch(getAuthUser())
  }, [dispatch])

  return (
    <>
      <ConnectedRouter history={history}>
        <Header />
        <Drawer />

        <AppView />
        <Toaster />
        <ProgressBar />
        <Footer />
      </ConnectedRouter>
    </>
  )
}

// <Hamburger />

App.propTypes = {
  history: PropTypes.object
}
