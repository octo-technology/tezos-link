import * as React from 'react'
import { RequestsCounterCard, BigNumber } from './RequestsCounter.style'
import * as PropTypes from 'prop-types'

type RequestsCounterViewProps = {
  count: number
}

export const RequestsCounterView = ({ count }: RequestsCounterViewProps) => {
  return (
    <RequestsCounterCard>
      <BigNumber>{count}</BigNumber>
      {' '} requests
    </RequestsCounterCard>
  )
}

RequestsCounterView.propTypes = {
  token: PropTypes.number
}

RequestsCounterView.defaultProps = {
  token: undefined
}
