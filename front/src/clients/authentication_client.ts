import { createConnectTransport, createPromiseClient } from "@bufbuild/connect-web"
import type { PromiseClient } from "@bufbuild/connect-web"
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

  async signUp(email: string, name: string, password: string): Promise<void> {
    await this.client.signUp(new SignUpRequest({email, name, password}))
  }

  async isSignedIn(): Promise<void> {
    await this.client.isSignedIn(null)
  }
}

export default new AuthenticationClient("/rpc")
