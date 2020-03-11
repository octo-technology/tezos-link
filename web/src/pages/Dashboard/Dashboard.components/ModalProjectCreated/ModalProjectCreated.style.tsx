import styled from 'styled-components/macro'
import { Card } from '../../../../styles'

export const ModalTitle = styled.h2`
  text-align: center;
`

export const Modal = styled(Card)`
  background: #fff;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 600px;
  height: 400px;
  box-shadow: 0 5px 10px 2px rgba(195, 192, 192, 0.5);
  padding: 20px;
`

export const DashboardOverlay = styled.div`
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: black;
  opacity: 0.5;
`

export const Code = styled.span`
  background-color: #eee;
  border-radius: 2px;
  color: #002b36;
  -moz-osx-font-smoothing: auto;
  -webkit-font-smoothing: auto;
  font-family: monospace;
  padding: 0.25em 0.5em;
`
