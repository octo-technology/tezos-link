import * as React from 'react'
import * as PropTypes from 'prop-types'

import { InputStyled, InputComponent, InputSatus, InputIcon } from './Input.style'

type InputViewProps = {
  icon?: string
  placeholder: string
  name?: string
  value?: string
  onChange: any
  onBlur: any
  inputStatus?: 'success' | 'error'
  type: string
}

export const InputView = ({ icon, placeholder, name, value, onChange, onBlur, inputStatus, type }: InputViewProps) => (
  <InputStyled>
    {icon && <InputIcon alt={icon} src={`/icons/${icon}.svg`} />}
    <InputComponent
      type={type}
      name={name}
      className={inputStatus}
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
      autoComplete={name}
    />
    <InputSatus className={inputStatus} />
  </InputStyled>
)

InputView.propTypes = {
  icon: PropTypes.string,
  placeholder: PropTypes.string,
  name: PropTypes.string,
  value: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  onBlur: PropTypes.func.isRequired,
  inputStatus: PropTypes.string,
  type: PropTypes.string
}

InputView.defaultProps = {
  icon: undefined,
  placeholder: undefined,
  name: undefined,
  value: undefined,
  inputStatus: undefined,
  type: 'text'
}
