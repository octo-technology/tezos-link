import * as React from 'react'
import { useDispatch, useSelector } from 'react-redux'

import { DrawerView } from './Drawer.view'
import { hideDrawer } from './Drawer.actions'
//import { removeAuthUser } from 'src/pages/SignUp/SignUp.actions'

export const Drawer = () => {
  const dispatch = useDispatch()
  const showing = useSelector((state: any) => state.drawer && state.drawer.showing)
  const route = useSelector((state: any) => state.router && state.router.location && state.router.location.pathname)
  const user = useSelector((state: any) => state && state.auth && state.auth.user)

  const hideCallback = () => {
    dispatch(hideDrawer())
  }

  function removeAuthUserCallback() {
    //dispatch(removeAuthUser())
  }

  return (
    <DrawerView
      showing={showing}
      hideCallback={hideCallback}
      route={route}
      user={user}
      removeAuthUserCallback={removeAuthUserCallback}
    />
  )
}
