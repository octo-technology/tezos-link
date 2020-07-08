import * as React from 'react'

import { ProjectNetworkCard } from './ProjectNetwork.style'
import * as PropTypes from 'prop-types'

type ProjectNetworkViewProps = {
  network: string
}

export const ProjectNetworkView = ({ network }: ProjectNetworkViewProps) => {
  return <ProjectNetworkCard>Network {network}</ProjectNetworkCard>
}

ProjectNetworkView.propTypes = {
  network: PropTypes.string
}

ProjectNetworkView.defaultProps = {
  network: undefined
}


