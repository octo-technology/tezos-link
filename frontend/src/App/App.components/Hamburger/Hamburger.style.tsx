import { textColor } from 'src/styles'
import styled from 'styled-components/macro'
import { HamburgerTopForward } from './Hamburger.styles/HamburgerTopForward'
import { HamburgerTopBackward } from './Hamburger.styles/HamburgerTopBackward'
import { HamburgerBottomForward } from './Hamburger.styles/HamburgerBottomForward'
import { HamburgerBottomBackward } from './Hamburger.styles/HamburgerBottomBackward'

export const HamburgerStyled = styled.div`
  position: fixed;
  left: 20px;
  top: 20px;
  overflow: visible;
  margin: 0;
  height: 24px;
  box-sizing: content-box;
  cursor: pointer;
  z-index: 11;
`

export const HamburgerBox = styled.div`
  position: relative;
  display: inline-block;
  width: 40px;
  height: 24px;
`

export const HamburgerInner = styled.div`
  position: absolute;
  width: 40px;
  height: 1px;
  border-radius: 1px;
  will-change: transform;
  background-color: ${textColor};
`

export const HamburgerInnerTop = styled(HamburgerInner)`
  top: 0;

  &.true {
    animation: ${HamburgerTopForward} 1s linear;
    animation-fill-mode: forwards;
  }

  &.false {
    animation: ${HamburgerTopBackward} 1s linear;
    animation-fill-mode: forwards;
  }
`

export const HamburgerInnerMiddle = styled(HamburgerInner)`
  display: block;
  top: calc(50% - 1px);
`

export const HamburgerInnerBottom = styled(HamburgerInner)`
  bottom: 1px;

  &.true {
    animation: ${HamburgerBottomForward} 1s linear;
    animation-fill-mode: forwards;
  }

  &.false {
    animation: ${HamburgerBottomBackward} 1s linear;
    animation-fill-mode: forwards;
  }
`
