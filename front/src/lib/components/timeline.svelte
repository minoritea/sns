<script lang="ts">
	import client from '$lib/clients/post_client';
	import type { Post } from '$lib/clients/post_client';
	import session from '$lib/stores/session';

	let posts: Post[] = [];

	async function load() {
		for await (const { post } of client.openStream({})) {
			if (post != null) {
				posts = [post].concat(posts);
			}
		}
	}

	$: {
		if ($session === 'authenticated') {
			load().then(console.error);
		}
	}
</script>

<ul>
	{#each posts as post}
		<li>
			<img src="/person.svg" alt="portlait" />
			<div>
				<div class="userName">@{post.userName}</div>
				<pre class="body">{post.body}</pre>
			</div>
		</li>
	{/each}
</ul>

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
