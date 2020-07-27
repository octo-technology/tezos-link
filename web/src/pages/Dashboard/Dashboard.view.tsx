import * as PropTypes from 'prop-types'
import * as React from 'react'
import { Button } from '../../App/App.components/Button/Button.controller'

import { ProjectTokenView } from './Dashboard.components/ProjectToken/ProjectToken.view'
import {
  DashboardHeader,
  DashboardLeftSide,
  DashboardRightSide,
  DashboardBottomStyled,
  DashboardTitle,
  DashboardTopStyled,
  DashboardStyled
} from './Dashboard.style'
import { ProjectWithMetrics } from '../../entities/ProjectWithMetrics'
import { ProjectNameView } from './Dashboard.components/ProjectName/ProjectName.view'
import { RequestsByDayLineView } from './Dashboard.components/RequestsByDayLine/RequestsByDayLine.view'
import { RPCUsagePieView } from './Dashboard.components/RPCUsagePie/RPCUsagePie.view'
import { LastRequestsView } from './Dashboard.components/LastRequests/LastRequests.view'

type DashboardViewProps = { project: ProjectWithMetrics; loading: boolean }

export const DashboardView = ({ loading, project }: DashboardViewProps) => (
  <DashboardStyled>
    <DashboardTopStyled>
      <DashboardLeftSide>
        <DashboardHeader>
          <DashboardTitle>
            <h1>Dashboard</h1>
          </DashboardTitle>
          <Button text={project.network} color="secondary" />
        </DashboardHeader>
        {loading ? (
          'Loading...'
        ) : (
          <>
            <RequestsByDayLineView requestsByDay={project.metrics.requestsByDay} />
          </>
        )}
      </DashboardLeftSide>
      <DashboardRightSide>
        <ProjectNameView name={project.title} />
        <ProjectTokenView token={project.uuid} />
      </DashboardRightSide>
    </DashboardTopStyled>
    <DashboardBottomStyled>
      <RPCUsagePieView rpcUsage={project.metrics.rpcUsage} rpcTotalCount={project.metrics.requestsCount} />
      <LastRequestsView lastRequests={project.metrics.lastRequests} />
    </DashboardBottomStyled>
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
