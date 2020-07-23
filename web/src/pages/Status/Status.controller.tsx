import * as React from 'react'
import { useEffect, useState } from 'react'
import axios, { AxiosResponse } from 'axios'
import { StatusView } from './Status.view'

export const Status = () => {
  const [proxyMainnetStatus, setMainnetProxyStatus] = useState(true)
  const [proxyCarthagenetStatus, setCarthagenetProxyStatus] = useState(true)
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
      setMainnetProxyStatus(false)
    })

    axios.get(carthagenetProxyUrl + '/health').catch((error: any) => {
      console.error(error)
      setCarthagenetProxyStatus(false)
    })

    axios({
      method: 'get',
      url: mainnetProxyUrl + '/status'
    }).then((response : AxiosResponse) => {
      let jdata = response.data.data
      setMainnetArchiveNodeStatus(jdata["archive_node"])
      setMainnetRollingNodeStatus(jdata["rolling_node"])
    }).catch((error: any) => {
      console.error(error)
      setMainnetArchiveNodeStatus(false)
      setMainnetRollingNodeStatus(false)
    })

    axios({
      method: 'get',
      url: carthagenetProxyUrl + '/status'
    }).then((response : AxiosResponse) => {
      let jdata = response.data.data
      setCarthagenetArchiveNodeStatus(jdata["archive_node"])
      setCarthagenetRollingNodeStatus(jdata["rolling_node"])
    }).catch((error: any) => {
      console.error(error)
      setCarthagenetArchiveNodeStatus(false)
      setCarthagenetRollingNodeStatus(false)
    })
  })

  return (
    <>
      <StatusView
      proxyMainnetStatus={proxyMainnetStatus}
      proxyCarthagenetStatus={proxyCarthagenetStatus}
      nodeMainnetArchiveStatus={nodeMainnetArchiveStatus}
      nodeMainnetRollingStatus={nodeMainnetRollingStatus}
      nodeCarthagenetArchiveStatus={nodeCarthagenetArchiveStatus}
      nodeCarthagenetRollingStatus={nodeCarthagenetRollingStatus}
      date={date} />
    </>
  )
}
