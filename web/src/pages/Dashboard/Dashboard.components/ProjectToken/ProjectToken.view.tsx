import * as React from 'react'

import { ProjectTokenCard, ProjectTokenStyled, ProjectTokenTitle } from './ProjectToken.style'
import * as PropTypes from 'prop-types'
import { TokenCopy } from '../TokenCopy/TokenCopy.controller'
import { Link } from "react-router-dom";

type ProjectTokenViewProps = {
  token: string
}

export const ProjectTokenView = ({ token }: ProjectTokenViewProps) => {
  return (
    <ProjectTokenStyled>
      <ProjectTokenCard>
        <ProjectTokenTitle>private token</ProjectTokenTitle>
        <TokenCopy token={token} />
        <p>
          Make sure to <b>save this token</b>, it is both your <b>access to this dashboard</b> and the <b>API key</b> to
          interact with the proxy.
        </p>
        <p>You can find information about how to use our gateway here : <Link to={"/documentation"}>documentation</Link>.</p>
      </ProjectTokenCard>
    </ProjectTokenStyled>
  )
}

ProjectTokenView.propTypes = {
  token: PropTypes.string
}

ProjectTokenView.defaultProps = {
  token: undefined
}
