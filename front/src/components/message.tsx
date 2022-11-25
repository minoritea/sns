import type { Message as MessageType } from "~/lib/message_client"
export default function Message({ message }: { message: MessageType }) {
  return <div>{ message.userName }: { message.body }</div>
}
