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

const NewProject = lazy(() =>
  import(
    /*
      webpackChunkName: "new-project",
      webpackPrefetch: true
      */ '../pages/NewProject/NewProject.controller'
  ).then(({ NewProject }) => ({ default: NewProject }))
)

const SignInProject = lazy(() =>
  import(
    /*
      webpackChunkName: "sign-in-project",
      webpackPrefetch: true
      */ '../pages/SignInProject/SignInProject.controller'
  ).then(({ SignInProject }) => ({ default: SignInProject }))
)

const NotFound = lazy(() =>
  import(
    /*
  webpackChunkName: "not-found",
  webpackPrefetch: true
  */ '../pages/NotFound/NotFound.view'
  ).then(({ NotFound }) => ({ default: NotFound }))
)

const ProjectNotFound = lazy(() =>
  import(
    /*
  webpackChunkName: "project-not-found",
  webpackPrefetch: true
  */ '../pages/ProjectNotFound/ProjectNotFound.view'
  ).then(({ ProjectNotFound }) => ({ default: ProjectNotFound }))
)

const Home = lazy(() =>
  import(
    /*
      webpackChunkName: "home",
      webpackPrefetch: true
      */ '../pages/Home/Home.view'
  ).then(({ Home }) => ({ default: Home }))
)

const Documentation = lazy(() =>
  import(
    /*
      webpackChunkName: "documentation",
      webpackPrefetch: true
      */ '../pages/Documentation/Documentation.controller'
  ).then(({ Documentation }) => ({ default: Documentation }))
)

const Status = lazy(() =>
  import(
    /*
      webpackChunkName: "status",
      webpackPrefetch: true
      */ '../pages/Status/Status.controller'
  ).then(({ Status }) => ({ default: Status }))
)

export const AppView = () => (
  <AppStyled>
    <Suspense fallback={<div>Loading...</div>}>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/project/:projectId" component={Dashboard} />
        <Route exact path="/new-project" component={NewProject} />
        <Route exact path="/sign-in-project" component={SignInProject} />
        <Route exact path="/documentation" component={Documentation} />
        <Route exact path="/status" component={Status} />
        <Route exact path="/project-not-found" component={ProjectNotFound} />
        <Route component={NotFound} />
      </Switch>
    </Suspense>
  </AppStyled>
)

// <LoggedOutRoute exact path="/login" component={Login} />
