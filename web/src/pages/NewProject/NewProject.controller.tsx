// eslint-disable-next-line import/no-unresolved,import/no-duplicates
import * as React from 'react'
import { useState } from 'react'
import { useDispatch } from 'react-redux'
import { Redirect } from 'react-router-dom'
import axios from 'axios'
import { hideProgressBar, showProgressBar } from '../../App/App.components/ProgressBar/ProgressBar.actions'
import { showToaster } from '../../App/App.components/Toaster/Toaster.actions'
import { ERROR, SUCCESS } from '../../App/App.components/Toaster/Toaster.constants'
import { NewProjectView } from './NewProject.view'

export const NewProject = () => {
  const dispatch = useDispatch()
  const [loading, setLoading] = useState(false)
  const [redirect, setRedirect] = useState(false)
  const [redirectURL, setRedirectURL] = useState('')

  const handleSubmitForm = async (values: any, actions: any) => {
    const { title, network } = values
    const { setErrors, setSubmitting } = actions
    dispatch(showProgressBar())
    setLoading(true)

    axios({
      method: 'post',
      url: process.env.REACT_APP_BACKEND_URL + '/api/v1/projects',
      data: {
        title: title,
        network: network
      }
    })
      .then(function(response: any) {
        dispatch(showToaster(SUCCESS, 'Welcome!', 'Project created'))
        setRedirectURL(response.headers.location)
        setRedirect(true)
        dispatch(hideProgressBar())
      })
      .catch(function(error: any) {
        console.error(error)
        dispatch(hideProgressBar())
        setSubmitting(false)
        setLoading(false)

        if (error.response) {
          dispatch(showToaster(ERROR, 'Project creation error', error.response.data.data))
          setErrors(error.response.data)
        } else {
          dispatch(showToaster(ERROR, 'Project creation error', error.message))
          setErrors(error.message)
        }
      })
  }

  return (
    <>
      {redirect ? (
        <Redirect to={`project/${redirectURL}?first-time`} />
      ) : (
        <NewProjectView handleSubmitForm={handleSubmitForm} loading={loading} />
      )}
    </>
  )
}
