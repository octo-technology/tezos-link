import * as React from 'react'
import { useEffect, useState } from 'react'

import { GraphLoaderView } from './GraphLoader.view'

export const GraphLoader = () => {
  const [loaded, setLoaded] = useState(false)

  useEffect(() => {
    let imageToDisplay
    imageToDisplay = new Image()
    imageToDisplay.onload = () => onImageLoad()
  })

  const onImageLoad = () => {
    setLoaded(true)
  }

  return <GraphLoaderView loaded={loaded} />
}

GraphLoader.defaultProps = {
  style: {}
}
