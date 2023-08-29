<script lang="ts">
	import type { Toolbox } from '$lib/utils/types';
	import ToolboxCard from '$lib/components/ToolboxCard.svelte';

	export let toolbox: Toolbox;

	const entries = Object.entries(toolbox).map(([key, value]) => ({ key, value }));
</script>

<div class="mt-2">
	<h2 class="text-md font-medium text-gray-500 dark:text-slate-300">Toolbox</h2>
	<ul role="list" class="mt-3 grid grid-cols-1 gap-5 sm:grid-cols-2 sm:gap-6 lg:grid-cols-4">
		{#each entries as entry (entry.key)}
			{#if entry.value.length > 0}
				<ToolboxCard title={entry.key} items={entry.value} />
			{/if}
		{/each}
	</ul>
</div>
