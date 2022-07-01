import adapter from '@sveltejs/adapter-static';
// import { index } from 'd3';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter: adapter({fallback: 'index.html'}),
		adapter: adapter(),
		prerender: {
			default: true
		},
		vite: {
            define: {
                global: {} // workaround for dev issue
            },
			resolve: { // workaround for build issue
				alias: {
				'./runtimeConfig': './runtimeConfig.browser',
				},
			},
        }
	}
};

export default config;
