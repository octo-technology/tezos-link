import * as React from 'react'
import { useDispatch, useSelector } from 'react-redux'

import { HeaderView } from './Header.view'

export const Header = () => {
  const dispatch = useDispatch()
  const user = useSelector((state: any) => (state && state.auth ? state.auth.user : {}))
  const router = useSelector((state: any) => state.router)

  function removeAuthUserCallback() {
    dispatch(null)
  }

  return <HeaderView user={user} router={router} removeAuthUserCallback={removeAuthUserCallback} />
}
