/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./**/*.{html,js,go}",
		"!./dist/**/*.{html,js,go}"
	],
	theme: {
		extend: {},
	},
	plugins: [],
}
