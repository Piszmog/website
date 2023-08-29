<script lang="ts">
	import { format } from '$lib/utils/date';

	export let start: string;
	export let end: string;
	export let school: string;
	export let degree: string;
	export let field: string;
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
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="icon icon-tabler icon-tabler-school"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						stroke-width="2"
						stroke="currentColor"
						fill="none"
						stroke-linecap="round"
						stroke-linejoin="round"
					>
						<path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
						<path d="M22 9l-10 -4l-10 4l10 4l10 -4v6"></path>
						<path d="M6 10.6v5.4a6 3 0 0 0 12 0v-5.4"></path>
					</svg>
				</div>
			</div>
		</div>
		<div class="min-w-0 flex-1">
			<div>
				<div class="text-sm">
					<p class="font-medium text-gray-900">{school}</p>
				</div>
				<p class="mt-0.5 text-sm text-gray-500">{degree}</p>
				<p class="mt-0.5 text-xs text-gray-500">{format(start)} - {format(end)}</p>
			</div>
			<div class="mt-2 text-sm text-gray-700">
				<p>
					{field}
				</p>
			</div>
		</div>
	</div>
</div>
