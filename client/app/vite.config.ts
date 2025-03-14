import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';
import Icons from 'unplugin-icons/vite'


export default defineConfig({
	plugins: [
		sveltekit(),
		Icons({
			compiler: 'svelte',
			autoInstall: true,
    })
	],

	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
