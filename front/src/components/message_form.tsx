import client, { Message } from "~/lib/message_client"
import { useCallback, useRef, FormEvent } from "react"

export default function MessageForm() {
  const ref = useRef<HTMLInputElement>(null)
  const post = useCallback((event: FormEvent) => {
    event.preventDefault()

    const body = ref.current?.value
    if (body == null || body === "") {
      throw new Error("body is empty")
    }

    client.post(new Message({ body }))
    ref.current.value = ""
  }, [])
  return (
    <form onSubmit={post}>
      <input ref={ref} type="text" />
      <input type="submit" disabled={ref.current && ref.current?.value === ""} value="Post" />
    </form>
  );
}
