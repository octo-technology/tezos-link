import * as React from 'react'
import {
  RequestsByDayLineViewCard,
  RequestsByDayLineViewStyled,
  RequestsByDayLineViewTitle
} from './RequestsByDayLine.style'
import * as PropTypes from 'prop-types'
import { ResponsiveLine } from '@nivo/line'
import { RequestByDay } from '../../../../entities/ProjectWithMetrics'

type RequestsByDayLineViewProps = { requestsByDay: RequestByDay[] }

export const RequestsByDayLineView = ({ requestsByDay }: RequestsByDayLineViewProps) => {
  const requestsByDayGraph =
    requestsByDay && Array.isArray(requestsByDay)
      ? requestsByDay.map(requestByDay => {
        return {
          x: requestByDay.date,
          y: requestByDay.value
        }
      })
      : []

  const data = [
    {
      id: 'Requests',
      data: requestsByDayGraph
    }
  ]

  const hasRequests =
    requestsByDay.filter(requestByDay => {
      return requestByDay.value > 10
    }).length > 0

  return (
    <RequestsByDayLineViewStyled>
      <RequestsByDayLineViewCard>
        <RequestsByDayLineViewTitle>Requests last 30 days</RequestsByDayLineViewTitle>
        <ResponsiveLine
          data={data}
          margin={{ top: 30, right: 10, bottom: 55, left: 40 }}
          xScale={{
            type: 'time',
            format: '%Y-%m-%d',
            precision: 'day'
          }}
          xFormat="time:%Y-%m-%d"
          yScale={{
            type: 'linear',
            stacked: false,
            max: hasRequests ? 'auto' : 10
          }}
          axisLeft={{
            legendOffset: 12
          }}
          axisBottom={{
            format: '%b %d',
            tickValues: 'every 2 days',
            legendOffset: -12
          }}
          enablePointLabel={true}
          pointSize={5}
          pointBorderWidth={1}
          pointBorderColor={{
            from: 'color',
            modifiers: [['darker', 0.3]]
          }}
          useMesh={true}
          enableSlices={false}
          enableArea={true}
          enableGridX={false}
        />
      </RequestsByDayLineViewCard>
    </RequestsByDayLineViewStyled>
  )
}

RequestsByDayLineView.propTypes = {
  token: PropTypes.number
}

RequestsByDayLineView.defaultProps = {
  token: undefined
}
