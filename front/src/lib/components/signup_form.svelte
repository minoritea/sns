<script lang="ts">
	import { authentication } from '$lib/store/session';
	import client from '$lib/rpc/authentication_client';

	let name = '';
	let password = '';

	function signUp() {
		event.preventDefault();
		authentication(client.signUp(name, password));
	}

	$: submitDisabled = name === '' || password === '';
</script>

<form on:submit={signUp}>
	<label for="name">Name</label>
	<input name="name" type="text" bind:value={name} />
	<label for="password">Password</label>
	<input name="password" type="password" bind:value={password} />
	<input type="submit" disabled={submitDisabled} value="Sign up" />
</form>
