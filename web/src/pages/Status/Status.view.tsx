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
  nodeMainnetArchiveStatus: boolean
  nodeMainnetRollingStatus: boolean
  nodeCarthagenetArchiveStatus: boolean
  nodeCarthagenetRollingStatus: boolean
  date: string
}

export const StatusView = ({ proxyStatus, nodeMainnetArchiveStatus, nodeMainnetRollingStatus, nodeCarthagenetArchiveStatus, nodeCarthagenetRollingStatus, date }: StatusViewProps) => (
  <StatusViewStyled>
    <StatusViewContent>
      <h2>Services status</h2>
      <StatusViewHeader>
        {proxyStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Proxy service is {proxyStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {nodeMainnetArchiveStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Archive-node RPC services for Alphanet are {nodeMainnetArchiveStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {nodeMainnetRollingStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Rolling-node RPC services for Alphanet are {nodeMainnetRollingStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {nodeCarthagenetArchiveStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Archive-node RPC services for Carthagenet Testnet are {nodeCarthagenetArchiveStatus ? 'online' : 'offline'}.</StatusViewTitle>
        <StatusViewSubtitle>As of {date}.</StatusViewSubtitle>
      </StatusViewHeader>
      <StatusViewHeader>
        {nodeCarthagenetRollingStatus ? <StatusViewIndicatorGreen /> : <StatusViewIndicatorRed />}
        <StatusViewTitle>Rolling-node RPC services for Carthagenet Testnet are {nodeCarthagenetRollingStatus ? 'online' : 'offline'}.</StatusViewTitle>
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
