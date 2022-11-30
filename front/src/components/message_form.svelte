<script lang="ts">
	import client, { Message } from "~/lib/message_client"

	let body = ""

	function submit() {
    event.preventDefault()
    client.post(new Message({ body }))
		body = ""
	}

	$: submitDisabled = body === ""
</script>

<form on:submit={submit}>
	<input type="text" bind:value={body} />
	<input type="submit" disabled={submitDisabled} value="Post" />
</form>
