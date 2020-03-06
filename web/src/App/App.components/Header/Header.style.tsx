import styled from 'styled-components/macro'

export const HeaderStyled = styled.div`
  z-index: 1;
  height: 64px;
  background-color: white;
  position: fixed;
  top: 0;
  width: 100vw;
  box-shadow: 0 1px 10px rgba(151, 164, 175, 0.1);
`

export const HeaderLogo = styled.div`
  position: absolute;
  left: 20px;

  img {
    height: 30px;
    margin-top: 17px;
    filter: drop-shadow(0 1px 10px rgba(151, 164, 175, 0.1));
  }
`

export const HeaderMenu = styled.div`
  position: absolute;
  top: 14px;
  left: 60px;
  display: grid;
  grid-template-columns: 80px 80px 80px;
  grid-gap: 10px;
`

export const HeaderLoggedOut = styled.div`
  position: absolute;
  top: 14px;
  right: 14px;
  display: grid;
  grid-template-columns: 130px 120px;
  grid-gap: 10px;
`

export const Beta = styled.span`
  vertical-align: 25px;
  margin-left: 0px;
  color: grey;
  font-size: 0.8em;
`
