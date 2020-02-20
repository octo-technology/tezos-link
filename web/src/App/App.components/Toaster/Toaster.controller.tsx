import * as React from 'react'
import { useSelector } from 'react-redux'

import { ToasterView } from './Toaster.view'

export const Toaster = () => {
  const toaster = useSelector((state: any) => state.toaster)
  return <ToasterView status={toaster.status} title={toaster.title} message={toaster.message} />
}
