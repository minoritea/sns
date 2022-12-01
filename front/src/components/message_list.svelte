<script lang="ts">
	import client, { Message } from "~/clients/message_client"
	import session from "~/stores/session"

	let messages: Message[] = []

	async function load(){
		for await (const { message } of client.openStream({})) {
			messages = [message].concat(messages)
		}
	}

	$: {
	  if ($session === "authenticated") {
			load().then(console.error)
		}
	}
</script>

<ul>
{ #each messages as message }
	<li>
		<img src="/person.svg" alt="portlait" />
		<div>
			<div class="userName">@{ message.userName }</div>
			<pre class="body">{ message.body }</pre>
		</div>
	</li>
{ /each }
</ul>

<style lang="postcss">
	ul {
		@apply mx-auto md:w-[600px] lg:w-[800px] bg-gray-300 rounded-lg flex flex-col gap-4 divide-y-2 border-gray-200;
		min-height: 100svh;
	}

	li {
		@apply flex py-8 px-4;
	}

	img {
		@apply bg-white rounded-full;
	}

	li > div {
		@apply flex flex-col ml-4;
	}

	.userName {
		@apply font-bold text-lg;
	}
</style>
