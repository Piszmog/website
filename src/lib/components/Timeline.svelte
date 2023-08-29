<script lang="ts">
	import type { Education, Job } from '$lib/utils/types';
	import TimelineEntryJob from '$lib/components/TimelineEntryJob.svelte';
	import TimelineEntryEducation from '$lib/components/TimelineEntryEducation.svelte';

	export let jobs: Job[] = [];
	export let education: Education[] = [];

	const data = [...jobs, ...education];

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
					details={d.details}
					changeReason={d.changeReason}
					toolbox={d.toolbox}
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
