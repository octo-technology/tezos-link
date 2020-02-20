import { keyframes } from 'styled-components/macro'

export const fadeIn = keyframes`
from {
  opacity: 0;
}
to {
  opacity: 1;
}
`

export const fadeInFromLeft = keyframes`
from {
  opacity: 0;
  transform: translate3d(-50px, 0, 0);
}
to {
  opacity: 1;
  transform: translate3d(0px, 0, 0);
}
`

export const fadeInFromRight = keyframes`
from {
  opacity: 0;
  transform: translate3d(50px, 0, 0);
}
to {
  opacity: 1;
  transform: translate3d(0px, 0, 0);
}`

export const fadeInFromTop = keyframes`
  from {
    opacity: 0;
    transform: translate3d(0, -50px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0px, 0);
  }
`

export const fadeInFromBottom = keyframes`
  from {
    opacity: 0;
    transform: translate3d(0, 50px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0px, 0);
  }
`
