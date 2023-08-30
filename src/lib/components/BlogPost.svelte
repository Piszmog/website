<script lang="ts">
	export let title: string;
	export let date: string;
	export let length: number;
</script>

<div class="px-6 py-10 lg:px-8">
	<div class="mx-auto max-w-3xl text-base leading-7 text-gray-700 dark:text-slate-200">
		<h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 dark:text-slate-50 sm:text-4xl">
			{title}
		</h1>
		<div class="mt-1 flex items-center gap-x-2 text-xs leading-5 text-gray-500">
			<p class="text-xs leading-7">{new Date(date).toDateString()}</p>
			<svg viewBox="0 0 2 2" class="h-0.5 w-0.5 fill-current">
				<circle cx="1" cy="1" r="1" />
			</svg>
			<p class="text-xs leading-7">{length} mins</p>
		</div>
		<slot />
	</div>
</div>
