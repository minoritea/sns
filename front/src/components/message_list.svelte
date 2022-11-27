<script lang="ts">
	import client, { Message } from "~/lib/message_client.ts"
	import { onMount } from "svelte"
	import { writable } from "svelte/store"

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
