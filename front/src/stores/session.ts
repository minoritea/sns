import { writable } from "svelte/store"

export type Status = "loading" | "authenticated" | "unauthenticated"
const SessionStore = writable<Status>("loading")
export default SessionStore

export async function authentication(promise: Promise<any>) {
  return promise.then(
    () => SessionStore.set("authenticated"),
    err => {
      SessionStore.set("unauthenticated")
      throw err;
    },
  )
}
