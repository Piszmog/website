<script lang="ts">
	import LanguageBadge from '$lib/components/LanguageBadge.svelte';

	export let name: string;
	export let description: string;
	export let url: string;
	export let languages: string[];
</script>

<li class="divide-y divide-gray-200 rounded-lg bg-white shadow dark:bg-slate-700">
	<div class="flex w-full items-center justify-between p-6">
		<div class="flex-1 truncate">
			<div class="flex items-center space-x-3">
				<h3 class="truncate text-sm font-medium text-gray-900 dark:text-white">{name}</h3>
			</div>
			<p class="mt-1 truncate text-sm text-gray-500 dark:text-slate-200">{description}</p>
			<div class="mt-3 space-x-1">
				{#each languages as language, index (index)}
					<LanguageBadge {language} />
				{/each}
			</div>
		</div>
	</div>
	<div>
		<div class="-mt-px flex divide-x divide-gray-200">
			<div class="flex w-0 flex-1">
				<a
					href={url}
					target="_blank"
					class="relative -mr-px inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-bl-lg border border-transparent py-4 text-sm font-semibold text-gray-900 dark:text-slate-50"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="icon icon-tabler icon-tabler-world"
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
						<path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0 -18 0"></path>
						<path d="M3.6 9h16.8"></path>
						<path d="M3.6 15h16.8"></path>
						<path d="M11.5 3a17 17 0 0 0 0 18"></path>
						<path d="M12.5 3a17 17 0 0 1 0 18"></path>
					</svg>
					URL
				</a>
			</div>
		</div>
	</div>
</li>
