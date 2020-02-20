import * as PropTypes from 'prop-types'
import * as React from 'react'
import { useSelector } from 'react-redux'
import { Redirect, Route } from 'react-router-dom'

type LoggedInRouteProps = {
  component: any
  path: string
  exact: boolean
}

export const LoggedInRoute = ({ component: Component, path, exact }: LoggedInRouteProps) => {
  const user = useSelector((state: any) => (state && state.auth ? state.auth.user : {}))

  return (
    <Route
      path={path}
      exact={exact}
      render={props =>
        user ? (
          <>
            {!user.emailVerified && props.location.pathname !== '/verify-email' ? (
              <Redirect to={{ pathname: '/verify-email', state: { from: props.location } }} />
            ) : (
              <Component {...props} />
            )}
          </>
        ) : (
          <Redirect to={{ pathname: '/login', state: { from: props.location } }} />
        )
      }
    />
  )
}

LoggedInRoute.propTypes = {
  component: PropTypes.func.isRequired,
  path: PropTypes.string.isRequired,
  exact: PropTypes.bool
}

LoggedInRoute.defaultProps = {
  exact: false
}
