import styled from 'styled-components/macro'
import { backgroundColor, shadowColor } from 'src/styles'

export const HeaderStyled = styled.div`
  z-index: 10;
  height: 64px;
  background-color: ${backgroundColor};
  position: fixed;
  top: 0;
  width: 100vw;
  box-shadow: 0 6px 10px ${shadowColor}80;

  .header-triangle {
    position: absolute;
    top: 64px;
    left: calc(50vw - 70px);
    width: 0;
    height: 0;
    border-left: 70px solid transparent;
    border-right: 70px solid transparent;
    border-top: 40px solid ${backgroundColor};
    filter: drop-shadow(0 4px 5px ${shadowColor});
  }
`

export const HeaderMenu = styled.div`
  position: absolute;
  height: 64px;
  top: 4px;
  left: 60px;
  line-height: 64px;

  @media (max-width: 900px) {
    display: none;
  }
`

export const HeaderLogo = styled.div`
  position: absolute;
  top: 13px;
  left: calc(50vw - 30px);

  img {
    height: 68px;
    margin-top: 5px;
    filter: drop-shadow(0 4px 4px rgba(0, 0, 0, 0.25));
  }
`

export const HeaderLoggedOut = styled.div`
  position: absolute;
  top: 14px;
  right: 14px;
  display: grid;
  grid-template-columns: 130px 120px;
  grid-gap: 10px;

  @media (max-width: 900px) {
    display: none;
  }
`
