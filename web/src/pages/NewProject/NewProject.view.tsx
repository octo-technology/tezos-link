import { Field, Formik } from 'formik'
import * as PropTypes from 'prop-types'
import * as React from 'react'

import { Button } from '../../App/App.components/Button/Button.controller'
import { Input } from '../../App/App.components/Input/Input.controller'
import { FormInputField } from './NewProject.components/FormInputField/FormInputField.controller'
import { FormInputNetworkSelector } from './NewProject.components/FormInputNetworkSelector/FormInputNetworkSelector.controller'
import { NewProjectCard } from './NewProject.style'
import { newProjectValidator } from './NewProject.validator'

type NewPostViewProps = {
  handleSubmitForm: (values: any, actions: any) => void
  loading: boolean
}

export const NewProjectView = ({ handleSubmitForm, loading }: NewPostViewProps) => (
  <NewProjectCard>
    <h1>New Project</h1>
    <Formik
      initialValues={{
        title: ''
      }}
      validationSchema={newProjectValidator}
      validateOnBlur={false}
      onSubmit={(values, actions) => handleSubmitForm(values, actions)}
    >
      {formikProps => {
        const { handleSubmit } = formikProps
        return (
          <form onSubmit={handleSubmit}>
              <Field InputType={Input} component={FormInputNetworkSelector} icon="network" name="network" />

              <Field InputType={Input} component={FormInputField} icon="user" name="title" placeholder="Project title" />

            <Button text="Create project" icon="sign-up" type="submit" loading={loading} />
          </form>
        )
      }}
    </Formik>
  </NewProjectCard>
)

NewProjectView.propTypes = {
  handleSubmitForm: PropTypes.func.isRequired,
  loading: PropTypes.bool
}

NewProjectView.defaultProps = {
  loading: false
}
