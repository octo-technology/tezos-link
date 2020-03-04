import * as PropTypes from 'prop-types'
import * as React from 'react'

import { InputView } from './Input.view'
import { RefObject } from 'react'

type InputProps = {
  inputRef?: RefObject<HTMLInputElement>
  icon?: string
  placeholder: string
  name?: string
  value?: string
  onChange: any
  onBlur: any
  inputStatus?: 'success' | 'error'
  type: string
}

export const Input = ({
  inputRef,
  icon,
  placeholder,
  name,
  value,
  onChange,
  onBlur,
  inputStatus,
  type
}: InputProps) => {
  return (
    <InputView
      inputRef={inputRef}
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
  inputRef: PropTypes.any,
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
  inputRef: undefined,
  icon: undefined,
  placeholder: undefined,
  name: undefined,
  value: undefined,
  inputStatus: undefined,
  type: 'text'
}
