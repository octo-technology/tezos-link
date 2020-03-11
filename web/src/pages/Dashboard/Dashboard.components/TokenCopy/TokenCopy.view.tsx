import * as React from 'react'

import { Input } from '../../../../App/App.components/Input/Input.controller'
import * as PropTypes from 'prop-types'
import { Button } from '../../../../App/App.components/Button/Button.controller'
import { RefObject } from 'react'

type TokenCopyViewProps = {
  token: string
  copySuccess: string
  copyToClipboard: () => void
  textAreaRef: RefObject<HTMLInputElement>
}

export const TokenCopyView = ({ token, copySuccess, copyToClipboard, textAreaRef }: TokenCopyViewProps) => {
  return (
    <>
      <Input inputRef={textAreaRef} icon="password" value={token} name="token" onChange={() => {}} onBlur={() => {}} />
      {/* Logical shortcut for only displaying the
             button if the copy command exists */
        document.queryCommandSupported('copy') && (
          <span>
            <Button onClick={copyToClipboard} text={copySuccess || 'Copy to clipboard'} icon="copy" type="button" />
          </span>
        )}
    </>
  )
}

TokenCopyView.propTypes = {
  token: PropTypes.string
}

TokenCopyView.defaultProps = {
  token: undefined
}
