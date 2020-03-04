import { SHOW_PROGRESS_BAR, HIDE_PROGRESS_BAR } from './ProgressBar.actions'

const initialState = {
  loading: false
}

export function progressBarReducers(state = initialState, action: any) {
  if (action.type === SHOW_PROGRESS_BAR) {
    return {
      ...state,
      loading: true
    }
  } else if (action.type === HIDE_PROGRESS_BAR) {
    return {
      ...state,
      loading: false
    }
  } else {
    return state
  }
}
