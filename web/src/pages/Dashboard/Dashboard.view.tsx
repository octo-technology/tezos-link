import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Button } from '../../App/App.components/Button/Button.controller'

import { Categories } from './Dashboard.components/Categories/Categories.controller'
import { ProjectToken } from './Dashboard.components/ProjectToken/ProjectToken.controller'
import {
  DashboardHeader,
  DashboardLeftSide,
  DashboardRightSide,
  DashboardStyled,
  DashboardTitle
} from './Dashboard.style'
import { ProjectWithMetrics } from '../../entities/ProjectWithMetrics'
import { RequestsCounterView } from './Dashboard.components/RequestsCounter/RequestsCounter.view'
import { NoRequestInfoView } from './Dashboard.components/NoRequestInfo/NoRequestInfo.view'
import { ProjectNameView } from './Dashboard.components/ProjectName/ProjectName.view'

type DashboardViewProps = { project: ProjectWithMetrics; loading: boolean }

export const DashboardView = ({ loading, project }: DashboardViewProps) => (
  <DashboardStyled>
    <DashboardLeftSide>
      <DashboardHeader>
        <DashboardTitle>
          <h1>Dashboard</h1>
        </DashboardTitle>
        <Button text="Mainnet" color="secondary" />
      </DashboardHeader>
      {loading ? (
        'Loading...'
      ) : (
        <>
          {project.metrics.requestsCount === 0 ? <NoRequestInfoView /> : <></>}
          <RequestsCounterView count={project.metrics.requestsCount} />
        </>
      )}
    </DashboardLeftSide>
    <DashboardRightSide>
      <ProjectNameView name={project.title} />
      <ProjectToken token={project.uuid} />
      <Categories />
    </DashboardRightSide>
  </DashboardStyled>
)

DashboardView.propTypes = {
  project: PropTypes.any,
  loading: PropTypes.bool
}

DashboardView.defaultProps = {
  metrics: {},
  loading: true
}
