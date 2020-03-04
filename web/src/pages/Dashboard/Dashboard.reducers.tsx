import { SET_METRICS } from './Dashboard.actions'

const initialState = {
  metrics: {}
}

export function metricsReducers(state = initialState, action: any) {
  if (action.type === SET_METRICS) {
    return {
      ...state,
      metrics: action.metrics
    }
  } else {
    return state
  }
}
