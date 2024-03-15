<script lang="ts">
	import { Greet } from "$lib/wailsjs/go/main/App.js";

	let greeting: any;
	$: console.log(greeting);
</script>

<section>
	<h1 class="text-red-500">Wails test</h1>
	<button on:click={() => Greet("This is a greeting").then(res => greeting = res)}>Greet</button>
	{#if greeting !== undefined}
		<p>{greeting.Message}</p>
	{/if}
</section>
