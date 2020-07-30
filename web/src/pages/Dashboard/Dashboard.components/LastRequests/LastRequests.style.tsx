import styled from 'styled-components/macro'
import { Card, FadeInFromLeft, primaryColor, secondaryColor } from 'src/styles'

export const LastRequestsViewCard = styled(Card)`
  margin: 0 0 20px 0;
  padding: 20px 20px 15px 10px;
  height: 300px;
  max-width: calc(100vw - 40px);
`

export const LastRequestsViewStyled = styled(FadeInFromLeft)`
  height: 300px;
  width: inherit;
`

export const LastRequestsViewTitle = styled.span`
  text-transform: uppercase;
  font-size: large;
  margin-left: 10px;
`

export const LastRequestsViewNoData = styled.div`
  font-size: 1em;
  color: grey;
  position: relative;
  top: 40%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
  text-align: center;
  margin: auto;
`

export const LastRequestsViewList = styled.ol`
  list-style: none;
  counter-reset: my-counter;
  padding-inline-start: 20px;
  overflow: hidden;
`

export const LastRequestsViewListItem = styled.li`
  margin: 5px 5px 5px 0px;
  padding: 5px 5px 5px 5px;
  counter-increment: my-counter;
  list-style-position: inside;

  -space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  background-color: ${secondaryColor};
  border-radius: 2px;
  color: #d4d4d4;
  width: 500px;

  &::before {
    content: counter(my-counter) '. ';
    color: ${primaryColor};
    font-weight: bold;
    margin-right: 5px;
    font-size: 1.3em;
  }
  &:first-child {
    font-size: 1.5em;
  }
`

export const LastRequestsViewListItemWithTooltip = styled.span`
  overflow-y: hidden;
  text-overflow: ellipsis;
  position: fixed;
  margin: 5px 0 0 5px;
  display: inline-block;
`
