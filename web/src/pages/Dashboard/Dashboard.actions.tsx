export const SET_POSTS = 'SET_POSTS'

export const setPosts = (posts: [any]) => (dispatch: any) => {
  dispatch({
    type: SET_POSTS,
    posts
  })
}
