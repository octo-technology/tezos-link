import styled from 'styled-components/macro'

import { Card, FadeInFromLeft } from '../../styles'

export const StatusViewStyled = styled(FadeInFromLeft)`
  width: 50%;
  max-width: 840px;
  margin: 100px auto 20px auto;
`

export const StatusViewContent = styled(Card)`
  padding: 20px 20px 10px 20px;
`

export const StatusViewHeader = styled.div`
  padding: 20px 40px 40px 40px;
`

export const StatusViewIndicatorGreen = styled.div`
  display: inline-block;
  width: 15px;
  height: 15px;
  border-radius: 100%;
  margin-right: 20px;
  background: #17d766;
  opacity: 0.8;
  filter: alpha(opacity=80);
`

export const StatusViewIndicatorRed = styled.div`
  display: inline-block;
  width: 15px;
  height: 15px;
  border-radius: 100%;
  margin-right: 20px;
  background: #ff0000;
  opacity: 0.8;
  filter: alpha(opacity=80);
`

export const StatusViewTitle = styled.div`
  font-size: 1.5em;
  display: inline-block;
`

export const StatusViewSubtitle = styled.div`
  display: block;
  margin-left: 35px;
`
