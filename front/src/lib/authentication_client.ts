import { createConnectTransport, createPromiseClient, PromiseClient } from "@bufbuild/connect-web"
import { AuthenticationService } from "~/proto/authentication_connectweb"
import { SignUpRequest } from "~/proto/authentication_pb"

export class AuthenticationClient {
  client: PromiseClient<typeof AuthenticationService>;
  constructor(serverUrl: string) {
    const transport = createConnectTransport({
      baseUrl: serverUrl
    })
    this.client = createPromiseClient(AuthenticationService, transport)
  }

  async isSignedIn(): Promise<void> {
    await this.client.isSignedIn(null)
  }

  async signUp(name: string, password: string): Promise<void> {
    await this.client.signUp(new SignUpRequest({name, password}))
  }
}

export default new AuthenticationClient("/rpc")
