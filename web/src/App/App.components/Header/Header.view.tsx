import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'

import { Button } from '../Button/Button.controller'
// prettier-ignore
import {
  HeaderLoggedOut,
  HeaderMenu,
  HeaderStyled
} from "./Header.style"

export const HeaderView: any = ({ router }: any) => {
  const path = router && router.location ? [router.location.pathname] : undefined
  console.log(path)
  return (
    <HeaderStyled>
      <HeaderMenu>
        <Link to="/">
          <img alt="Tezos Link" src="/link_logo.svg" />
        </Link>
        <Link to="/status">
          <Button color="transparent" text="STATUS" />
        </Link>
        <Link to="/documentation">
          <Button color="transparent" text="DOCS" />
        </Link>
      </HeaderMenu>

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
