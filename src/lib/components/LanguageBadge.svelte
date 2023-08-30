<script lang="ts">
	export let language: string;
</script>

<span
	class:go={language.toLowerCase() === 'go'}
	class:rust={language.toLowerCase() === 'rust'}
	class:javascript={language.toLowerCase() === 'javascript'}
	class:java={language.toLowerCase() === 'java'}
	class:spring-boot={language.toLowerCase() === 'spring boot'}
	class:typescript={language.toLowerCase() === 'typescript'}
	class:svelte={language.toLowerCase() === 'svelte'}
	class="inline-flex items-center rounded-full px-2 py-1 text-xs font-medium ring-1 ring-inset"
>
	{language}
</span>

<style lang="postcss">
	.go {
		@apply bg-blue-50 text-blue-700 ring-blue-600/20;
	}

	.rust {
		@apply bg-red-50 text-red-700 ring-red-600/20;
	}

	.javascript {
		@apply bg-yellow-50 text-yellow-700 ring-yellow-600/20;
	}

	.java {
		@apply bg-purple-50 text-purple-700 ring-purple-600/20;
	}

	.spring-boot {
		@apply bg-green-50 text-green-700 ring-green-600/20;
	}

	.typescript {
		@apply bg-indigo-50 text-indigo-700 ring-indigo-600/20;
	}

	.svelte {
		@apply bg-gray-50 text-gray-700 ring-gray-600/20;
	}
</style>
