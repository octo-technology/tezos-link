import * as PropTypes from 'prop-types'
import * as React from 'react'
import { downColor, upColor } from 'src/styles'

// prettier-ignore
import { ToasterClose, ToasterContent, ToasterCountdown, ToasterGrid, ToasterIcon, ToasterMessage, ToasterStyled, ToasterTitle } from './Toaster.style'

type ToasterViewProps = { status?: string; title: string; message: string }

export const ToasterView = ({ status, title, message }: ToasterViewProps) => {
  const backgroundColor = status === 'success' ? upColor : downColor

  return (
    <ToasterStyled className={status ? 'showing' : 'hidden'}>
      <ToasterGrid>
        <ToasterIcon style={{ backgroundColor }}>
          {status === 'success' ? (
            <img alt={status} src={`/icons/success.svg`} />
          ) : (
            <img alt={status} src={`/icons/error.svg`} />
          )}
        </ToasterIcon>
        <ToasterContent>
          <ToasterTitle>{title}</ToasterTitle>
          <ToasterMessage>{message}</ToasterMessage>
        </ToasterContent>
        <ToasterClose>
          <img alt={status} src={`/icons/close.svg`} />
        </ToasterClose>
      </ToasterGrid>
      <ToasterCountdown className={status ? 'showing' : 'hidden'} style={{ backgroundColor }} />
    </ToasterStyled>
  )
}

ToasterView.propTypes = {
  status: PropTypes.string,
  title: PropTypes.string,
  message: PropTypes.string
}

ToasterView.defaultProps = {
  status: undefined,
  title: 'Error',
  message: 'Undefined error'
}
