import { createConnectTransport, createPromiseClient } from "@bufbuild/connect-web"
import { MessageService } from "~/proto/message_connectweb"
import type { Message } from "~/proto/message_pb"

export default createPromiseClient(MessageService, createConnectTransport({ baseUrl: "/rpc" }))
export type { Message }
