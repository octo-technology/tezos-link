import * as React from 'react'
import { RequestsCounterCard, BigNumber, RequestsCounterStyled } from './RequestsCounter.style'
import * as PropTypes from 'prop-types'

type RequestsCounterViewProps = {
  count: number
}

export const RequestsCounterView = ({ count }: RequestsCounterViewProps) => {
  return (
    <RequestsCounterStyled>
      <RequestsCounterCard>
        <BigNumber>{count}</BigNumber> requests
      </RequestsCounterCard>
    </RequestsCounterStyled>
  )
}

RequestsCounterView.propTypes = {
  token: PropTypes.number
}

RequestsCounterView.defaultProps = {
  token: undefined
}
