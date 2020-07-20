import * as React from 'react'
import { useEffect, useState } from 'react'
import axios, { AxiosResponse } from 'axios'
import { StatusView } from './Status.view'

export const Status = () => {
  const [proxyStatus, setProxyStatus] = useState(true)
  const [nodeMainnetArchiveStatus, setMainnetArchiveNodeStatus] = useState(true)
  const [nodeCarthagenetArchiveStatus, setCarthagenetArchiveNodeStatus] = useState(true)
  const [nodeMainnetRollingStatus, setMainnetRollingNodeStatus] = useState(true)
  const [nodeCarthagenetRollingStatus, setCarthagenetRollingNodeStatus] = useState(true)
  const [date, setDate] = useState('')

  useEffect(() => {
    setDate(new Date(Date.now()).toLocaleString())
    const mainnetProxyUrl = 'https://mainnet.tezoslink.io'
    const carthagenetProxyUrl = 'https://carthagenet.tezoslink.io'

    axios.get(mainnetProxyUrl + '/health').catch((error: any) => {
      console.error(error)
      setProxyStatus(false)
    })

    axios({
      method: 'get',
      url: mainnetProxyUrl + '/status'
    }).then((response : AxiosResponse) => {
      let jdata = JSON.parse(response.data)
      setMainnetArchiveNodeStatus(jdata["data"]["archive_node"])
      setMainnetRollingNodeStatus(jdata["data"]["rolling_node"])
    }).catch((error: any) => {
      console.error(error)
      setMainnetArchiveNodeStatus(false)
      setMainnetRollingNodeStatus(false)
    })

    axios({
      method: 'get',
      url: carthagenetProxyUrl + '/status'
    }).then((response : AxiosResponse) => {
      let jdata = JSON.parse(response.data)
      setCarthagenetArchiveNodeStatus(jdata["data"]["archive_node"])
      setCarthagenetRollingNodeStatus(jdata["data"]["rolling_node"])
    }).catch((error: any) => {
      console.error(error)
      setCarthagenetArchiveNodeStatus(false)
      setCarthagenetRollingNodeStatus(false)
    })
  })

  return (
    <>
      <StatusView proxyStatus={proxyStatus}
      nodeMainnetArchiveStatus={nodeMainnetArchiveStatus}
      nodeMainnetRollingStatus={nodeMainnetRollingStatus}
      nodeCarthagenetArchiveStatus={nodeCarthagenetArchiveStatus}
      nodeCarthagenetRollingStatus={nodeCarthagenetRollingStatus}
      date={date} />
    </>
  )
}
