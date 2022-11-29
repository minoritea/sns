<script lang="ts">
	import { browser } from '$app/environment';
	import MessageList from '$lib/components/message_list.svelte';
	import MessageForm from '$lib/components/message_form.svelte';
	import SignUpForm from '$lib/components/signup_form.svelte';
	import client from '$lib/rpc/authentication_client';
	import session, { authentication } from '$lib/store/session';

	if (browser) {
		authentication(client.isSignedIn());
	}
</script>

{#if $session == 'authenticated'}
	<MessageForm />
	<MessageList />
{:else if $session == 'unauthenticated'}
	<SignUpForm />
{:else}
	<div>loading...</div>
{/if}
