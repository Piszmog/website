<script lang="ts">
	import type { ChangeReason } from '$lib/utils/types';
	import JobIcon from '$lib/components/JobIcon.svelte';
	import { format } from '$lib/utils/date';

	export let start: string;
	export let end: string;
	export let company: string;
	export let title: string;
	export let description: string;
	export let changeReason: ChangeReason;
	export let isLast = false;
</script>

<div class="relative pb-8">
	{#if !isLast}
		<span class="absolute left-5 top-5 -ml-px h-full w-0.5 bg-gray-200" aria-hidden="true"></span>
	{/if}
	<div class="relative flex items-start space-x-3">
		<div>
			<div class="relative px-1">
				<div
					class="flex h-8 w-8 items-center justify-center rounded-full bg-gray-100 ring-8 ring-white"
				>
					<JobIcon {changeReason} />
				</div>
			</div>
		</div>
		<div class="min-w-0 flex-1">
			<div>
				<div class="text-sm">
					<p class="font-medium text-gray-900">{company}</p>
				</div>
				<p class="mt-0.5 text-sm text-gray-500">{title}</p>
				<p class="mt-0.5 text-xs text-gray-500">{format(start)} - {format(end)}</p>
			</div>
			<div class="mt-2 text-sm text-gray-700">
				<p>
					{description}
				</p>
			</div>
		</div>
	</div>
</div>
