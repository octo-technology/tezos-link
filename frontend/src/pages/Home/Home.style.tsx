import styled from 'styled-components/macro'

import { AnimatedCard, FadeInFromTop } from '../../styles'

export const Error404Styled = styled.div`
  margin: 100px auto 20px auto;
`

export const Error404Card = styled(AnimatedCard)`
  width: 400px;
  max-width: 100vw;
  padding: 20px;
  text-align: center;
`

export const Error404Title = styled(FadeInFromTop)``

export const Error404Message = styled.div`
  margin-bottom: 18px;
`
