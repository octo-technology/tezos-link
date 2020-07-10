import { Field, Formik } from 'formik'
import * as PropTypes from 'prop-types'
import * as React from 'react'
import { useState } from 'react'

import { Button } from '../../App/App.components/Button/Button.controller'
import { Input } from '../../App/App.components/Input/Input.controller'
import { FormInputField } from './NewProject.components/FormInputField/FormInputField.controller'
import { Select } from './NewProject.components/FormInputNetworkSelector/Select.controller'
import { NewProjectCard } from './NewProject.style'
import { newProjectValidator } from './NewProject.validator'

type NewPostViewProps = {
  handleSubmitForm: (values: any, actions: any) => void
  loading: boolean
}

export const NewProjectView = ({ handleSubmitForm, loading  }: NewPostViewProps) => {
        const [networkValue, setNetwork] = useState('MAINNET')

        return (
        <NewProjectCard>
            <h1>New Project</h1>
            <Formik
                initialValues={{title: '', network: ''}}
                validationSchema={newProjectValidator}
                validateOnBlur={false}
                onSubmit={(values, actions) => handleSubmitForm({ ...values, network: networkValue }, actions)}
            >
                {formikProps => {
                    const {handleSubmit} = formikProps
                    return (
                        <div>
                            <Select
                                options={['MAINNET', 'CARTHAGENET']}
                                defaultOption={networkValue}
                                selectCallback={(e: string) => setNetwork(e)}
                            />
                            <form onSubmit={handleSubmit}>

                                <Field InputType={Input} component={FormInputField} icon="user" name="title"
                                       placeholder="Project title"/>

                                <Button text="Create project" icon="sign-up" type="submit" loading={loading}/>
                            </form>
                        </div>
                    )
                }}
            </Formik>
        </NewProjectCard>
    )
}


NewProjectView.propTypes = {
  handleSubmitForm: PropTypes.func.isRequired,
  loading: PropTypes.bool
}

NewProjectView.defaultProps = {
  loading: false
}
