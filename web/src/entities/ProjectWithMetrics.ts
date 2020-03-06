export interface ProjectWithMetrics {
  title: string
  uuid: string
  metrics: {
    requestsCount: number
    requestsByDay: RequestByDay[]
  }
}

export interface RequestByDay {
  date: string
  value: number
}
