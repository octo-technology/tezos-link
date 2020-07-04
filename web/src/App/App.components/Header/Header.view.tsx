import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'

import { Button } from '../Button/Button.controller'
// prettier-ignore
import {
  HeaderLoggedOut,
  HeaderMenu,
  HeaderStyled,
  HeaderLogo
} from "./Header.style"

export const HeaderView: any = () => {
  return (
    <HeaderStyled>
      <HeaderMenu>
        <Link to="/">
          <img alt="Tezos Link" src="/link_logo.svg" />
        </Link>
      </HeaderMenu>

      <div className="header-triangle"></div>
      <HeaderLogo>
        <Link to="/">
          <img alt="entire stack" src="/logo.svg" />
        </Link>
      </HeaderLogo>

      {loggedOutHeader()}
    </HeaderStyled>
  )
}

function loggedOutHeader() {
  return (
    <HeaderLoggedOut>
      <Link to="/sign-in-project">
        <Button color="transparent" text="SIGN IN" icon="login" />
      </Link>
      <Link to="/new-project">
        <Button text="CREATE" icon="plus-card" />
      </Link>
    </HeaderLoggedOut>
  )
}

HeaderView.propTypes = {
  router: PropTypes.object.isRequired
}
