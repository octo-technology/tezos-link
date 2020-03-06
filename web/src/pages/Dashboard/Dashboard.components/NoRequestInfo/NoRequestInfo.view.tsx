import * as React from 'react'
import { NoRequestInfoCard, NoRequestInfoStyled } from './NoRequestInfo.style'
import { Link } from 'react-router-dom'

export const NoRequestInfoView = () => {
  return (
    <NoRequestInfoStyled>
      <NoRequestInfoCard>
        <img height="30" alt="warning" src="/icons/warning.svg" />
        <p>
          It looks like you haven't done any request for now. Please have a look to the{' '}
          <Link to="http://google.fr">documentation</Link>
        </p>
      </NoRequestInfoCard>
    </NoRequestInfoStyled>
  )
}
