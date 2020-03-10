import * as React from 'react'

import * as PropTypes from 'prop-types'
import { useRef, useState } from 'react'
import { TokenCopyView } from './TokenCopy.view'

type TokenCopyProps = {
  token: string
}

export const TokenCopy = ({ token }: TokenCopyProps) => {
  const [copySuccess, setCopySuccess] = useState('')
  const textAreaRef = useRef(null)

  const copyToClipboard = () => {
    // @ts-ignore
    textAreaRef.current.select()
    document.execCommand('copy')
    setCopySuccess('Copied!')
  }

  return (
    <TokenCopyView
      token={token}
      copyToClipboard={copyToClipboard}
      copySuccess={copySuccess}
      textAreaRef={textAreaRef}
    />
  )
}

TokenCopy.propTypes = {
  token: PropTypes.string
}

TokenCopy.defaultProps = {
  token: undefined
}
