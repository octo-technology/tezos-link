import * as React from 'react'
import { useDispatch } from 'react-redux'

import { hideProgressBar, showProgressBar } from '../../App/App.components/ProgressBar/ProgressBar.actions'
import { showToaster } from '../../App/App.components/Toaster/Toaster.actions'
import { ERROR } from '../../App/App.components/Toaster/Toaster.constants'
import { setPosts } from './Dashboard.actions'
import { GET_LATEST_POSTS } from './Dashboard.query'
import { DashboardView } from './Dashboard.view'

export const Dashboard = () => {
  const dispatch = useDispatch()

  // useEffect(() => {
  //   const handleScroll = (event: any) => {
  //     // TODO : Load more post on scroll
  //   }

  //   window.addEventListener('scroll', handleScroll)

  //   return () => window.removeEventListener('scroll', handleScroll)
  // }, [])

  dispatch(showProgressBar())
  let { postsLoading, postsError, posts, fetchMorePosts } = useGetPosts()

  if (postsError) { // @ts-ignore
    // @ts-ignore
    dispatch(showToaster(ERROR, 'Query Error', postsError.message))
  }

  if (!postsLoading && posts) {
    // @ts-ignore
    dispatch(setPosts(posts))
    dispatch(hideProgressBar())
  }

  function useFetchMoreCallback() {
    console.log('yo')
    useFetchMore(fetchMorePosts, posts)
  }

  // @ts-ignore
  return <DashboardView posts={posts} loading={postsLoading} fetchMoreCallback={useFetchMoreCallback} />
}

function useFetchMore(fetchMorePosts: any, posts: any) {
  fetchMorePosts({
    query: GET_LATEST_POSTS,
    variables: { cursorId: posts[posts.length - 1]._id },
    updateQuery: (prev: any, { fetchMoreResult }: any) => {
      if (!fetchMoreResult) return prev
      return Object.assign({}, prev, {
        getLatestPosts: [...prev.getLatestPosts, ...fetchMoreResult.getLatestPosts]
      })
    }
  })
}

function useGetPosts() {
  //const { loading, error, data, fetchMore } = {}
  return {
    postsLoading: null,
    postsError: null,
    posts: null,
    fetchMorePosts: null
  }
}
