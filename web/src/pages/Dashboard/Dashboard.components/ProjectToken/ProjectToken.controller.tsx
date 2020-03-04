import * as React from 'react'

import { ProjectTokenView } from './ProjectToken.view'
import * as PropTypes from 'prop-types'
import { useRef, useState } from 'react'

type ProjectTokenProps = {
  token: string
}

export const ProjectToken = ({ token }: ProjectTokenProps) => {
  const [copySuccess, setCopySuccess] = useState('')
  const textAreaRef = useRef(null)

  const copyToClipboard = () => {
    // @ts-ignore
    textAreaRef.current.select()
    document.execCommand('copy')
    setCopySuccess('Copied!')
  }

  return (
    <ProjectTokenView
      token={token}
      copyToClipboard={copyToClipboard}
      copySuccess={copySuccess}
      textAreaRef={textAreaRef}
    />
  )
}

ProjectToken.propTypes = {
  token: PropTypes.string
}

ProjectToken.defaultProps = {
  token: undefined
}
