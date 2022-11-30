import { writable } from "svelte/store"

export type Status = "loading" | "authenticated" | "unauthenticated"
const SessionStore = writable<Status>("loading")
export default SessionStore

export function authentication(promise: Promise<any>) {
  promise.then(() => SessionStore.set("authenticated"), err => {
    console.error(err)
    SessionStore.set("unauthenticated")
  })
}
