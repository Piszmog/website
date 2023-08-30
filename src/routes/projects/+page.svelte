<script lang="ts">
	import type { PageData } from './$types';
	import ProjectCards from '$lib/components/ProjectCards.svelte';

	export let data: PageData;
</script>

<div class="py-10 sm:py-10">
	<div class="flex justify-center pb-10">
		<div class="mx-auto max-w-2xl lg:mx-0">
			<h2 class="text-3xl font-bold tracking-tight text-gray-900 dark:text-white sm:text-4xl">
				My Projects 🚀
			</h2>
			<p class="mt-2 text-lg leading-8 text-gray-600 dark:text-white">
				Some of my open source projects that I have created or contributed to.
			</p>
		</div>
	</div>
	<div class="flex justify-center border-t border-gray-200 py-10">
		<ProjectCards projects={data.projects} />
	</div>
</div>
