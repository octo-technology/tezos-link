import { Field, Formik } from 'formik'
import * as PropTypes from 'prop-types'
import * as React from 'react'

import { Button } from '../../App/App.components/Button/Button.controller'
import { Input } from '../../App/App.components/Input/Input.controller'
import { FormInputField } from './SignInProject.components/FormInputField/FormInputField.controller'
import { SignInProjectCard } from './SignInProject.style'
import { signInProjectValidator } from './SignInProject.validator'

type NewPostViewProps = {
  handleSubmitForm: (values: any) => void
}

export const SignInProjectView = ({ handleSubmitForm }: NewPostViewProps) => (
  <SignInProjectCard>
    <h1>Sign In</h1>
    <Formik
      initialValues={{
        uuid: ''
      }}
      validationSchema={signInProjectValidator}
      validateOnBlur={false}
      onSubmit={values => handleSubmitForm(values)}
    >
      {formikProps => {
        const { handleSubmit } = formikProps
        return (
          <form onSubmit={handleSubmit}>
            <Field
              InputType={Input}
              component={FormInputField}
              icon="user"
              name="uuid"
              placeholder="e4efba4d-47e6-42a5-905e-589d3f673853"
            />

            <Button text="Sign in project" icon="login" type="submit" />
          </form>
        )
      }}
    </Formik>
  </SignInProjectCard>
)

SignInProjectView.propTypes = {
  handleSubmitForm: PropTypes.func.isRequired
}
