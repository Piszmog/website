<script lang="ts">
	import type { Project } from '$lib/utils/types';
	import ProjectCard from '$lib/components/ProjectCard.svelte';

	export let projects: Project[];
</script>

<ul role="list" class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
	{#each projects as project (project.id)}
		<ProjectCard {...project} />
	{/each}
</ul>
