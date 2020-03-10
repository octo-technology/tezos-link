import * as React from 'react'
import { useDispatch } from 'react-redux'

import { hideProgressBar, showProgressBar } from '../../App/App.components/ProgressBar/ProgressBar.actions'
import { DashboardView } from './Dashboard.view'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { Redirect } from 'react-router-dom'
import { Modal } from './Dashboard.components/Modal/Modal.controller'
import { Code, DashboardModal, DashboardOverlay, ModalTitle } from "./Dashboard.style";
import { Button } from '../../App/App.components/Button/Button.controller'
import { TokenCopy } from './Dashboard.components/TokenCopy/TokenCopy.controller'

export const Dashboard = () => {
  const dispatch = useDispatch()
  const [loading, setLoading] = useState(true)
  const [redirect, setRedirect] = useState(false)
  const [showModal, setShowModal] = useState(false)
  const [project, setProject] = useState({
    title: '',
    uuid: '',
    metrics: {
      requestsCount: 0,
      requestsByDay: [],
      rpcUsage: []
    }
  })

  const closeModal = () => {
    setShowModal(false)
  }

  const showModalIfFirstTime = () => {
    const search = window.location.search
    const params = new URLSearchParams(search)
    const firstTime = params.get('first-time')
    if (firstTime === '') {
      setShowModal(true)
    }
  }

  useEffect(() => {
    showModalIfFirstTime()
    setLoading(true)
    dispatch(showProgressBar())
    setRedirect(false)

    const pageURL = window.location.href
    const uuid = pageURL.substr(pageURL.lastIndexOf('/') + 1)
    axios(process.env.REACT_APP_BACKEND_URL + '/api/v1/projects/' + uuid)
      .then(function(response: any) {
        // TODO: Mock a long request to have a nice loader
        setProject(response.data.data)
        dispatch(hideProgressBar())
        setLoading(false)
      })
      .catch(function(error: any) {
        console.error(error)
        dispatch(hideProgressBar())
        setRedirect(true)
        setLoading(false)
      })
  }, [])

  return (
    <>
      {redirect ? <Redirect to={`/project-not-found`} /> : <DashboardView project={project} loading={loading} />}}
      {showModal ? (
        <>
          <DashboardOverlay />
          <Modal>
            <DashboardModal>
              <ModalTitle>Well done!</ModalTitle>
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
                <p>
                  <b>Make sure to save the following Project ID:</b>
                  <TokenCopy token={project.uuid} />
                </p>
                <p>
                  <Button color={'secondary'} onClick={closeModal} text="Got it!" />
                </p>
              </div>
            </DashboardModal>
          </Modal>
        </>
      ) : (
        <></>
      )}
      }
    </>
  )
}
