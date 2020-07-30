import * as React from 'react'
// prettier-ignore
import {
  RPCUsagePieViewCard,
  RPCUsagePieViewStyled,
  RPCUsagePieViewTitle,
  RPCUsagePieViewNoData,
  RPCUsagePieViewTotalRequestsCount,
  RPCUsagePieViewTotalRequestsCountNumber,
  RPCUsagePieContainer
} from './RPCUsagePie.style'
import * as PropTypes from 'prop-types'
import { RPCUsage } from '../../../../entities/ProjectWithMetrics'
import { ResponsivePie, PieDatum } from '@nivo/pie'

type RPCUsagePieViewProps = { rpcUsage: RPCUsage[]; rpcTotalCount: number }

export const RPCUsagePieView = ({ rpcUsage, rpcTotalCount }: RPCUsagePieViewProps) => {
  const rpcUsagePieData =
    rpcUsage && Array.isArray(rpcUsage)
      ? rpcUsage.map(rpc => {
          const rpcPieUsage: PieDatum = {
            id: rpc.id,
            label: rpc.label,
            value: rpc.value,
            color: 'hsl(177, 78%, 58%)'
          }

          return rpcPieUsage
        })
      : []

  return (
    <RPCUsagePieViewStyled>
      <RPCUsagePieViewCard>
        <RPCUsagePieViewTitle>RPC Usage</RPCUsagePieViewTitle>
        {rpcUsagePieData.length > 0 ? (
          <RPCUsagePieViewTotalRequestsCount>
            Total requests <br />
            <RPCUsagePieViewTotalRequestsCountNumber>{rpcTotalCount}</RPCUsagePieViewTotalRequestsCountNumber>
          </RPCUsagePieViewTotalRequestsCount>
        ) : null}
        {rpcUsagePieData.length > 0 ? (
          <RPCUsagePieContainer>
            <ResponsivePie
              data={rpcUsagePieData}
              margin={{ top: 20, left: 20, bottom: 20, right: 20 }}
              innerRadius={0.7}
              padAngle={0.7}
              cornerRadius={3}
              colors={{ scheme: 'nivo' }}
              borderWidth={1}
              borderColor={{ theme: 'background' }}
              enableRadialLabels={false}
              radialLabelsLinkColor={{ from: 'color' }}
              slicesLabelsSkipAngle={10}
              slicesLabelsTextColor="#333333"
              animate={true}
              sortByValue={true}
              motionStiffness={90}
              motionDamping={15}
            />
          </RPCUsagePieContainer>
        ) : (
          <RPCUsagePieViewNoData>No data</RPCUsagePieViewNoData>
        )}
      </RPCUsagePieViewCard>
    </RPCUsagePieViewStyled>
  )
}

RPCUsagePieView.propTypes = {
  token: PropTypes.number
}

RPCUsagePieView.defaultProps = {
  token: undefined
}
