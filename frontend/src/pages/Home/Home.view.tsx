import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Link } from 'react-router-dom'
import { Button } from 'src/App/App.components/Button/Button.controller'

import { Error404Card, Error404Styled, Error404Title, Error404Message } from './Error404.style'

type Error404ViewProps = { error: string }

export const Error404View = ({ error }: Error404ViewProps) => (
  <Error404Styled>
    <Error404Title>
      <h1>404</h1>
    </Error404Title>
    <Error404Card>
      <Error404Message>{error}</Error404Message>
      <Link to="/">
        <Button text="Go to Dashboard" icon="cards" />
      </Link>
    </Error404Card>
  </Error404Styled>
)

Error404View.propTypes = {
  error: PropTypes.string
}

Error404View.defaultProps = {
  error: 'Not Found'
}
