<script>
    import { isAuthenticated, contentLoading } from '$lib/store';
    import { Auth } from 'aws-amplify';

    let emailInput = "";
    let password = "";
    let loginError = false;
    let msg = window.location.href.split("=")[1];

    async function signIn() {
        contentLoading.set(true);
        try {
            await Auth.signIn(emailInput, password);
            isAuthenticated.set(true);
            window.location.href = "/";
        } catch (error) {
            console.log('error signing in', error);
            loginError = true;
        }
        contentLoading.set(false);
    }

</script>

<div class="login-register-container">
    <h1>Login</h1>
    {#if loginError}
        <div class="small-centered err">Incorrect username or password.</div>
    {/if}
    {#if msg == "verified"}
        <div class="small-centered">Account has been verified. Please sign in.</div>
    {/if}
    <form on:submit|preventDefault={signIn}>
        <label for="email">Email</label>
        <input class="form-input" type="email" bind:value={emailInput} id="email" required/>
        <label for="password">Password</label>
        <input class="form-input" type="password" bind:value={password} id="password" required/>
        <a href="/forgotpassword" class="link">Forgot Password?</a>
        <button class="form-btn clickable rounded">Submit</button>
        <div class="small-centered"><span>Don't have an account? <a href="/register" class="link">Sign Up</a></span></div>
    </form>
</div>
