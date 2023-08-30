<script lang="ts">
	import TagBadge from '$lib/components/TagBadge.svelte';

	export let title: string;
	export let date: string;
	export let tags: string[];
	export let preview: string;
	export let slug: string;
</script>

<article class="mt-5 flex max-w-xl flex-col items-start justify-between">
	<div class="flex items-center gap-x-4 text-xs">
		<p class="text-gray-500">{new Date(date).toDateString()}</p>
		{#each tags as tag, index (index)}
			<TagBadge {tag} />
		{/each}
	</div>
	<div class="group relative">
		<h3 class="mt-3 text-lg font-semibold leading-6 text-gray-900 group-hover:text-gray-600">
			<a href={slug}>
				<span class="absolute inset-0"></span>
				{title}
			</a>
		</h3>
		<p class="mt-5 line-clamp-3 text-sm leading-6 text-gray-600">
			{preview}
		</p>
	</div>
</article>
