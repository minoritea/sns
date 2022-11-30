<script lang="ts">
	import client, { Message } from "~/clients/message_client.ts"
	import { onMount } from "svelte"
	import { writable } from "svelte/stores"

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
