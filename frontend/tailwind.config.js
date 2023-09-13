import colors from 'tailwindcss/colors';

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				textPrimary: colors.slate[100],
				textSecondary: colors.slate[300],
				textTertiary: colors.slate[500],

				surfacePrimary: colors.slate[900],
				surfaceSecondary: colors.slate[800],
				surfaceTertiary: colors.slate[700],

				primary: colors.indigo[800],
				primaryLight: colors.indigo[700],
				primaryDark: colors.indigo[900]
			}
		}
	},
	plugins: []
};
