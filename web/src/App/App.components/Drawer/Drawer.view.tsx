import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'

import { DrawerItem, DrawerMask, DrawerStyled } from './Drawer.style'

type DrawerViewProps = {
  showing: boolean
  hideCallback: () => void
  pathname: string
  user: any
  removeAuthUserCallback: () => void
}

export const DrawerView = ({ showing, hideCallback, pathname, user, removeAuthUserCallback }: DrawerViewProps) => (
  <>
    <DrawerMask className={`${showing}`} onClick={() => hideCallback()} />
    <DrawerStyled className={`${showing}`}>
      <h1>Menu</h1>

      {user
        ? loggedInDrawer(user, pathname, hideCallback, removeAuthUserCallback)
        : loggedOutDrawer(pathname, hideCallback)}

      <DrawerItem className={pathname === '/status' ? 'current-path' : 'other-path'}>
        <Link to="/status" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#cards" />
          </svg>
          Status
        </Link>
      </DrawerItem>

      <DrawerItem className={pathname === '/documentation' ? 'current-path' : 'other-path'}>
        <Link to="/documentation" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#documentation" />
          </svg>
          Documentation
        </Link>
      </DrawerItem>
    </DrawerStyled>
  </>
)

function loggedOutDrawer(pathname: any, hideCallback: any) {
  return (
    <>
      <DrawerItem className={pathname === '/sign-in-project' ? 'current-path' : 'other-path'}>
        <Link to="/sign-in-project" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#login" />
          </svg>
          Sign In
        </Link>
      </DrawerItem>

      <DrawerItem className={pathname === '/new-project' ? 'current-path' : 'other-path'}>
        <Link to="/new-project" onClick={() => hideCallback()}>
          <svg>
            <use xlinkHref="/icons/sprites.svg#plus-card" />
          </svg>
          Create
        </Link>
      </DrawerItem>
    </>
  )
}

function loggedInDrawer(user: any, pathname: any, hideCallback: any, removeAuthUserCallback: any) {
  return (
    <>
      <DrawerItem className={pathname === '/profile' ? 'current-path' : 'other-path'}>
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
  pathname: PropTypes.string.isRequired,
  user: PropTypes.object,
  removeAuthUserCallback: PropTypes.func.isRequired
}

DrawerView.defaultProps = {
  showing: false,
  user: undefined
}
