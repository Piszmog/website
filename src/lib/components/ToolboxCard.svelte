<script lang="ts">
	import { capitalize } from '$lib/utils/strings';

	export let title: string;
	export let items: string[];

	let colorClass = '';
	switch (title) {
		case 'plan':
			colorClass = 'bg-indigo-600';
			break;
		case 'code':
			colorClass = 'bg-green-600';
			break;
		case 'model':
			colorClass = 'bg-yellow-600';
			break;
		case 'view':
			colorClass = 'bg-red-600';
			break;
		case 'build':
			colorClass = 'bg-blue-600';
			break;
		case 'run':
			colorClass = 'bg-pink-600';
			break;
		case 'persist':
			colorClass = 'bg-purple-600';
			break;
		case 'move':
			colorClass = 'bg-gray-500';
			break;
	}
</script>

<li class="col-span-1 flex rounded-md shadow-sm">
	<div
		class="dark flex w-16 flex-shrink-0 items-center justify-center rounded-l-md text-sm font-medium text-white {colorClass}"
	>
		{capitalize(title)}
	</div>
	<div
		class="flex flex-1 items-center justify-between truncate rounded-r-md border-b border-r border-t border-gray-200 bg-white dark:border-slate-600 dark:bg-slate-800"
	>
		<div class="flex-1 truncate px-4 py-2 text-sm">
			{#each items as item, index (index)}
				<p class="text-gray-500 dark:text-slate-300">{item}</p>
			{/each}
		</div>
	</div>
</li>

<style lang="postcss">
</style>
