import * as React from 'react'

import { ProjectTokenCard, ProjectTokenStyled, ProjectTokenTitle } from './ProjectToken.style'
import { Input } from '../../../../App/App.components/Input/Input.controller'
import * as PropTypes from 'prop-types'
import { Button } from '../../../../App/App.components/Button/Button.controller'
import { RefObject } from 'react'

type ProjectTokenViewProps = {
  token: string
  copySuccess: string
  copyToClipboard: () => void
  textAreaRef: RefObject<HTMLInputElement>
}

export const ProjectTokenView = ({ token, copySuccess, copyToClipboard, textAreaRef }: ProjectTokenViewProps) => {
  return (
    <ProjectTokenStyled>
      <ProjectTokenCard>
        <ProjectTokenTitle>private token</ProjectTokenTitle>
        <Input
          inputRef={textAreaRef}
          icon="password"
          value={token}
          name="token"
          onChange={() => {}}
          onBlur={() => {}}
        />
        {/* Logical shortcut for only displaying the
             button if the copy command exists */
          document.queryCommandSupported('copy') && (
            <div>
              <Button onClick={copyToClipboard} text={copySuccess || 'Copy to clipboard'} icon="copy" type="button" />
            </div>
          )}
        <p>
          Make sure to <b>save this token</b>, it is both your <b>access to this dashboard</b> and the <b>API key</b> to
          interact with the proxy.
        </p>
        <p>You can find information about how to use our gateway here : documentation.</p>
      </ProjectTokenCard>
    </ProjectTokenStyled>
  )
}

ProjectTokenView.propTypes = {
  token: PropTypes.string
}

ProjectTokenView.defaultProps = {
  token: undefined
}
