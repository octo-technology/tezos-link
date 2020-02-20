import * as React from 'react'

import {ProjectTokenCard, ProjectTokenStyled} from './ProjectToken.style'
import {Input} from '../../../../App/App.components/Input/Input.controller'
import * as PropTypes from "prop-types";

type ProjectTokenViewProps = {
    token: string
}

export const ProjectTokenView = ({token}: ProjectTokenViewProps) => {
  return (
    <ProjectTokenStyled>
      <ProjectTokenCard>
          Your private token:
        <Input value={token} name="token" onChange={() => {}} onBlur={() => {}} />
      </ProjectTokenCard>
    </ProjectTokenStyled>
  )
}

ProjectTokenView.propTypes = {
    token: PropTypes.string
}

ProjectTokenView.defaultProps = {
    token: undefined,
}
