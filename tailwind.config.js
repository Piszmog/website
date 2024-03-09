/** @type {import('tailwindcss').Config} */
module.exports = {
	darkMode: 'class',
	content: [
		"./components/**/*.templ",
	],
	theme: {
		extend: {},
	},
	plugins: [
		require('@tailwindcss/forms'),
		require('@tailwindcss/typography'),
	],
}

