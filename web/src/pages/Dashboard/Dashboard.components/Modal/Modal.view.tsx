import * as React from 'react'
import { ProjectNameCard } from './Modal.style'
import * as PropTypes from 'prop-types'

type ModalViewProps = {
  uuid: string
}

export const ModalView = ({ uuid }: ModalViewProps) => {
  return <ProjectNameCard>{uuid}</ProjectNameCard>
}

ModalView.propTypes = {
  uuid: PropTypes.string
}

ModalView.defaultProps = {
  uuid: undefined
}
