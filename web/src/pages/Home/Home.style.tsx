import styled from 'styled-components/macro'
import { backgroundColor2, backgroundColor3, primaryColor } from 'src/styles'

export const HomeStyled = styled.div``

export const HomeSection = styled.div`
  width: 90vw;
  max-width: 1280px;
  margin: 100px auto 20px auto;
`

export const HomeHead = styled(HomeSection)`
  display: grid;
  grid-template-columns: 1fr 1fr;
  text-align: left;
  height: 700px;
  position: relative;

  h1 {
    margin: 20px 0 10px 0;
  }

  button {
    margin-top: 40px;
    width: 200px;
  }

  @media (max-width: 900px) {
    grid-template-columns: 1fr;
  }
`

export const HomeLeft = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-around;
`

export const HomeLeftInside = styled.div`
  height: 400px;
`

export const HomeRight = styled.div`
  position: relative;

  @media (max-width: 900px) {
    display: none;
  }

  > img {
    position: absolute;
    bottom: 100px;
    left: 100px;
  }

  @keyframes breathing {
    0% {
      opacity: 0.65;
    }

    25% {
      opacity: 1;
    }

    50% {
      opacity: 0.65;
    }

    75% {
      opacity: 0.3;
    }

    100% {
      opacity: 0.65;
    }
  }

  .torus-cables {
    z-index: 0;
  }

  .torus-bg {
    z-index: 1;
  }

  .torus-fg {
    z-index: 3;
  }

  .torus-logo {
    z-index: 4;
  }

  .torus-logo,
  .torus-cables {
    animation: breathing 3s linear infinite normal;
  }

  @keyframes shooting {
    0% {
      transform: translate(0px, 171px);
      opacity: 0;
    }

    20% {
      opacity: 1;
    }

    80% {
      transform: translate(300px, 0px);
      opacity: 1;
    }

    100% {
      transform: translate(300px, 0px);
      opacity: 0;
    }
  }

  .meteor {
    width: 165px;
    height: 100px;
    background-image: url('/images/particle.svg');
    z-index: 2;
    position: absolute;
    opacity: 0;
    animation: shooting 1.5s linear infinite;

    &.meteor1 {
      bottom: 350px;
      left: -70px;
      animation-delay: 0s;
    }

    &.meteor2 {
      bottom: 300px;
      left: -70px;
      animation-delay: 0.3s;
    }

    &.meteor3 {
      bottom: 400px;
      left: -70px;
      animation-delay: 0.5s;
    }

    &.meteor4 {
      bottom: 380px;
      left: -70px;
      animation-delay: 0.9s;
    }
  }
`

export const HomeBrought = styled.img`
  position: absolute;
  bottom: 0px;
  left: calc(50% - 132px);
`

export const HomeBuilt = styled.div`
  background-color: ${backgroundColor2};
  text-align: center;
`

export const HomeH1 = styled.h1`
  margin: 0 auto;
  padding: 70px 0 0 0;
`

export const HomeH3 = styled.h3`
  margin: 10px auto 30px auto;
  padding: 0;
`

export const HomeBuiltArg = styled.div`
  text-align: left;
  display: flex;
  align-items: center;
  justify-content: space-around;

  @media (max-width: 900px) {
    display: grid;
    grid-template-columns: auto;
    justify-content: initial;
    margin: 40px 0;
  }
`

export const HomeBuiltArgRev = styled.div`
  text-align: left;
  display: flex;
  align-items: center;
  justify-content: space-around;

  @media (max-width: 900px) {
    display: grid;
    grid-template-columns: auto;
    justify-content: initial;
    margin: 40px 0;
  }
`

export const HomeBuiltArgTitle = styled.div`
  font-weight: bold;
  font-size: 26px;
  line-height: 32px;
  color: ${primaryColor};
`

export const HomeBuiltArgText = styled.div`
  font-size: 18px;
  line-height: 22px;
  text-align: justify;
  font-weight: 300;
`

export const HomeBuiltArgImg = styled.img`
  width: 477px;
  height: 477px;

  @media (max-width: 900px) {
    width: 300px;
    height: 300px;
  }
`

export const HomeTrusted = styled(HomeSection)`
  text-align: center;
  margin-top: 0;

  button {
    margin: 100px auto;
    width: 350px;

    @media (max-width: 900px) {
      max-width: 90vw;
    }
  }
`

export const HomeTrustedGrid = styled(HomeSection)`
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: space-around;

  img {
    max-width: 200px;
  }

  @media (max-width: 900px) {
    display: grid;
    grid-gap: 20px;
    grid-template-columns: auto;
    justify-content: initial;

    img {
      max-width: 70vw;
      margin: auto;
    }
  }
`

export const HomeMetrics = styled.div`
  text-align: center;
  background-color: ${backgroundColor2};

  > img {
    width: 1139px;
    max-width: 90vw;
    margin: 60px auto 100px auto;
  }
`

export const HomePanels = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 40px;
  width: 1140px;
  max-width: 90vw;
  margin: 50px auto;

  @media (max-width: 900px) {
    display: none;
  }
`

export const HomePanel = styled.div`
  height: 200px;
  max-width: 90vw;
  background-color: ${primaryColor};
  display: grid;
  grid-template-columns: 70px auto 30px;
  grid-gap: 40px;
  padding: 65px 45px;
  color: #000;

  h3 {
    font-size: 20px;
    margin-top: 10px;
    font-weight: 600;
  }

  p {
    font-size: 16px;
    margin: 0;
  }

  .logo {
    width: 70px;
  }

  .arrow {
    margin-top: 17px;
    width: 30px;
  }
`

export const HomePanel2 = styled.div`
  height: 200px;
  max-width: 90vw;
  background-color: ${backgroundColor2};
  display: grid;
  grid-template-columns: 70px auto 30px;
  grid-gap: 40px;
  padding: 65px 45px;

  h3 {
    font-size: 20px;
    margin-top: 10px;
    font-weight: 600;
  }

  p {
    font-size: 16px;
    margin: 0;
  }

  .logo {
    width: 70px;
  }

  .arrow {
    margin-top: 17px;
    width: 30px;
  }
`

export const HomeGetStarted = styled.div`
  background: url('/images/grid-bg.svg'), ${backgroundColor2};
  text-align: center;
  position: relative;
  height: 300px;

  h1 {
    padding-top: 100px;
  }

  img.left {
    position: absolute;
    left: 0;
    top: 0;
    height: 300px;
  }

  img.right {
    position: absolute;
    right: 0;
    top: 0;
    height: 300px;
    transform: scaleX(-1);
  }

  @media (max-width: 900px) {
    display: none;
  }

  button {
    margin-top: 20px;
    width: 200px;
  }
`

export const HomeFooter = styled.div`
  background-color: ${backgroundColor3};
  width: 100%;
`

export const HomeFooterGrid = styled.div`
  width: 100%;
  max-width: 1280px;
  margin: auto;
  padding: 50px 0;
  display: grid;
  grid-template-columns: 2fr 1fr 1fr 1fr 2fr;
  grid-gap: 30px;
  font-family: 'Proxima Nova';

  a {
    text-decoration: underline !important;
    display: block;
    line-height: 20px;
  }

  @media (max-width: 900px) {
    grid-template-columns: auto;
    text-align: center;

    img {
      margin: auto;
    }
  }
`
