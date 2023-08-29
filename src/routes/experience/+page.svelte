<script lang="ts">
	import Timeline from '$lib/components/Timeline.svelte';
	import { ChangeReason } from '$lib/utils/types';
</script>

<div class="bg-white py-10 sm:py-10">
	<div class="flex justify-center py-10">
		<div class="mx-auto max-w-2xl lg:mx-0">
			<h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">My Experience</h2>
			<p class="mt-2 text-lg leading-8 text-gray-600">Places I have worked and what I learned.</p>
		</div>
	</div>
	<div class="flex justify-center border-t border-gray-200 py-10">
		<Timeline
			jobs={[
				{
					id: 0,
					start: '2021-01-01',
					end: '2022-01-01',
					company: 'Company B',
					title: 'Title B',
					description: 'Fo bar baz',
					changeReason: ChangeReason.New,
					skills: ['Java', 'MongoDB']
				},
				{
					id: 1,
					start: '2020-01-01',
					end: '2021-01-01',
					company: 'Company A',
					title: 'Title A',
					description: 'Fo bar baz',
					changeReason: ChangeReason.Promotion,
					skills: ['Java', 'MongoDB']
				}
			]}
		/>
	</div>
</div>
