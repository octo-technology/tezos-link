import * as React from 'react'
import * as PropTypes from 'prop-types'

import { InputStyled, InputComponent, InputSatus, InputIcon } from './Input.style'
import { RefObject } from 'react'

type InputViewProps = {
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

export const InputView = ({
  inputRef,
  icon,
  placeholder,
  name,
  value,
  onChange,
  onBlur,
  inputStatus,
  type
}: InputViewProps) => (
  <InputStyled>
    {icon && <InputIcon alt={icon} src={`/icons/${icon}.svg`} />}
    <InputComponent
      ref={inputRef}
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

InputView.defaultProps = {
  inputRef: undefined,
  icon: undefined,
  placeholder: undefined,
  name: undefined,
  value: undefined,
  inputStatus: undefined,
  type: 'text'
}
