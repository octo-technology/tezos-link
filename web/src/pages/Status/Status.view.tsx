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
  proxyMainnetStatus: boolean
  nodeMainnetArchiveStatus: boolean
  nodeMainnetRollingStatus: boolean
  proxyCarthagenetStatus: boolean
  nodeCarthagenetArchiveStatus: boolean
  nodeCarthagenetRollingStatus: boolean
  date: string
}

export const StatusView = ({ proxyMainnetStatus, nodeMainnetArchiveStatus, nodeMainnetRollingStatus, proxyCarthagenetStatus, nodeCarthagenetArchiveStatus, nodeCarthagenetRollingStatus, date }: StatusViewProps) => (
  <StatusViewStyled>
    <StatusViewContent>
      <h2>Services status</h2>
      <StatusViewHeader>
        {proxyMainnetStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Proxy service for Mainnet is {proxyMainnetStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        <StatusViewTitle>Nodes RPC services for Mainnet:</StatusViewTitle>
        <StatusViewSubtitle>{nodeMainnetArchiveStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />} Archive nodes are {nodeMainnetArchiveStatus ? 'online' : 'offline'}.</StatusViewSubtitle>
        <StatusViewSubtitle>{nodeMainnetRollingStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />} Rolling nodes are {nodeMainnetRollingStatus ? 'online' : 'offline'}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {proxyCarthagenetStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Proxy service for Carthagenet is {proxyCarthagenetStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        <StatusViewTitle>Nodes RPC services for Carthagenet:</StatusViewTitle>
        <StatusViewSubtitle>{nodeCarthagenetArchiveStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />} Archive nodes are {nodeCarthagenetArchiveStatus ? 'online' : 'offline'}.</StatusViewSubtitle>
        <StatusViewSubtitle>{nodeCarthagenetRollingStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />} Rolling nodes are {nodeCarthagenetRollingStatus ? 'online' : 'offline'}.</StatusViewSubtitle>
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
