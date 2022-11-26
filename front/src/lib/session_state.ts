import { BehaviorSubject, distinctUntilChanged, share } from "rxjs"
import { useState, useEffect } from "react"

export type Status = "loading" | "authenticated" | "unauthenticated"
const subject = new BehaviorSubject<Status>("loading") 
const sessionState$ = subject.pipe(distinctUntilChanged(), share({ connector: () => new BehaviorSubject<Status>("loading"), resetOnError: false, resetOnComplete: false, resetOnRefCountZero: false }))

export function useSessionState(): Status {
  const [status, setStatus] = useState<Status>("loading")
  useEffect(() => {
    const subscription = sessionState$.subscribe(setStatus)
    return () => subscription.unsubscribe()
  }, [])
  return status
}

export function setSessionState(status: Status) { subject.next(status) }
export function authentication(promise: Promise<any>) {
  promise.then(() => setSessionState("authenticated"), err => {
    console.error(err)
    setSessionState("unauthenticated")
  })
}
