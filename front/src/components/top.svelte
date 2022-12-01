<script lang="ts">
	import { onMount } from "svelte"
	import MessageList from "~/components/message_list.svelte"
	import MessageForm from "~/components/message_form.svelte"
	import AuthenticationForm from "~/components/authentication_form.svelte"
	import client from "~/clients/authentication_client"
	import session, { authentication } from "~/stores/session"

	onMount(() => {
		authentication(client.isSignedIn({}))
	})
</script>

{#if $session === "authenticated"}
	<MessageForm />
	<MessageList />
{:else if $session === "unauthenticated"}
	<AuthenticationForm />
{:else}
	<div class="text-white">loading...</div>
{/if}
