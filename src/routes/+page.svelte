<script lang="ts">
	import avatar from '$lib/assets/avatar.png';
</script>

<div class="py-10 sm:py-10">
	<div
		class="mx-auto max-w-2xl px-4 sm:px-6 lg:grid lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8"
	>
		<div class="lg:max-w-lg lg:self-end">
			<div class="mx-auto max-w-2xl lg:mx-0">
				<h2 class="text-4xl font-bold tracking-tight text-gray-900 dark:text-white sm:text-5xl">
					Dad, Runner, Engineer, Tinkerer
				</h2>

				<p class="mt-6 text-lg leading-8 text-gray-600 dark:text-white">
					I am Randell, a software engineer and tinkerer based in colorful Colorado. I enjoy
					building things, whether it be software, hardware, or something in between. I am a father
					of two (👶 and 🐶), a runner, and a lover of the outdoors.
				</p>
				<div class="mt-5 flex space-x-6 md:order-2">
					<a href="https://github.com/Piszmog" class="text-gray-400 hover:text-gray-500">
						<span class="sr-only">GitHub</span>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="icon icon-tabler icon-tabler-brand-github"
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
							<path
								d="M9 19c-4.3 1.4 -4.3 -2.5 -6 -3m12 5v-3.5c0 -1 .1 -1.4 -.5 -2c2.8 -.3 5.5 -1.4 5.5 -6a4.6 4.6 0 0 0 -1.3 -3.2a4.2 4.2 0 0 0 -.1 -3.2s-1.1 -.3 -3.5 1.3a12.3 12.3 0 0 0 -6.2 0c-2.4 -1.6 -3.5 -1.3 -3.5 -1.3a4.2 4.2 0 0 0 -.1 3.2a4.6 4.6 0 0 0 -1.3 3.2c0 4.6 2.7 5.7 5.5 6c-.6 .6 -.6 1.2 -.5 2v3.5"
							>
							</path>
						</svg>
					</a>
					<a
						href="https://www.linkedin.com/in/randell-callahan-083aa951"
						class="text-gray-400 hover:text-gray-500"
					>
						<span class="sr-only">LinkedIn</span>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="icon icon-tabler icon-tabler-brand-linkedin"
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
							<path
								d="M4 4m0 2a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2z"
							></path>
							<path d="M8 11l0 5"></path>
							<path d="M8 8l0 .01"></path>
							<path d="M12 16l0 -5"></path>
							<path d="M16 16v-3a2 2 0 0 0 -4 0"></path>
						</svg>
					</a>
				</div>
			</div>
		</div>
		<div class="mt-10 lg:col-start-2 lg:row-span-2 lg:mt-0 lg:self-center">
			<div class="aspect-h-1 aspect-w-1 overflow-hidden rounded-xl">
				<img src={avatar} alt="" class="h-full w-full object-cover object-center" />
			</div>
		</div>
	</div>
</div>
