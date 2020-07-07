export interface ProjectWithMetrics {
  title: string
  uuid: string
  network : string
  metrics: {
    requestsCount: number
    requestsByDay: RequestByDay[]
    rpcUsage: RPCUsage[]
    lastRequests: string[]
  }
}

export interface RequestByDay {
  date: string
  value: number
}

export interface RPCUsage {
  id: string
  label: string
  value: number
}
