import { createConnectTransport, createPromiseClient, PromiseClient } from "@bufbuild/connect-web"
import { MessageService }from "~/proto/message_connectweb"
import { Message }from "~/proto/message_pb"

export class Client {
  client: PromiseClient<typeof MessageService>;
  constructor(serverUrl: string) {
    const transport = createConnectTransport({
      baseUrl: serverUrl
    })
    this.client = createPromiseClient(MessageService, transport)
  }

  async fetchList(): Promise<Message[]> {
    const { messages }= await this.client.list(null)
    return messages
  }
}

export default new Client("/rpc")
export type Message = Message
