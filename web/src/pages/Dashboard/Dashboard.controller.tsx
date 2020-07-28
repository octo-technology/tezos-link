import * as React from 'react'
import { useDispatch } from 'react-redux'

import { hideProgressBar, showProgressBar } from '../../App/App.components/ProgressBar/ProgressBar.actions'
import { DashboardView } from './Dashboard.view'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { Redirect } from 'react-router-dom'
import { showConfetis } from './Dashboard.components/Confetis/Confetis.controller'
import { ModalProjectCreatedView } from './Dashboard.components/ModalProjectCreated/ModalProjectCreated.view'

export const Dashboard = () => {
  const dispatch = useDispatch()
  const [loading, setLoading] = useState(true)
  const [redirect, setRedirect] = useState(false)
  const [showModal, setShowModal] = useState(false)
  const [project, setProject] = useState({
    title: '',
    uuid: '',
    network: '',
    metrics: {
      requestsCount: 0,
      requestsByDay: [],
      rpcUsage: [],
      lastRequests: []
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
      setTimeout(() => {
        showConfetis()
      }, 500)
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
        response.data.data.metrics.requestsByDay.sort((a: any, b: any) => Date.parse(b.date) - Date.parse(a.date))
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
      {redirect ? <Redirect to={`/project-not-found`} /> : <DashboardView project={project} loading={loading} />}
      {showModal ? <ModalProjectCreatedView uuid={project.uuid} closeModal={closeModal} /> : <></>}
    </>
  )
}
