import styled from 'styled-components/macro'
import { Card, FadeInFromLeft, secondaryColor } from 'src/styles'

export const RequestsByDayLineViewCard = styled(Card)`
  margin-bottom: 20px;
  padding: 20px 20px 15px 10px;
  width: 100%;
  height: 300px;
  color: ${secondaryColor};
`

export const RequestsByDayLineViewStyled = styled(FadeInFromLeft)``

export const RequestsByDayLineViewTitle = styled.span`
  text-transform: uppercase;
  font-size: large;
  margin-left: 10px;
  color: white;
`
