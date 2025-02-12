import containerQueries from '@tailwindcss/container-queries';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				'basic': '#171717',
				'standard': '#242424',
			},
			fontSize: {
				'xxs': '0.6rem',
			}
		}
	},

	plugins: [typography, forms, containerQueries]
} satisfies Config;
