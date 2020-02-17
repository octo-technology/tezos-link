import { SET_POSTS } from './Dashboard.actions'

const initialState = {
  posts: []
}

export function postReducers(state = initialState, action: any) {
  switch (action.type) {
    case SET_POSTS:
      return {
        ...state,
        posts: action.posts
      }
    default:
      return state
  }
}
