import { createConnectTransport, createPromiseClient, PromiseClient } from "@bufbuild/connect-web"
import { MessageStream }from "~/proto/message_connectweb"
import { Message, Response }from "~/proto/message_pb"
import { Observable, from } from "rxjs"

export class Client {
  client: PromiseClient<typeof MessageStream>;
  constructor(serverUrl: string) {
    const transport = createConnectTransport({
      baseUrl: serverUrl
    })
    this.client = createPromiseClient(MessageStream, transport)
  }

  open(): Observable<Response> {
    return from(this.client.open(null))
  }
}

export default new Client("/rpc")
export { Message, Response }
