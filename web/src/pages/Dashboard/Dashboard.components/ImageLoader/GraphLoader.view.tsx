import * as PropTypes from 'prop-types'
import * as React from 'react'

import { GraphLoaderImg, GraphLoaderStyled } from './ImageLoader.style'

type ImageLoaderViewProps = { loaded: boolean }

export const GraphLoaderView = ({ loaded }: ImageLoaderViewProps) => (
  <GraphLoaderStyled>
    <GraphLoaderImg className={loaded ? 'loaded' : ''} />
  </GraphLoaderStyled>
)

GraphLoaderView.propTypes = {
  loaded: PropTypes.bool.isRequired
}

GraphLoaderView.defaultProps = {
  style: {}
}
