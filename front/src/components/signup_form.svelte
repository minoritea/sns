<script lang="ts">
	import { authentication } from "~/stores/session"
	import client from "~/clients/authentication_client"

	let name = ""
	let email = ""
	let password = ""
	
	function signUp(event: SubmitEvent) {
    event.preventDefault()
		authentication(client.signUp(name, email, password))
	}

	$: submitDisabled = name === "" || password === ""
</script>

<form on:submit={signUp}>
	<label for="name">Name</label>
	<input name="name" type="text" bind:value={name} />
	<label for="email">Email</label>
	<input name="email" type="email" bind:value={email} />
	<label for="password">Password</label>
	<input name="password" type="password" bind:value={password} />
	<input type="submit" disabled={submitDisabled} value="Sign up" />
</form>
