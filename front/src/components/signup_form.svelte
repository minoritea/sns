<script lang="ts">
	import { authentication } from "~/stores/session"
	import client from "~/clients/authentication_client"

	let name = ""
	let password = ""
	
	function signUp() {
    event.preventDefault()
		authentication(client.signUp(name, password))
	}

	$: submitDisabled = name === "" || password === ""
</script>

<form on:submit={signUp}>
	<label htmlFor="name">Name</label>
	<input name="name" type="text" bind:value={name} />
	<label htmlFor="password">Password</label>
	<input name="password" type="password" bind:value={password} />
	<input type="submit" disabled={submitDisabled} value="Sign up" />
</form>
