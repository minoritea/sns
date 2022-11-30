<script lang="ts">
	import { browser } from '$app/environment';
	import client, { Message } from '$lib/rpc/message_client.ts';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let messages: Message[] = [];

	if (browser) {
		async function load() {
			for await (const { message } of client.openStream()) {
				messages = [message].concat(messages);
			}
		}

		load().catch(console.error);
	}
</script>

<ul>
	{#each messages as message}
		<li class="message">
			<img src="/person.svg" />
			<div class="box">
				<div class="userName">
					@{message.userName}
				</div>
				<div class="body">
					{message.body}
				</div>
			</div>
		</li>
	{/each}
</ul>

<style>
	li.message {
		margin: 1rem 0;
		display: flex;
		align-items: center;
	}

	li.message .box {
		flex: auto;
		display: flex;
		flex-direction: column;
	}

	li.message > img {
		object-fit: cover;
		width: 5rem;
		height: 5rem;
	}

	li.message div.userName {
		margin: 0 0 0 1rem;
		font-size: 125%;
	}

	li.message div.body {
		margin: 1rem 0 0 1rem;
	}
</style>
