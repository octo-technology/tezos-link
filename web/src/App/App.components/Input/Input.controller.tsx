import * as PropTypes from 'prop-types'
import * as React from 'react'

import { InputView } from './Input.view'

type InputProps = {
  icon?: string
  placeholder: string
  name?: string
  value?: string
  onChange: any
  onBlur: any
  inputStatus?: 'success' | 'error'
  type: string
}

export const Input = ({ icon, placeholder, name, value, onChange, onBlur, inputStatus, type }: InputProps) => {
  return (
    <InputView
      type={type}
      icon={icon}
      name={name}
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
      inputStatus={inputStatus}
    />
  )
}

Input.propTypes = {
  icon: PropTypes.string,
  placeholder: PropTypes.string,
  name: PropTypes.string,
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  onBlur: PropTypes.func.isRequired,
  inputStatus: PropTypes.string,
  type: PropTypes.string
}

Input.defaultProps = {
  icon: undefined,
  placeholder: undefined,
  name: undefined,
  value: undefined,
  inputStatus: undefined,
  type: 'text'
}
