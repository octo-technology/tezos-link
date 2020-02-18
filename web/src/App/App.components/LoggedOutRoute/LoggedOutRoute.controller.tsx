import * as PropTypes from 'prop-types'
import * as React from 'react'
import { useSelector } from 'react-redux'
import { Redirect, Route } from 'react-router-dom'

type LoggedOutRouteProps = {
  component: any
  path: string
  exact: boolean
}

export const LoggedOutRoute = ({ component: Component, path, exact }: LoggedOutRouteProps) => {
  const user = useSelector((state: any) => state?.auth?.user)

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
              <Redirect to={{ pathname: '/', state: { from: props.location } }} />
            )}
          </>
        ) : (
          <Component {...props} />
        )
      }
    />
  )
}

LoggedOutRoute.propTypes = {
  component: PropTypes.func.isRequired,
  path: PropTypes.string.isRequired,
  exact: PropTypes.bool
}

LoggedOutRoute.defaultProps = {
  exact: false
}
