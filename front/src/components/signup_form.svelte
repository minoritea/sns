<script lang="ts">
	import { authentication } from "~/stores/session"
	import client from "~/clients/authentication_client"

	let name = ""
	let email = ""
	let password = ""
	
	function signUp(event: SubmitEvent) {
    event.preventDefault()
		authentication(client.signUp(name, email, password).then(() => {
			email = ""
			name = ""
			password =""
		}))
	}

	$: submitDisabled = email === "" || name === "" || password === ""
</script>

<div class="signup-box">
	<form on:submit={signUp} class="">
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
</div>

<style lang="postcss">
	.signup-box {
		@apply max-w-sm mx-auto mt-16 bg-gray-500 p-4 rounded-lg;
	}
	.signup-box form {
		@apply flex flex-col gap-4 justify-center;
	}
	.signup-box > form > input {
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
	.signup-box input[type=submit] {
		@apply rounded-lg w-1/3 text-center mt-2 mb-4 py-1 bg-cyan-200 text-gray-800 font-bold;
	}
</style>
