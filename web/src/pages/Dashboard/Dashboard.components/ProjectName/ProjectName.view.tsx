import * as React from 'react'
import { ProjectNameCard } from './ProjectName.style'
import * as PropTypes from 'prop-types'

type ProjectNameViewProps = {
  name: string
}

export const ProjectNameView = ({ name }: ProjectNameViewProps) => {
  return <ProjectNameCard>Project {name}</ProjectNameCard>
}

ProjectNameView.propTypes = {
  name: PropTypes.string
}

ProjectNameView.defaultProps = {
  name: undefined
}
