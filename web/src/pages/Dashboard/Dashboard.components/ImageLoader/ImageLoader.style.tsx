import styled, { keyframes } from 'styled-components/macro'

const GraphLoaderAnimation = keyframes`
to {
    background-position: 350% 0, var(--soft-image-position), 0 0;
  }
`

export const GraphLoaderStyled = styled.div`
  width: 100%;
  height: 100%;

  --soft-card-skeleton: linear-gradient(rgba(255, 255, 255, 1) 66px, transparent 0);

  --soft-image-height: 66px;
  --soft-image-width: 117px;
  --soft-image-position: 0 0;
  --soft-image-skeleton: linear-gradient(rgba(245, 249, 252, 1) var(--soft-image-height), transparent 0);

  --soft-blur-size: 200px 66px;

  display: block;

  background-image: linear-gradient(
      90deg,
      rgba(255, 255, 255, 0) 0,
      rgba(255, 255, 255, 0.8) 50%,
      rgba(255, 255, 255, 0) 100%
    ),
    var(--soft-image-skeleton), var(--soft-card-skeleton);
  background-size: var(--soft-blur-size), var(--soft-image-width) var(--soft-image-height), 100% 100%;
  background-position: -150% 0, var(--soft-image-position), 0 0;
  background-repeat: no-repeat;
  animation: ${GraphLoaderAnimation} 1.5s infinite;
`

export const GraphLoaderImg = styled.img`
  width: 100%;
  height: 100%;
  opacity: 0;
  transition: opacity 1s;
  will-change: opacity;

  &.loaded {
    opacity: 0;
  }
`
