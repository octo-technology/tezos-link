import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'

import { DrawerItem, DrawerMask, DrawerStyled } from './Drawer.style'

type DrawerViewProps = {
  showing: boolean
  hideCallback: () => void
  route: string
  user: any
  removeAuthUserCallback: () => void
}

export const DrawerView = ({ showing, hideCallback, route, user, removeAuthUserCallback }: DrawerViewProps) => (
  <>
    <DrawerMask className={`${showing}`} onClick={() => hideCallback()} />
    <DrawerStyled className={`${showing}`}>
      <h1>Menu</h1>

      {user ? loggedInDrawer(user, route, hideCallback, removeAuthUserCallback) : loggedOutDrawer(route, hideCallback)}

      <DrawerItem className={route === '/' ? 'current-path' : 'other-path'}>
        <Link to="/" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#cards" />
          </svg>
          Dashboard
        </Link>
      </DrawerItem>

      <DrawerItem className={route === '/new-post' ? 'current-path' : 'other-path'}>
        <Link to="/new-post" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#plus-card" />
          </svg>
          New Post
        </Link>
      </DrawerItem>
    </DrawerStyled>
  </>
)

function loggedOutDrawer(route: any, hideCallback: any) {
  return (
    <>
      <DrawerItem className={route === '/login' ? 'current-path' : 'other-path'}>
        <Link to="/login" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#login" />
          </svg>
          Login
        </Link>
      </DrawerItem>

      <DrawerItem className={route === '/sign-up' ? 'current-path' : 'other-path'}>
        <Link to="/sign-up" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#sign-up" />
          </svg>
          Sign Up
        </Link>
      </DrawerItem>
    </>
  )
}

function loggedInDrawer(user: any, route: any, hideCallback: any, removeAuthUserCallback: any) {
  return (
    <>
      <DrawerItem className={route === '/profile' ? 'current-path' : 'other-path'}>
        <Link to="/profile" onClick={() => hideCallback()}>
          <img alt={user.username} src="/images/user1.jpg" />
          {user.username}
        </Link>
      </DrawerItem>

      <DrawerItem className={'other-path'}>
        <Link
          to="/"
          onClick={() => {
            removeAuthUserCallback()
            hideCallback()
          }}
        >
          <svg>
            <use xlinkHref="/icons/sprites.svg#log-out" />
          </svg>
          Log Out
        </Link>
      </DrawerItem>
    </>
  )
}

DrawerView.propTypes = {
  showing: PropTypes.bool,
  hideCallback: PropTypes.func.isRequired,
  route: PropTypes.string.isRequired,
  user: PropTypes.object,
  removeAuthUserCallback: PropTypes.func.isRequired
}

DrawerView.defaultProps = {
  showing: false,
  user: undefined
}
