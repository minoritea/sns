export default function Message({ message }: { message: { body: string } }) {
  return <div>{ message.body }</div>
}
