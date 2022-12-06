<script lang="ts">
	import client, { Message } from "~/clients/message_client"
	import { onMount } from "svelte"

	let messages: Message[] = []

	onMount(() => {
		(async function load(){
			for await (const { message } of client.openStream()) {
				messages = [message].concat(messages)
			}
		})().catch(console.error)
	})
</script>

<ul>
{ #each messages as message }
	<li>{ message.body }</li>
{ /each }
</ul>
