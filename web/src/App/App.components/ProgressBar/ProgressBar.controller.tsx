import * as React from 'react'
import { useState } from 'react'
import { useSelector } from 'react-redux'

import { ready, showing, done } from './ProgressBar.constants'
import { ProgressBarView } from './ProgressBar.view'

export const ProgressBar = () => {
  const progressBar = useSelector((state: any) => state.progressBar)
  const [status, setStatus] = useState(ready)

  if (status === ready && progressBar.loading) {
    setStatus(showing)
    setTimeout(() => setStatus(done), 30000)
  } else if (status === 'showing' && !progressBar.loading) {
    setStatus(done)
    setTimeout(() => setStatus(ready), 500)
  }

  return <ProgressBarView status={status} />
}
