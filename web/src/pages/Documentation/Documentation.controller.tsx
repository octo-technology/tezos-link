import * as React from 'react'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { DocumentationView } from './Documentation.view'

export const Documentation = () => {
  const [menu, setMenu] = useState('')
  const [content, setContent] = useState('')

  useEffect(() => {
    axios({
      method: 'get',
      url: 'docs/menu.md'
    })
      .then(function(response: any) {
        setMenu(response.data)
      })
      .catch(function(error: any) {
        console.error(error)
      })

    axios({
      method: 'get',
      url: 'docs/content.md'
    })
      .then(function(response: any) {
        setContent(response.data)
      })
      .catch(function(error: any) {
        console.error(error)
      })
  })

  return (
    <>
      <DocumentationView content={content} menu={menu} />
    </>
  )
}
