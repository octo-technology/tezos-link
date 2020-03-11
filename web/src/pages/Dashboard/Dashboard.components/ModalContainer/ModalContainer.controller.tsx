import React, { useEffect } from 'react'
import { createPortal } from 'react-dom'

type ModalProps = {
  children: React.ReactElement
}

export const ModalContainer = ({ children }: ModalProps) => {
  const modalRoot = document.getElementById('modal')
  const element = document.createElement('div')

  useEffect(() => {
    // @ts-ignore
    modalRoot.appendChild(element)

    return function cleanup() {
      // @ts-ignore
      modalRoot.removeChild(element)
    }
  })

  return createPortal(children, element)
}
