import { createConnectTransport, createPromiseClient, PromiseClient } from "@bufbuild/connect-web"
import { MessageService }from "~/proto/message_connectweb"
import { Message, Response }from "~/proto/message_pb"
import { Observable, from } from "rxjs"

export class MessageClient {
  client: PromiseClient<typeof MessageService>;
  constructor(serverUrl: string) {
    const transport = createConnectTransport({
      baseUrl: serverUrl
    })
    this.client = createPromiseClient(MessageService, transport)
  }

  openStream(): Observable<Response> {
    return from(this.client.openStream(null))
  }

  async post(message: Message): Promise<void> {
    await this.client.post(message)
  }
}

export default new MessageClient("/rpc")
export { Message, Response }
