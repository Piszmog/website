<script lang="ts">
	import Timeline from '$lib/components/Timeline.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="py-10 sm:py-10">
	<div class="flex justify-center pb-10">
		<div class="mx-auto max-w-2xl lg:mx-0">
			<h2 class="text-3xl font-bold tracking-tight text-gray-900 dark:text-white sm:text-4xl">
				My Experiences 💪
			</h2>
			<p class="mt-2 text-lg leading-8 text-gray-600 dark:text-white">
				Places I have worked and what I learned.
			</p>
		</div>
	</div>
	<div class="flex justify-center border-t border-gray-200 py-10">
		<Timeline jobs={data.jobs} education={data.education} />
	</div>
</div>
