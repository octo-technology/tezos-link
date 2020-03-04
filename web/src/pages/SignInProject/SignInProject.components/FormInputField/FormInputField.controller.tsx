import * as PropTypes from 'prop-types'
import * as React from 'react'

import { FormInputFieldView } from './FormInputField.view'

type FormInputFieldProps = {
  field: any
  form: any
  InputType: any
  hideErrorMessage: boolean
}

export const FormInputField = ({
  field: { name, value, onChange, onBlur },
  form: { touched, errors }, // also values, setXXXX, handleXXXX, dirty, isValid, status, etc.
  InputType,
  hideErrorMessage,
  ...props
}: FormInputFieldProps) => {
  const errorMessage = hideErrorMessage ? undefined : touched[name] && errors[name]

  let inputStatus

  if (errorMessage) {
    inputStatus = 'error'
  } else if (touched[name] && value) {
    inputStatus = 'success'
  }

  return (
    <FormInputFieldView
      errorMessage={errorMessage}
      inputStatus={inputStatus}
      name={name}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
      InputType={InputType}
      props={props}
    />
  )
}

FormInputField.propTypes = {
  field: PropTypes.object,
  form: PropTypes.object,
  inputStatus: PropTypes.string,
  InputType: PropTypes.func,
  props: PropTypes.object
}
