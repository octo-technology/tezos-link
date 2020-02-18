import React, { lazy, Suspense } from 'react'
import { Route, Switch } from 'react-router-dom'

import { AppStyled } from './App.style'

const Dashboard = lazy(() =>
  import(
    /*
  webpackChunkName: "dashboard",
  webpackPrefetch: true
  */ '../pages/Dashboard/Dashboard.controller'
  ).then(({ Dashboard }) => ({ default: Dashboard }))
)

const NotFound = lazy(() =>
  import(
    /*
  webpackChunkName: "not-found",
  webpackPrefetch: true
  */ '../pages/NotFound/NotFound.view'
  ).then(({ NotFound }) => ({ default: NotFound }))
)

const Home = lazy(() =>
    import(
        /*
      webpackChunkName: "home",
      webpackPrefetch: true
      */ '../pages/Home/Home.view'
        ).then(({ Home }) => ({ default: Home }))
)

export const AppView = () => (
  <AppStyled>
    <Suspense fallback={<div>Loading...</div>}>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/project/:projectId" component={Dashboard}/>
        <Route component={NotFound} />
      </Switch>
    </Suspense>
  </AppStyled>
)

//<LoggedOutRoute exact path="/login" component={Login} />
