<script lang="ts" context="module">
  export const prerender = true
</script>

<script lang="ts">
	import { onMount } from "svelte"
	import TimeLine from "$lib/components/timeline.svelte"
	import AuthenticationForm from "$lib/components/authentication_form.svelte"
	import Sidebar from "$lib/components/sidebar.svelte"
	import client from "$lib/clients/authentication_client"
	import session, { authentication } from "$lib/stores/session"

	onMount(() => {
		authentication(client.isSignedIn({})).catch(console.error)
	})
</script>

{#if $session !== "unauthenticated"}
	<div class="flex justify-center mt-12 gap-4">
		<div>
			<Sidebar />
		</div>
		<div>
			<TimeLine />
		</div>
		<div></div>
	</div>
{:else}
	<AuthenticationForm />
{/if}
