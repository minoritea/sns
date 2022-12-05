<script lang="ts" context="module">
  export const prerender = true
</script>

<script lang="ts">
	import { authentication } from "$lib/stores/session"
	import client from "$lib/clients/authentication_client"

	let name = ""
	let email = ""
	let password = ""
	
	function signUp(event: SubmitEvent) {
    event.preventDefault()
		authentication(client.signUp({name, email, password})).then(() => {
			window.location.href = "/"
		}).catch(console.error)
	}

	$: submitDisabled = email === "" || name === "" || password === ""
</script>

<div class="authenticatoin-box">
	<form on:submit={signUp}>
		<label for="name">Name</label>
		<div class="name-box">
			<div class="name-box-inline">
				<span>@</span>
				<input name="name" type="text" bind:value={name} />
			</div>
		</div>
		<label for="email">Email</label>
		<input name="email" type="email" bind:value={email} />
		<label for="password">Password</label>
		<input name="password" type="password" bind:value={password} />
		<div class="w-full flex justify-center">
			<input type="submit" disabled={submitDisabled} value="Sign up" />
		</div>
	</form>
	<div class="flex flex-row-reverse">
		<a href="/signin" class="underline">Sign in</a>
	</div>
</div>

<style lang="postcss">
	.authenticatoin-box {
		@apply max-w-sm mx-auto mt-16 bg-gray-500 p-4 rounded-lg;
	}
	.authenticatoin-box form {
		@apply flex flex-col gap-4 justify-center;
	}
	.authenticatoin-box > form > input {
		@apply rounded w-full bg-gray-300 focus:ring-blue-500;
	}
	.name-box-inline {
		@apply flex w-full;
	}
	.name-box-inline {
		@apply inline-flex;
	}
	.name-box-inline > span {
		@apply px-2 rounded-l bg-gray-600;
	}
	.name-box-inline > input {
		@apply rounded-r w-full bg-gray-300 focus:ring-blue-500;
	}
	.authenticatoin-box input[type=submit] {
		@apply rounded-lg w-1/3 text-center mt-2 mb-4 py-1 bg-cyan-200 text-gray-800 font-bold;
	}
</style>
