import * as React from 'react'
import { Link } from 'react-router-dom'
import { Button } from '../../App/App.components/Button/Button.controller'

import { HomeButton, ProjectNotFoundStyled, ProjectNotFoundText } from './ProjectNotFound.style'

export const ProjectNotFound = () => (
  <ProjectNotFoundStyled>
    <img alt="Unplugged" height="50" src="/icons/island.svg" /> Oops
    <ProjectNotFoundText>Unknown project...</ProjectNotFoundText>
    <HomeButton>
      <Link to="/">
        <Button text="Go to Home" icon="cards" />
      </Link>
    </HomeButton>
  </ProjectNotFoundStyled>
)
