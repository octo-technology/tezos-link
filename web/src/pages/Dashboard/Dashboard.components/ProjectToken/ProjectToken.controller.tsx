import * as React from 'react'

import {ProjectTokenView} from './ProjectToken.view'
import * as PropTypes from "prop-types";

type ProjectTokenProps = {
  token: string
}

export const ProjectToken = ({token}: ProjectTokenProps) => {
  return <ProjectTokenView token={token}/>
}

ProjectToken.propTypes = {
  token: PropTypes.string
}

ProjectToken.defaultProps = {
  token: undefined,
}
