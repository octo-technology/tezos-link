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

export const HeaderMenu = styled.div`
  position: absolute;
  height: 30px;
  top: 0;
  left: 20px;
  line-height: 64px;
  display: grid;
  grid-template-columns: 205px 80px 80px;
  grid-gap: 10px;

  > a > img {
    padding-top: 15px;
  }
`

export const HeaderLoggedOut = styled.div`
  position: absolute;
  top: 14px;
  right: 14px;
  display: grid;
  grid-template-columns: 130px 120px;
  grid-gap: 10px;
`
