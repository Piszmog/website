<script lang="ts">
	import type { ChangeReason, Toolbox } from '$lib/utils/types';
	import JobIcon from '$lib/components/JobIcon.svelte';
	import { format } from '$lib/utils/date';
	import ToolboxCards from '$lib/components/ToolboxCards.svelte';

	export let start: string;
	export let end: string;
	export let company: string;
	export let title: string;
	export let details: string[];
	export let changeReason: ChangeReason;
	export let toolbox: Toolbox;
	export let isLast = false;
</script>

<div class="relative pb-8">
	{#if !isLast}
		<span
			class="absolute left-5 top-5 -ml-px h-full w-0.5 bg-gray-200 dark:bg-slate-400"
			aria-hidden="true"
		></span>
	{/if}
	<div class="relative flex items-start space-x-3">
		<div>
			<div class="relative px-1">
				<div
					class="flex h-8 w-8 items-center justify-center rounded-full bg-gray-100 ring-8 ring-white dark:bg-slate-600 dark:text-white dark:ring-slate-800"
				>
					<JobIcon {changeReason} />
				</div>
			</div>
		</div>
		<div class="min-w-0 flex-1">
			<div>
				<div class="text-sm">
					<p class="font-medium text-gray-900 dark:text-white">{company}</p>
				</div>
				<p class="mt-0.5 text-sm text-gray-500 dark:text-slate-300">{title}</p>
				<p class="mt-0.5 text-xs text-gray-500 dark:text-slate-300">
					{format(start)} - {format(end)}
				</p>
			</div>
			<ul class="list-none">
				{#each details as detail, index (index)}
					<li class="mt-2 flex items-center">
						<span class="mr-2 inline-block h-1 w-1 rounded-full bg-black"></span>
						<p class="text-sm text-gray-700 dark:text-slate-400">{detail}</p>
					</li>
				{/each}
			</ul>
			<ToolboxCards {toolbox} />
		</div>
	</div>
</div>
