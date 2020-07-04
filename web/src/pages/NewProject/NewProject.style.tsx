import styled from 'styled-components/macro'

import { Card, fadeInFromLeft } from '../../styles'

export const NewProjectCard = styled(Card)`
  margin-top: 150px;
  padding: 20px;
  width: 100%;
  height: 100%;
  max-width: 500px;
  will-change: transform, opacity;
  animation: ${fadeInFromLeft} 500ms;

  > h1 {
    margin-top: 0;
  }

  button {
    margin-top: 10px;
  }
`
