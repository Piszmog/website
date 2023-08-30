<script lang="ts">
	export let title: string;
</script>

<div class="px-6 py-10 lg:px-8">
	<div class="mx-auto max-w-3xl text-base leading-7 text-gray-700 dark:text-slate-200">
		<h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 dark:text-slate-50 sm:text-4xl">
			{title}
		</h1>
		<slot />
	</div>
</div>
