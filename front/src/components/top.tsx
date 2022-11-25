import { useState, useEffect } from "react"
import MessageList from "~/components/message_list"
import MessageForm from "~/components/message_form"
import SignUpForm from "~/components/signup_form"
import sessionState from "~/lib/session_state"
import { distinctUntilChanged } from "rxjs"

export default function Top() {
  const [isSignedIn, setSignedIn] = useState<boolean | null>(null)

  useEffect(() => {
    const subscription = sessionState.pipe(distinctUntilChanged()).subscribe(setSignedIn)
    return () => subscription.unsubscribe()
  }, [])

  return (<>
    { isSignedIn == null ? <div></div> : !isSignedIn ? <SignUpForm /> : <MessageForm /> }
    <MessageList />
  </>);
}
