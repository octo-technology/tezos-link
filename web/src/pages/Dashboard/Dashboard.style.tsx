import { FadeInFromTop, FadeInFromRight } from 'src/styles'
import styled from 'styled-components/macro'

export const DashboardStyled = styled.div`
  width: 80%;
  max-width: 1140px;
  margin: 100px auto 20px auto;
`

export const DashboardTopStyled = styled.div`
  display: grid;
  grid-template-columns: calc(100% - 320px) 300px;
  grid-gap: 20px;
`

export const DashboardBottomStyled = styled.div`
  display: grid;
  grid-template-columns: 49% 49%;
  grid-gap: 20px;
`

export const DashboardTitle = styled.div``

export const DashboardLeftSide = styled.div``

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
