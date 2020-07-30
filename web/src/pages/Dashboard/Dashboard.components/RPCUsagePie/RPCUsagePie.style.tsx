import styled from 'styled-components/macro'
import { Card, FadeInFromLeft } from 'src/styles'

export const RPCUsagePieViewCard = styled(Card)`
  padding: 20px 20px 15px 10px;
  height: 300px;
  margin: 0 0 20px 0;
  max-width: calc(100vw - 40px);
`

export const RPCUsagePieContainer = styled.span`
  color: grey;
`

export const RPCUsagePieViewStyled = styled(FadeInFromLeft)`
  position: relative;
  float: left;
`

export const RPCUsagePieViewTitle = styled.span`
  text-transform: uppercase;
  font-size: large;
  margin-left: 10px;
`

export const RPCUsagePieViewNoData = styled.div`
  font-size: 1em;
  color: grey;
  position: relative;
  top: 45%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
  text-align: center;
  margin: auto;
`

export const RPCUsagePieViewTotalRequestsCount = styled.div`
  text-transform: uppercase;
  color: grey;
  position: absolute;
  top: 56%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
  text-align: center;
  width: inherit;
  left: calc(50% - 62px);
`

export const RPCUsagePieViewTotalRequestsCountNumber = styled.div`
  font-size: 3em;
  color: grey;
`
