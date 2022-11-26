import { useCallback, useRef, FormEvent } from "react"
import client from "~/lib/authentication_client"
import { authentication } from "~/lib/session_state"

export default function SignUpForm() {
  const nameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)
  const signUp = useCallback((event: FormEvent) => {
    event.preventDefault()
    const name = nameRef.current?.value
    if (name == null || name === "") {
      throw new Error("name is empty")
    }

    const password = passwordRef.current?.value
    if (password == null || password === "") {
      throw new Error("password is empty")
    }

    authentication(client.signUp(name, password))
  }, [])

  return (
    <form onSubmit={signUp}>
      <label htmlFor="name">Name</label>
      <input name="name" ref={nameRef} type="text" />
      <label htmlFor="password">Password</label>
      <input name="password" ref={passwordRef} type="password" />
      <input type="submit" disabled={nameRef.current && nameRef.current?.value === "" && passwordRef.current && passwordRef.current?.value === ""} value="Sign up" />
    </form>
  )
}
