const monthAbbreviations = [
	'Jan',
	'Feb',
	'Mar',
	'Apr',
	'May',
	'Jun',
	'Jul',
	'Aug',
	'Sep',
	'Oct',
	'Nov',
	'Dec'
];

export const format = (input: string): string => {
	const date = new Date(input);
	const month = date.getMonth();
	const year = date.getFullYear();
	return `${monthAbbreviations[month]} ${year}`;
};
