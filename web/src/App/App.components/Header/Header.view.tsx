import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'

import { Button } from '../Button/Button.controller'
// prettier-ignore
import {
    HeaderLoggedOut,
    HeaderLogo,
    HeaderMenu,
    HeaderStyled,
} from './Header.style'

export const HeaderView: any = ({ router }: any) => {
  const path = router && router.location ? [router.location.pathname] : undefined
  console.log(path)
  return (
    <HeaderStyled>
      <HeaderLogo>
          <Link to="/">
              <img alt="Tezos Link" src="/logo.svg" />
          </Link>
          <HeaderMenu>
              <Link to="/product">
                  <Button color="transparent" text="PRODUCT" />
              </Link>
              <Link to="/product">
                  <Button color="transparent" text="STATUS" />
              </Link>
              <a href="/docs">
                  <Button color="transparent" text="DOCS" />
              </a>
          </HeaderMenu>
      </HeaderLogo>

      {loggedOutHeader()}
    </HeaderStyled>
  )
}

function loggedOutHeader() {
  return (
    <HeaderLoggedOut>
      <Link to="/project/test">
        <Button color="transparent" text="MY PROJECT" icon="sign-up" />
      </Link>
      <Link to="/new-project">
        <Button text="CREATE" icon="login" />
      </Link>
    </HeaderLoggedOut>
  )
}

HeaderView.propTypes = {
  router: PropTypes.object.isRequired
}
