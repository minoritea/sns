<script lang="ts">
	import type { Post } from '$lib/clients/post_client';
	import socialClient from '$lib/clients/social_client';

	export let posts: Post[] = [];
	
	async function follow(followerName: string) {
		await socialClient.follow({ followerName })
	}
</script>

<ul>
	{#each posts as post}
		<li>
			<img src="/person.svg" alt="portlait" />
			<div>
				<div class="flex">
					<div class="userName">@{post.userName}</div>
					<button class="ml-4 rounded-lg p-1 bg-gray-400" on:click={() => follow(post.userName)}>Follow</button>
				</div>
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
