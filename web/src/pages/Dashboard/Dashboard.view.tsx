import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Button } from 'src/App/App.components/Button/Button.controller'

import { Categories } from './Dashboard.components/Categories/Categories.controller'
import { ProjectToken } from './Dashboard.components/ProjectToken/ProjectToken.controller'
import { postsMock } from './Dashboard.mock'
//prettier-ignore
import { DashboardHeader, DashboardLeftSide, DashboardRightSide, DashboardStyled, DashboardTitle } from './Dashboard.style'

export const DashboardView = () => (
  <DashboardStyled>
    <DashboardLeftSide>
      <DashboardHeader>
        <DashboardTitle>
          <h1>Dashboard</h1>
        </DashboardTitle>
        <Button text="Active" color="secondary"/>
      </DashboardHeader>
      Some charts here...
    </DashboardLeftSide>
    <DashboardRightSide>
      <ProjectToken token="EFJKFE7837REHFEBH"/>
      <Categories />
    </DashboardRightSide>
  </DashboardStyled>
)

DashboardView.propTypes = {
  posts: PropTypes.array,
  loading: PropTypes.bool,
  fetchMoreCallback: PropTypes.func.isRequired
}

DashboardView.defaultProps = {
  posts: postsMock,
  loading: true
}
