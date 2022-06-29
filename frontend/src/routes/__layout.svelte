<script>
    import Nav from "$lib/Nav.svelte";
    import Footer from "$lib/Footer.svelte";
    import Loader from "$lib/Loader.svelte";

	import { isAuthenticated, settings, contentLoading} from "$lib/store.js";
	import { userAPI } from "$lib/requests.js";
	import { setCSS } from "$lib/colors.js";

	import { onMount } from "svelte";
	import { Amplify, Auth } from "aws-amplify";

	// TODO change favicon and using it for loading icon
	
	let loading = true;

	onMount( async () => {
		// Configure authentificiation for aws cognito
		Amplify.configure({
			Auth: {
				region: "us-west-1",
				userPoolId: "us-west-1_yunzSfETg",
				userPoolWebClientId: "3hdjsr9rjseolsqstg4bq5q443",
			},
		});

		// Check if user is already logged in
		await Auth.currentSession()
			.then(() => {
				isAuthenticated.set(true);
			})
			.catch(() => isAuthenticated.set(false));

		// Load users settings
		if ($isAuthenticated) {
			await userAPI("get", "/settings", null)
			.then(resp => {
                settings.set(resp);
            })
            .catch(error => console.log(error));
		}

		setCSS($settings.theme);

		loading = false;
	});
</script>

{#if loading}
<Loader/>
{:else}
<div>
    <Nav/>
    {#if $contentLoading}
    <Loader/>
    {/if}
    <main>
        <slot></slot>
    </main>
    <Footer/>
</div>
{/if}

<style>
div {
    padding: 0;
    min-height: 100vh;
    max-width: 900px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

main {
    padding: 2em;
}

:global(body) {
    font-family: 'Lato', sans-serif;
    min-height: 100vh
}

:global(h1) {
    text-transform: uppercase;
    text-align: center;
    color: var(--text-on-bg);
}

:global(h2) {
    margin-top: 1.5em;
    color: var(--text-on-bg);
}

:global(h3) {
    margin-bottom: 0;
    color: var(--text-on-bg);
}

:global(p) {
    margin-top: 0.25em;
    color: var(--text-on-bg);
}

:global(body) {
    margin: 0;
    padding: 0;
    background-color: var(--bg);
}

:global(input[type=text], input[type=password], input[type=email]) {
    width: 100%;
    box-sizing: border-box;
    border-radius: 8px;
    border: 2px solid var(--primary);
    height: 2em;
    background-color: var(--bg);
    color: var(--text-on-bg);
}

:global(.login-register-container) {
    display: block;
    padding: 1em;
    max-width: 350px;
    margin: 0 auto;
}

:global(label) {
    float: left;
    margin-top: 1em;
    color: var(--text-on-bg);
}

:global(.form-btn) {
    width: 100%;
    margin-top: 2em;
    display: block;
}

:global(.small-centered) {
    display: grid;
    justify-items: center;
    font-size: small;
    margin: 2em auto;
}

:global(.link) {
    color: var(--primary);
    text-decoration: underline;
    cursor: pointer;
}

:global(.err) {
    color: var(--error);
}

:global(button) {
    border: none;
    outline: none;
    cursor: pointer;
}

:global(.rounded) {
    border-radius: 50px;
    overflow: hidden;
    box-sizing: border-box;
}

:global(.clickable) {
    background-color: var(--primary);
    color: var(--text-on-primary);
    height: 2.5em;
    font-size: medium;
    transition: all 0.1s;
}

:global(.clickable:hover) {
    background-color: var(--hover);
    color: var(--text-on-primary);
}

:global(.clickable:active) {
    background-color: var(--press);
    color: var(--text-on-primary);
}

:global(footer, nav) {
    background-color: var(--primary);
    color: var(--text-on-primary);
}
</style>