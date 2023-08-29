import type { PageLoad } from '../../../.svelte-kit/types/src/routes/experience/$types';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch('/experience.json');
	return await res.json();
};
