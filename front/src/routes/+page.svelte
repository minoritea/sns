<script lang="ts" context="module">
	export const prerender = true;
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import LocalTimeline from '$lib/components/local_timeline.svelte';
	import SocialTimeline from '$lib/components/social_timeline.svelte';
	import AuthenticationForm from '$lib/components/authentication_form.svelte';
	import Sidebar from '$lib/components/sidebar.svelte';
	import client from '$lib/clients/authentication_client';
	import session, { authentication } from '$lib/stores/session';

	onMount(() => {
		authentication(client.isSignedIn({})).catch(console.error);
	});
</script>

{#if $session !== 'unauthenticated'}
	<div class="flex justify-center mt-12 gap-4">
		<div>
			<Sidebar />
		</div>
		<div>
		{#if $page.url.searchParams.get('tl') === 'social'}
			<SocialTimeline />
		{:else}
			<LocalTimeline />
		{/if}
		</div>
		<div />
	</div>
{:else}
	<AuthenticationForm />
{/if}
