import styled from 'styled-components/macro'
import { Card, backgroundColor, primaryColor, backgroundColor2 } from '../../../../styles'

export const ModalTitle = styled.h2`
  text-align: center;
`

export const Modal = styled(Card)`
  background: ${backgroundColor2};
  position: absolute;
  left: 0;
  right: 0;
  top: 150px;
  margin: auto;
  width: 600px;
  max-width: 90vw;
  padding: 20px;
`

export const DashboardOverlay = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: black;
  opacity: 0.5;
`

export const Code = styled.span`
  background-color: ${backgroundColor};
  border-radius: 2px;
  color: ${primaryColor};
  -moz-osx-font-smoothing: auto;
  -webkit-font-smoothing: auto;
  font-family: monospace;
  padding: 0.25em 0.5em;
  word-break: break-all;
`
