import * as React from 'react'
import {
  StatusViewContent,
  StatusViewHeader,
  StatusViewIndicatorGreen,
  StatusViewIndicatorRed,
  StatusViewStyled,
  StatusViewSubtitle,
  StatusViewTitle
} from './Status.style'

import * as PropTypes from 'prop-types'

type StatusViewProps = {
  proxyStatus: boolean
  nodeStatus: boolean
  date: string
}

export const StatusView = ({ proxyStatus, nodeStatus, date }: StatusViewProps) => (
  <StatusViewStyled>
    <StatusViewContent>
      <h2>Services status</h2>
      <StatusViewHeader>
        {proxyStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Proxy service is {proxyStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {nodeStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Node RPC services are {nodeStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
    </StatusViewContent>
  </StatusViewStyled>
)

StatusView.propTypes = {
  proxyStatus: PropTypes.bool,
  nodeStatus: PropTypes.bool,
  date: PropTypes.string
}

StatusView.defaultProps = {
  proxyStatus: true,
  nodeStatus: true,
  date: ''
}
