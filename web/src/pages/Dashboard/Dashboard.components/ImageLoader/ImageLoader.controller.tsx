import * as React from 'react'
import * as PropTypes from 'prop-types'
import { useEffect, useState } from 'react'

import { ImageLoaderView } from './ImageLoader.view'

type ImageLoaderProps = { src: string; alt: string; style: object }

export const ImageLoader = ({ src, alt, style }: ImageLoaderProps) => {
  const [loaded, setLoaded] = useState(false)

  useEffect(() => {
    let imageToDisplay
    imageToDisplay = new Image()
    imageToDisplay.src = src
    imageToDisplay.alt = alt
    imageToDisplay.onload = () => onImageLoad()
  })

  const onImageLoad = () => {
    setLoaded(true)
  }

  return <ImageLoaderView src={src} alt={alt} style={style} loaded={loaded} />
}

ImageLoader.propTypes = {
  src: PropTypes.string.isRequired,
  alt: PropTypes.string.isRequired,
  style: PropTypes.object
}

ImageLoader.defaultProps = {
  style: {}
}
