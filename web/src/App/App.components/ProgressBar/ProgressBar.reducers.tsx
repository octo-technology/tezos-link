import { SHOW_PROGRESS_BAR, HIDE_PROGRESS_BAR } from './ProgressBar.actions'

const initialState = {
  loading: false
}

export function progressBarReducers(state = initialState, action: any) {
  switch (action.type) {
    case SHOW_PROGRESS_BAR:
      return {
        ...state,
        loading: true
      }
    case HIDE_PROGRESS_BAR:
      return {
        ...state,
        loading: false
      }
    default:
      return state
  }
}
