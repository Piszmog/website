// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import type { Blog, Education, Job, Project } from '$lib/utils/types';

declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		interface PageData {
			jobs: Job[];
			education: Education[];
			projects: Project[];
			blogs: Blog[];
		}
		// interface Platform {}
	}
}

export {};
