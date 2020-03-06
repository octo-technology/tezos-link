import styled from 'styled-components/macro'
import { secondaryColor, tertiaryColor } from '../../../styles'

export const FooterStyled = styled.div`
  padding: 24px 50px;
  text-align: center;
  background-color: ${secondaryColor};
  color: ${tertiaryColor};
  a {
    color: ${tertiaryColor} !important;
  }
`
