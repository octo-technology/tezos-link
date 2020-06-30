import styled from 'styled-components/macro'

import { fadeIn, fadeInFromLeft, fadeInFromTop, fadeInFromRight, fadeInFromBottom } from './animations'
import { backgroundColor2, shadowColor } from '.'

export const Ellipsis = styled.div`
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
`

export const TextLight = styled.div`
  font-weight: 300;
`

export const IconSmall = styled.img`
  width: 16px;
  height: 16px;
`

export const Card = styled.div`
  border-radius: 6px;
  background: ${backgroundColor2};
  box-shadow: 0 1px 10px ${shadowColor}80;
  margin: auto;
`

export const AnimatedCard = styled.div`
  border-radius: 6px;
  background: ${backgroundColor2};
  box-shadow: 0 1px 10px ${shadowColor}80;
  will-change: opacity, transform;
  animation: ${fadeInFromLeft} 500ms;
`

export const FadeIn = styled.div`
  will-change: opacity;
  animation: ${fadeIn} 500ms;
`

export const FadeInFromTop = styled.div`
  will-change: opacity, transform;
  animation: ${fadeInFromTop} 500ms;
`

export const FadeInFromRight = styled.div`
  will-change: opacity, transform;
  animation: ${fadeInFromRight} 500ms;
`

export const FadeInFromBottom = styled.div`
  will-change: opacity, transform;
  animation: ${fadeInFromBottom} 500ms;
`

export const FadeInFromLeft = styled.div`
  will-change: opacity, transform;
  animation: ${fadeInFromLeft} 500ms;
`
