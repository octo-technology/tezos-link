import * as React from 'react'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { StatusView } from './Status.view'

export const Status = () => {
  const [proxyStatus, setProxyStatus] = useState(true)
  const [nodeStatus, setNodeStatus] = useState(true)
  const [date, setDate] = useState('')

  useEffect(() => {
    setDate(new Date(Date.now()).toLocaleString())

    axios.get('http://mainnet.tezoslink.io/health').catch((error: any) => {
      console.error(error)
      setProxyStatus(false)
    })

    axios({
      method: 'get',
      url: 'http://mainnet.tezoslink.io/v1/ece6cee2-09a3-4092-8cb2-a1c41602db86/chains/main/blocks/head'
    }).catch((error: any) => {
      console.error(error)
      setNodeStatus(false)
    })
  })

  return (
    <>
      <StatusView proxyStatus={proxyStatus} nodeStatus={nodeStatus} date={date} />
    </>
  )
}
