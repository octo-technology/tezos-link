import { SHOW_TOASTER, HIDE_TOASTER } from './Toaster.actions'

const initialState = {
  status: undefined,
  title: undefined,
  message: undefined
}

export function toasterReducers(state = initialState, action: any) {
  if (action.type === SHOW_TOASTER) {
    return {
      ...state,
      status: action.status,
      title: action.title,
      message: action.message
    }
  } else if (action.type === HIDE_TOASTER) {
    return {
      ...state,
      status: undefined
    }
  } else {
    return state
  }
}
