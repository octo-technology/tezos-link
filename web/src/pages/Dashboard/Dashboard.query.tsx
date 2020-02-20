import gql from 'graphql-tag'

export const GET_LATEST_POSTS = gql`
  query GetLatestPosts($cursorId: ID) {
    getLatestPosts(cursorId: $cursorId) {
      user {
        _id
        username
        role
        createdAt
        updatedAt
      }
      vote {
        _id
        userId
        target
        targetId
        direction
        createdAt
        updatedAt
      }
      _id
      userId
      title
      slug
      content
      category
      sticky
      image
      viewCount
      replyCount
      status
      upCount
      downCount
      lastestReply
      createdAt
      updatedAt
      __typename
    }
  }
`
