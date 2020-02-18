import { SHOW_TOASTER, HIDE_TOASTER } from './Toaster.actions'

const initialState = {
  status: undefined,
  title: undefined,
  message: undefined
}

export function toasterReducers(state = initialState, action: any) {
  switch (action.type) {
    case SHOW_TOASTER:
      return {
        ...state,
        status: action.status,
        title: action.title,
        message: action.message
      }
    case HIDE_TOASTER:
      return {
        ...state,
        status: undefined
      }
    default:
      return state
  }
}
