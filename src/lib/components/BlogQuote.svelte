<script lang="ts">
	export let author: string;
	export let title = '';
</script>

<figure class="mt-10 border-l border-blue-600 pl-9">
	<blockquote class="font-semibold text-gray-900 dark:text-slate-50">
		<p>
			“<slot />“
		</p>
	</blockquote>
	<figcaption class="mt-6 flex gap-x-4">
		<div class="text-sm leading-6">
			<strong class="font-semibold text-gray-900 dark:text-slate-50">{author}</strong>
			{#if title}
				- {title}
			{/if}
		</div>
	</figcaption>
</figure>
