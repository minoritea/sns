import { useState, useEffect } from "react"
import client, { Message as MessageType } from "~/lib/message_client"
import Message from "~/components/message"
import { share, scan } from "rxjs"
import sessionState from "~/lib/session_state"

function useMessageStream() {
  const [messages, setMessages] = useState<MessageType[] | null>(null)
  useEffect(() => {
    const message$ = client.openStream().pipe(share())
    message$.pipe(scan((acc, res) => [res.message].concat(acc), [] as MessageType[])).subscribe(setMessages)
    message$.subscribe({ next: () => sessionState.next(true), error: () => sessionState.next(false) })
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
