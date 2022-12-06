<script lang="ts">
	import { onMount } from "svelte"
	import TimeLine from "~/components/timeline.svelte"
	import AuthenticationForm from "~/components/authentication_form.svelte"
	import Sidebar from "~/components/sidebar.svelte"
	import client from "~/clients/authentication_client"
	import session, { authentication } from "~/stores/session"

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
