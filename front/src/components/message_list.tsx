import { useState, useEffect } from "react"
import client, { Message as MessageType } from "~/lib/client"
import Message from "~/components/message"

export default function MessageList() {
  const [messages, setMessages] = useState<MessageType[] | null>(null)
  useEffect(() => {
    client.fetchList().then(setMessages)
  }, [])

  if (messages == null) {
    return <div>loading...</div>
  }

  return <ul>
    { messages.map((message, index) => <li key={index}><Message message={message} /></li>) }
  </ul>;
}
