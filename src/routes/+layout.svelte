<script lang="ts">
	import '../app.css';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { dev } from '$app/environment';
	import { inject } from '@vercel/analytics';

	inject({ mode: dev ? 'development' : 'production' });
</script>

<Header />
<div class="flex min-h-screen flex-col">
	<main class="flex-grow dark:bg-slate-800">
		<div class="py-10">
			<div class="mx-auto sm:px-6 lg:px-8">
				<slot />
			</div>
		</div>
	</main>

	<Footer />
</div>
