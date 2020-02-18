import * as PropTypes from 'prop-types'
import * as React from 'react'

import { ImageLoaderImg, ImageLoaderStyled } from './ImageLoader.style'

type ImageLoaderViewProps = { src: string; alt: string; style: object; loaded: boolean }

export const ImageLoaderView = ({ src, alt, style, loaded }: ImageLoaderViewProps) => (
  <ImageLoaderStyled>
    <ImageLoaderImg src={src} alt={alt} style={style} className={loaded ? 'loaded' : ''} />
  </ImageLoaderStyled>
)

ImageLoaderView.propTypes = {
  src: PropTypes.string.isRequired,
  alt: PropTypes.string.isRequired,
  loaded: PropTypes.bool.isRequired,
  style: PropTypes.object
}

ImageLoaderView.defaultProps = {
  style: {}
}
