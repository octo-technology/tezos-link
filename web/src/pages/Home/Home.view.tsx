import * as React from 'react'
import { Link } from 'react-router-dom'
import { Button } from 'src/App/App.components/Button/Button.controller'

import { CreateButton, HomeText, HomeViewStyled, HomeViewTitle } from './Home.style'

export const Home = () => (
  <HomeViewStyled>
    <HomeViewTitle>
      <h1>YOUR GATEWAY TO THE TEZOS NETWORK</h1>
    </HomeViewTitle>
    <HomeText>
      <h3>We provide scalable API access to the Tezos network and usage analytics for your projects</h3>
    </HomeText>
    <CreateButton>
      <Link to="/new-project">
        <Button text="&nbsp;&nbsp;NEW PROJECT" icon="plus-card" />
      </Link>
    </CreateButton>
  </HomeViewStyled>
)
