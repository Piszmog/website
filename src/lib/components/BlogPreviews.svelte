<script lang="ts">
	import type { Blog } from '$lib/utils/types';
	import BlogPreview from '$lib/components/BlogPreview.svelte';

	export let blogs: Blog[];
</script>

{#each blogs as blog (blog.id)}
	<BlogPreview {...blog} />
{/each}
