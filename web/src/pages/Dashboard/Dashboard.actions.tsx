export const SET_METRICS = 'SET_METRICS'

export const setMetrics = (metrics: [any]) => (dispatch: any) => {
  dispatch({
    type: SET_METRICS,
    metrics: metrics
  })
}
