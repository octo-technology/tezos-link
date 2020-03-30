import styled from 'styled-components/macro'

import { FadeInFromTop } from '../../styles'

export const HomeViewStyled = styled.div`
  margin: 200px auto 20px auto;
`

export const HomeViewTitle = styled(FadeInFromTop)`
  text-align: center;

  > h1 {
    margin: 20px 0 0 0;
  }
`

export const HomeText = styled(FadeInFromTop)`
  text-align: center;
  margin-top: 10px;
`

export const CreateButton = styled(FadeInFromTop)`
  width: max-content;
  margin: 30px auto;
`
