import * as PropTypes from 'prop-types'
import * as React from 'react'

import { FormInputFieldViewErrorMessage } from './FormInputField.style'

type FormInputFieldViewProps = {
  errorMessage: any
  inputStatus: any
  props: any
  name: any
  value: any
  onChange: any
  onBlur: any
  InputType: any
}

export const FormInputFieldView = ({
  errorMessage,
  inputStatus,
  name,
  value,
  onChange,
  onBlur,
  InputType,
  props
}: FormInputFieldViewProps) => (
  <>
    <InputType name={name} value={value} onChange={onChange} onBlur={onBlur} inputStatus={inputStatus} {...props} />
    {errorMessage && <FormInputFieldViewErrorMessage>{errorMessage}</FormInputFieldViewErrorMessage>}
  </>
)

FormInputFieldView.propTypes = {
  errorMessage: PropTypes.string,
  inputStatus: PropTypes.string,
  name: PropTypes.string,
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  onChange: PropTypes.func.isRequired,
  onBlur: PropTypes.func.isRequired,
  InputType: PropTypes.func,
  props: PropTypes.object
}
