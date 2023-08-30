<script lang="ts">
	import type { Blog } from '$lib/utils/types';
	import BlogPreview from '$lib/components/BlogPreview.svelte';

	export let blogs: Blog[];

	const data = blogs.sort((a, b) => {
		return new Date(b.date).getTime() - new Date(a.date).getTime();
	});
</script>

{#each data as blog (blog.id)}
	<BlogPreview {...blog} />
{/each}
