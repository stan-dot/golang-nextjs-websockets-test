

import React, { useState, useEffect } from 'react'

function WebsocketTester({ startData }) {
  const [ws, setWs] = useState(null)
  const [data, setData] = useState(startData)
  useEffect(() => {
    const newWs = new WebSocket('ws://localhost:8000/socket')
    newWs.onerror = err => console.error(err)
    newWs.onopen = () => setWs(newWs)
    newWs.onmessage = msg => setData(JSON.parse(msg.data))
    return () => {
      newWs.close()
    }
  }, [])

  return (
    <div>WebsocketTester

      <div>Data is: {JSON.stringify(data)}</div>
    </div>
  )
}

export default WebsocketTester