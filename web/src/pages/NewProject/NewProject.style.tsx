import styled from 'styled-components/macro'

import { Card, fadeInFromLeft } from '../../styles'

export const NewProjectCard = styled(Card)`
  margin-top: 10%;
  padding: 10px;
  width: 100%;
  height: 100%;
  max-width: 500px;
  will-change: transform, opacity;
  animation: ${fadeInFromLeft} 500ms;
`
