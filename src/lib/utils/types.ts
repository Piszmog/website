export interface Job {
	id: number;
	start: string;
	end: string;
	company: string;
	title: string;
	description: string;
	changeReason: ChangeReason;
	skills: string[];
}

export enum ChangeReason {
	New = 'new',
	Promotion = 'promotion'
}

export interface Education {
	id: number;
	start: string;
	end: string;
	school: string;
	degree: string;
	field: string;
}
