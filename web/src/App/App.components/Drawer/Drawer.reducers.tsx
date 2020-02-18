import { SHOW_DRAWER, HIDE_DRAWER } from './Drawer.actions'

const initialState = {
  showing: false
}

export function drawerReducers(state = initialState, action: any) {
  switch (action.type) {
    case SHOW_DRAWER:
      return {
        ...state,
        showing: true
      }
    case HIDE_DRAWER:
      return {
        ...state,
        showing: false
      }
    default:
      return state
  }
}
