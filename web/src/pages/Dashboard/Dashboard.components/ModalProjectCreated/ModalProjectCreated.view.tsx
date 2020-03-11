import * as PropTypes from 'prop-types'
import * as React from 'react'
import { ModalContainer } from '../ModalContainer/ModalContainer.controller'
import { TokenCopy } from '../TokenCopy/TokenCopy.controller'
import { Button } from '../../../../App/App.components/Button/Button.controller'
import { Code, Modal, DashboardOverlay, ModalTitle } from './ModalProjectCreated.style'

type ModalProjectCreatedViewProps = {
  uuid: string
  closeModal: () => void
}

export const ModalProjectCreatedView = ({ uuid, closeModal }: ModalProjectCreatedViewProps) => {
  return (
    <>
      <DashboardOverlay />
      <ModalContainer>
        <Modal>
          <canvas id="confetis" />
          <ModalTitle id={'modal-title'}>Well done!</ModalTitle>
          <div>
            <h3>Usage</h3>
            <p>
              We've generated a Project ID for you, this identifier is both your <b>login to this dashboard</b> and
              <b> your credential to access the request URL.</b>
            </p>
            <p>
              <Code>{'curl https://<network>.tezoslink.io/v1/YOUR-PROJECT-ID'}</Code>
            </p>
            <h3>Your Project ID</h3>
            <div>
              <b>Make sure to save the following Project ID:</b>
              <TokenCopy token={uuid} />
            </div>
            <p>
              <Button color={'secondary'} onClick={closeModal} text="Got it!" />
            </p>
          </div>
        </Modal>
      </ModalContainer>
    </>
  )
}

ModalProjectCreatedView.propTypes = {
  uuid: PropTypes.string
}

ModalProjectCreatedView.defaultProps = {
  uuid: undefined
}
