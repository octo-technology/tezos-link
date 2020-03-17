import * as React from 'react'
import {
  RPCUsagePieViewCard,
  RPCUsagePieViewStyled,
  RPCUsagePieViewTitle,
  RPCUsagePieViewNoData,
  RPCUsagePieViewTotalRequestsCount, RPCUsagePieViewTotalRequestsCountNumber
} from "./RPCUsagePie.style";
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
          color: 'hsl(46, 87%, 62%)'
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
          <ResponsivePie
            data={rpcUsagePieData}
            margin={{ top: 20, left: -100, bottom: 30, right: 100 }}
            innerRadius={0.7}
            padAngle={0.7}
            cornerRadius={3}
            colors={{ scheme: 'nivo' }}
            borderWidth={1}
            borderColor={{ theme: 'background' }}
            enableRadialLabels={false}
            radialLabelsSkipAngle={10}
            radialLabelsTextXOffset={6}
            radialLabelsTextColor="#A9A9A9"
            radialLabelsLinkOffset={0}
            radialLabelsLinkDiagonalLength={16}
            radialLabelsLinkHorizontalLength={24}
            radialLabelsLinkStrokeWidth={1}
            radialLabelsLinkColor={{ from: 'color' }}
            slicesLabelsSkipAngle={10}
            slicesLabelsTextColor="#333333"
            animate={true}
            sortByValue={true}
            motionStiffness={90}
            motionDamping={15}
            legends={[
              {
                anchor: 'right',
                direction: 'column',
                itemWidth: 100,
                itemHeight: 18,
                itemTextColor: '#999',
                symbolSize: 18,
                symbolShape: 'circle',
                padding: { top: 5, right: 0, left: 0, bottom: 0 },
                effects: [
                  {
                    on: 'hover',
                    style: {
                      itemTextColor: '#000'
                    }
                  }
                ]
              }
            ]}
          />
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
