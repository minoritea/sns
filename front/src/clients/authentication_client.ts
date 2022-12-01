import { createConnectTransport, createPromiseClient } from "@bufbuild/connect-web"
import { AuthenticationService } from "~/proto/authentication_connectweb"

export default createPromiseClient(AuthenticationService, createConnectTransport({ baseUrl: "/rpc" }))
