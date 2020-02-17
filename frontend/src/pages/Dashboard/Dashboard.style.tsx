import { FadeInFromTop, FadeInFromRight, FadeInFromLeft } from 'src/styles'
import styled from 'styled-components/macro'

export const DashboardStyled = styled.div`
  width: 80%;
  max-width: 1140px;
  margin: 100px auto 20px auto;
  display: grid;
  grid-template-columns: calc(75% - 20px) 25%;
  grid-gap: 20px;
`

export const DashboardTitle = styled.div``

export const DashboardLeftSide = styled.div``

export const DashboardPostRows = styled(FadeInFromLeft)``

export const DashboardRightSide = styled(FadeInFromRight)``

export const DashboardHeader = styled(FadeInFromTop)`
  display: grid;
  grid-template-columns: calc(100% - 150px) 150px;
  margin: 20px 0;

  > div > h1 {
    margin: 0;
    line-height: 40px;
  }
`
