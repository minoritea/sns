<script lang="ts">
	import client from "~/clients/message_client"
	
	export let visible = false

	let body = ""

	function submit() {
    event.preventDefault()
    client.post({ body })
		body = ""
		visible = false
	}

	$: submitDisabled = body === ""
</script>

{#if visible}
	<div
	  class="fixed w-full h-full bg-gray-500 inset-0 bg-opacity-50 backdrop-blur-sm flex justify-center"
		on:click={() => { visible = false }}
		on:keypress={() => { visible = false }}
	>
		<div
		  class="mt-12"
			on:click|stopPropagation={() => {}}
			on:keypress|stopPropagation={() => {}}
		>
			<form on:submit={submit}>
				<textarea maxlength=1000 minlength=1 rows=4 bind:value={body} />
				<input type="submit" disabled={submitDisabled} value="Post" />
			</form>
		</div>
	</div>
{/if}

<style lang="postcss">
form {
	@apply bg-gray-300 p-8 flex flex-col gap-4 justify-center rounded-lg w-96;
}

textarea {
	@apply rounded-lg bg-gray-200;
}

input[type=submit] {
	@apply rounded-lg w-1/3 bg-gray-400 py-1 mx-auto;
}
</style>
