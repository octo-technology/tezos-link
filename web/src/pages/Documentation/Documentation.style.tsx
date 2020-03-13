import styled from 'styled-components/macro'

import { Card, FadeInFromLeft, FadeInFromRight } from '../../styles'

export const DocumentationViewStyled = styled.div`
  width: 80%;
  max-width: 1140px;
  margin: 100px auto 20px auto;
  display: grid;
  grid-template-columns: 25% calc(75% - 20px);
  grid-gap: 20px;
`

export const DocumentationViewLeftSide = styled(FadeInFromLeft)``

export const DocumentationViewRightSide = styled(FadeInFromRight)``

export const DocumentationViewMenu = styled(Card)`
  margin-bottom: 20px;
  padding: 20px 20px 10px 20px;
`

export const DocumentationViewContent = styled(Card)`
  margin-bottom: 20px;
  padding: 20px 20px 10px 20px;
  overflow-wrap: break-word;
  word-wrap: break-word;
`
