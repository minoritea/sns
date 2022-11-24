import { useState, useEffect } from "react"
import MessageList from "~/components/message_list"
import MessageForm from "~/components/message_form"
import SignUpForm from "~/components/signup_form"
import client from "~/lib/authentication_client"
import sessionState from "~/lib/session_state"

export default function Top() {
  const [isSignedIn, setSignedIn] = useState(false)
  useEffect(() => {
    const subscription = sessionState.subscribe(setSignedIn)
    client.isSignedIn().then(() => sessionState.next(true)).catch(() => sessionState.next(false))
    return () => subscription.unsubscribe()
  }, [])
  
  if (!isSignedIn) {
    return <SignUpForm />
  }
  return (<>
    <MessageForm />
    <MessageList />
  </>);
}
