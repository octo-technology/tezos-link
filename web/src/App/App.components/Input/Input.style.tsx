import styled, { keyframes } from 'styled-components/macro'

import { backgroundColor, downColor, upColor, borderColor, primaryColor } from '../../../styles'

export const InputStyled = styled.div`
  position: relative;
  margin-bottom: 5px;
`

export const InputComponent = styled.input`
  width: 100%;
  display: block;
  position: relative;
  height: 40px;
  padding: 12px 16px 12px 40px;
  border-width: 1px;
  border-style: solid;
  border-color: ${borderColor};
  border-radius: 4px;
  transition: border-color 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
  will-change: border-color, box-shadow;
  background-color: ${backgroundColor};

  &:hover {
    border-color: ${primaryColor}80;
  }

  &:focus {
    box-shadow: 0 0 0 2px ${primaryColor}20;
    border-color: ${primaryColor}80;
  }

  &.error {
    border-color: ${downColor};

    &:focus {
      box-shadow: 0 0 0 2px rgba(237, 29, 37, 0.1);
    }
  }

  &.success {
    border-color: ${upColor};

    &:focus {
      box-shadow: 0 0 0 2px rgba(0, 201, 167, 0.1);
    }
  }
`
const zoomIn = keyframes`
  from {
    transform:scale(.2);
    opacity:0
  }
  to {
    transform:scale(1);
    opacity:1
  }
`

export const InputSatus = styled.div`
  display: block;
  position: absolute;
  top: 50%;
  right: 10px;
  z-index: 1;
  width: 20px;
  height: 20px;
  margin-top: -10px;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  visibility: visible;
  pointer-events: none;
  will-change: transform, opacity;

  &.error {
    background-image: url('/icons/input-error.svg');
    animation: ${zoomIn} 0.3s cubic-bezier(0.12, 0.4, 0.29, 1.46);
  }

  &.success {
    background-image: url('/icons/input-success.svg');
    animation: ${zoomIn} 0.3s cubic-bezier(0.12, 0.4, 0.29, 1.46);
  }
`

export const InputIcon = styled.img`
  display: block;
  position: absolute;
  top: 50%;
  left: 10px;
  z-index: 1;
  width: 20px;
  height: 20px;
  margin-top: -10px;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  visibility: visible;
  pointer-events: none;
`
