<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import postClient from '$lib/clients/post_client';
	import type { Post } from '$lib/clients/post_client';
	import Timeline from '$lib/components/timeline.svelte';

	let posts: Post[] = [];

	const aborter = new AbortController()

	async function load() {
		for await (const { post } of postClient.openSocialStream({}, { signal: aborter.signal })) {
			if (post != null) {
				posts = [post].concat(posts);
			}
		}
	}

	onMount(() => {
		load().catch(err => console.error(`streaming error: ${err}`))
	})

	onDestroy(() => {
		aborter.abort()
	})
</script>

<Timeline posts={posts} />

<style lang="postcss">
	ul {
		@apply mx-auto md:w-[600px] lg:w-[800px] bg-gray-300 rounded-lg flex flex-col gap-4 divide-y-2 divide-gray-50;
		min-height: 100svh;
	}

	li {
		@apply flex py-8 px-4;
	}

	li > img {
		@apply bg-white rounded-full object-contain w-16 h-16;
	}

	li > div {
		@apply flex flex-col ml-4;
	}

	li pre {
		@apply font-sans font-medium text-base;
	}

	.userName {
		@apply font-bold text-lg;
	}
</style>
