

import TextEditor from './TextEditor'
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
      // newWs.close()
      // todo this was causing issues
    }
  }, [])

  return (
    <div>

      <h2>
        WebsocketTester
      </h2>

      <div>Data is: {JSON.stringify(data)}</div>
      {ws && <TextEditor ws={ws} startText={data} />}
    </div>
  )
}

export default WebsocketTester