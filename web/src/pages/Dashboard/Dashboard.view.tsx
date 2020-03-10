import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Button } from '../../App/App.components/Button/Button.controller'

import { ProjectTokenView } from './Dashboard.components/ProjectToken/ProjectToken.view'
import {
  DashboardHeader,
  DashboardLeftSide,
  DashboardRightSide,
  DashboardStyled,
  DashboardTitle
} from './Dashboard.style'
import { ProjectWithMetrics } from '../../entities/ProjectWithMetrics'
import { ProjectNameView } from './Dashboard.components/ProjectName/ProjectName.view'
import { RequestsByDayLineView } from './Dashboard.components/RequestsByDayLine/RequestsByDayLine.view'
import { RPCUsagePieView } from './Dashboard.components/RPCUsagePie/RPCUsagePie.view'

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
          <RequestsByDayLineView requestsByDay={project.metrics.requestsByDay} />
          <RPCUsagePieView rpcUsage={project.metrics.rpcUsage} rpcTotalCount={project.metrics.requestsCount} />
        </>
      )}
    </DashboardLeftSide>
    <DashboardRightSide>
      <ProjectNameView name={project.title} />
      <ProjectTokenView token={project.uuid} />
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
