// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import type { Education, Job, Project } from '$lib/utils/types';

declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		interface PageData {
			jobs: Job[];
			education: Education[];
			projects: Project[];
		}
		// interface Platform {}
	}
}

export {};
