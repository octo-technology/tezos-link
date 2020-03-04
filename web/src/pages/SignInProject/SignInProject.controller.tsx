import * as React from 'react'
import { useState } from 'react'
import { Redirect } from 'react-router-dom'
import { SignInProjectView } from './SignInProject.view'

export const SignInProject = () => {
  const [redirect, setRedirect] = useState(false)
  const [redirectURL, setRedirectURL] = useState('')

  const handleSubmitForm = async (values: any) => {
    const { uuid } = values
    setRedirectURL(uuid)
    setRedirect(true)
  }

  return (
    <>
      {redirect ? (
        <Redirect to={`project/${redirectURL}`} />
      ) : (
        <SignInProjectView handleSubmitForm={handleSubmitForm} />
      )}
    </>
  )
}
