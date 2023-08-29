<script lang="ts">
	import type { Education, Job } from '$lib/utils/types';
	import TimelineEntryJob from '$lib/components/TimelineEntryJob.svelte';
	import TimelineEntryEducation from '$lib/components/TimelineEntryEducation.svelte';

	export let jobs: Job[] = [];
	export let educations: Education[] = [];

	const data = [...jobs, ...educations];

	data.sort(
		(a: Job | Education, b: Job | Education) =>
			new Date(b.start).getTime() - new Date(a.start).getTime()
	);
</script>

<div class="flow-root">
	<ul role="list" class="-mb-8">
		{#each data as d, index (d.id)}
			{#if d.company !== undefined}
				<TimelineEntryJob
					start={d.start}
					end={d.end}
					company={d.company}
					title={d.title}
					description={d.description}
					changeReason={d.changeReason}
					skills={d.skills}
					isLast={index === data.length - 1}
				/>
			{:else if d.school !== undefined}
				<TimelineEntryEducation
					start={d.start}
					end={d.end}
					school={d.school}
					degree={d.degree}
					field={d.field}
					isLast={index === data.length - 1}
				/>
			{/if}
		{/each}
	</ul>
</div>
