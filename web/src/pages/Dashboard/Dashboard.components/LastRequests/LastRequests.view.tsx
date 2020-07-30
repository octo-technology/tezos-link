import * as React from 'react'
import {
  LastRequestsViewCard,
  LastRequestsViewStyled,
  LastRequestsViewTitle,
  LastRequestsViewNoData,
  LastRequestsViewListItem,
  LastRequestsViewList,
  LastRequestsViewListItemWithTooltip
} from './LastRequests.style'
import * as PropTypes from 'prop-types'

type LastRequestsViewProps = { lastRequests: string[] }

const showRequests = (request: string, isFirst: boolean) => {
  const stringLength = isFirst ? 40 : 60
  if (request.length > stringLength) {
    return request.substring(0, stringLength) + '...'
  }

  return request
}

export const LastRequestsView = ({ lastRequests }: LastRequestsViewProps) => {
  return (
    <LastRequestsViewCard>
      <LastRequestsViewStyled>
        <LastRequestsViewTitle>Last requests</LastRequestsViewTitle>
        {lastRequests && lastRequests.length > 0 ? (
          <LastRequestsViewList>
            {lastRequests.slice(0, 5).map((request: string, index: number) => {
              const isFirst = index === 0
              return (
                <LastRequestsViewListItem key={index}>
                  <LastRequestsViewListItemWithTooltip data-tooltip={request}>
                    {showRequests(request, isFirst)}
                  </LastRequestsViewListItemWithTooltip>
                </LastRequestsViewListItem>
              )
            })}
          </LastRequestsViewList>
        ) : (
          <LastRequestsViewNoData>No data</LastRequestsViewNoData>
        )}
      </LastRequestsViewStyled>
    </LastRequestsViewCard>
  )
}

LastRequestsView.propTypes = {
  lastRequests: PropTypes.array
}

LastRequestsView.defaultProps = {
  lastRequests: undefined
}
