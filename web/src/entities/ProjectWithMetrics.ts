export interface ProjectWithMetrics {
  title: string
  uuid: string
  metrics: {
    requestsCount: number
  }
}
