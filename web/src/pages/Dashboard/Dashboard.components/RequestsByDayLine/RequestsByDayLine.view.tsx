import * as React from 'react'
import {
  RequestsByDayLineViewCard,
  RequestsByDayLineViewStyled,
  RequestsByDayLineViewTitle
} from './RequestsByDayLine.style'
import * as PropTypes from 'prop-types'
import { ResponsiveLine } from '@nivo/line'
import { RequestByDay } from '../../../../entities/ProjectWithMetrics'
import { primaryColor, secondaryColor, tertiaryColor } from 'src/styles'

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

  const theme = {
    axis: {
      fontSize: '14px',
      tickColor: primaryColor,
      ticks: {
        line: {
          stroke: tertiaryColor
        },
        text: {
          fill: tertiaryColor
        }
      },
      legend: {
        text: {
          fill: tertiaryColor
        }
      }
    },
    grid: {
      line: {
        stroke: '#555555'
      }
    }
  }

  const hasRequests =
    requestsByDay.filter(requestByDay => {
      return requestByDay.value > 10
    }).length > 0

  const isMobile = Math.max(document.documentElement.clientWidth, window.innerWidth || 0) <= 700

  return (
    <RequestsByDayLineViewStyled>
      <RequestsByDayLineViewCard>
        <RequestsByDayLineViewTitle>Requests last 30 days</RequestsByDayLineViewTitle>
        <ResponsiveLine
          data={data}
          margin={{ top: 30, right: 10, bottom: 55, left: 40 }}
          colors={[primaryColor, secondaryColor, tertiaryColor]}
          theme={theme}
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
            format: isMobile ? '%d' : '%b %d',
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
