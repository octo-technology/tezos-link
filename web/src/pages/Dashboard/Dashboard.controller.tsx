import * as React from 'react'
import { useDispatch } from 'react-redux'

import { hideProgressBar, showProgressBar } from '../../App/App.components/ProgressBar/ProgressBar.actions'
import { DashboardView } from './Dashboard.view'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { Redirect } from 'react-router-dom'

export const Dashboard = () => {
  const dispatch = useDispatch()
  const [loading, setLoading] = useState(true)
  const [redirect, setRedirect] = useState(false)
  const [project, setProject] = useState({
    title: '',
    uuid: '',
    metrics: {
      requestsCount: 0
    }
  })

  useEffect(() => {
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
    <>{redirect ? <Redirect to={`/project-not-found`} /> : <DashboardView project={project} loading={loading} />}</>
  )
}
