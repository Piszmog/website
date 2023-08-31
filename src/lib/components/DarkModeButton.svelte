<script lang="ts">
	let isDarkMode = false;
	const toggleDarkMode = () => {
		isDarkMode = !isDarkMode;
		document.body.classList.toggle('dark', isDarkMode);
	};
</script>

<button
	on:click={toggleDarkMode}
	type="button"
	class="rounded-full bg-white px-2.5 py-2.5 text-sm font-semibold ring-1 ring-inset ring-gray-300 hover:bg-gray-50 dark:bg-slate-800 dark:text-slate-400 dark:ring-slate-400"
	aria-label="Toggle dark mode"
>
	{#if isDarkMode}
		<svg
			xmlns="http://www.w3.org/2000/svg"
			class="icon icon-tabler icon-tabler-moon"
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
			<path d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z">
			</path>
		</svg>
	{:else}
		<svg
			xmlns="http://www.w3.org/2000/svg"
			class="icon icon-tabler icon-tabler-sun-high"
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
			<path d="M14.828 14.828a4 4 0 1 0 -5.656 -5.656a4 4 0 0 0 5.656 5.656z"></path>
			<path d="M6.343 17.657l-1.414 1.414"></path>
			<path d="M6.343 6.343l-1.414 -1.414"></path>
			<path d="M17.657 6.343l1.414 -1.414"></path>
			<path d="M17.657 17.657l1.414 1.414"></path>
			<path d="M4 12h-2"></path>
			<path d="M12 4v-2"></path>
			<path d="M20 12h2"></path>
			<path d="M12 20v2"></path>
		</svg>
	{/if}
</button>
