<script lang="ts">
	import client from '$lib/clients/post_client';

	export let visible = false;

	let body = '';

	function submit() {
		client.publish({ body });
		body = '';
		visible = false;
	}

	$: submitDisabled = body === '';
</script>

{#if visible}
	<div
		class="post-form fixed w-full h-full bg-gray-500 inset-0 bg-opacity-50 backdrop-blur-sm flex justify-center"
		on:click={() => {
			visible = false;
		}}
		on:keypress={() => {
			visible = false;
		}}
	>
		<div
			class="mt-12"
			on:click|stopPropagation={() => {
				/* eslint-disable-line no-empty-function */
			}}
			on:keypress|stopPropagation={() => {
				/* eslint-disable-line no-empty-function */
			}}
		>
			<form on:submit|preventDefault={submit}>
				<textarea maxlength="1000" minlength="1" rows="4" bind:value={body} />
				<input type="submit" disabled={submitDisabled} value="Post" />
			</form>
		</div>
	</div>
{/if}

<style lang="postcss">
	.post-form form {
		@apply bg-gray-300 p-8 flex flex-col gap-4 justify-center rounded-lg w-96;
	}

	.post-form textarea {
		@apply rounded-lg bg-gray-200;
	}

	.post-form input[type='submit'] {
		@apply rounded-lg w-1/3 bg-gray-400 py-1 mx-auto;
	}
</style>
