<script lang="ts">
	import PostForm from "~/components/post_form.svelte"

	let visible = false
</script>

<aside class="h-full w-64 bg-gray-300 rounded-lg pt-4 grid grid-cols-1 justify-items-center content-start">
	<button class="mt-4 w-52 py-2 rounded-lg bg-gray-400 " on:click={() => { visible = true }}>Post</button>
</aside>

<PostForm bind:visible />

<style>
</style>
