export interface ProjectWithMetrics {
  title: string
  uuid: string
  metrics: {
    requestsCount: number
    requestsByDay: RequestByDay[]
    rpcUsage: RPCUsage[]
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
