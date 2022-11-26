import { useEffect } from "react"
import MessageList from "~/components/message_list"
import MessageForm from "~/components/message_form"
import SignUpForm from "~/components/signup_form"
import { useSessionState, authentication } from "~/lib/session_state"
import client from "~/lib/authentication_client"

export default function Top() {
  const status = useSessionState()
  useEffect(() => authentication(client.isSignedIn()), [])
  return (<>
    { status === "loading" ? <div></div> : status === "unauthenticated" ? <SignUpForm /> : <MessageForm /> }
    <MessageList />
  </>);
}
