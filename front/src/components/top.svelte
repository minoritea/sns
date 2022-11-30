<script lang="ts">
	import { onMount } from "svelte"
	import MessageList from "~/components/message_list.svelte"
	import MessageForm from "~/components/message_form.svelte"
	import SignUpForm from "~/components/signup_form.svelte"
	import client from "~/clients/authentication_client"
	import session, { authentication } from "~/stores/session"

	onMount(() => {
		authentication(client.isSignedIn())
	})

	$: isSignedIn = $session === "authenticated"
</script>

{#if isSignedIn}
	<MessageForm />
	<MessageList />
{:else}
	<SignUpForm />
{/if}
