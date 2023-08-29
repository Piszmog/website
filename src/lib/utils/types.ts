export interface Job {
	id: number;
	start: string;
	end: string;
	company: string;
	title: string;
	details: string[];
	changeReason: ChangeReason;
	toolbox: Toolbox;
}

export enum ChangeReason {
	New = 'new',
	Promotion = 'promotion'
}

export interface Toolbox {
	plan: string[];
	code: string[];
	model: string[];
	view: string[];
	build: string[];
	run: string[];
	persist: string[];
	move: string[];
}

export interface Education {
	id: number;
	start: string;
	end: string;
	school: string;
	degree: string;
	field: string;
}
