import { useState, useEffect } from "react"
import client, { Message as MessageType } from "~/lib/message_client"
import Message from "~/components/message"
import { scan } from "rxjs"

function useMessageStream() {
  const [messages, setMessages] = useState<MessageType[] | null>(null)
  useEffect(() => {
    client.openStream().pipe(scan((acc, res) => [res.message].concat(acc), [] as MessageType[])).subscribe(setMessages)
  }, [])
  return messages
}

export default function MessageList() {
  const messages = useMessageStream()

  if (messages == null) {
    return <div>loading...</div>
  }

  return <ul>
    { messages.map((message, index) => <li key={index}><Message message={message} /></li>) }
  </ul>;
}
