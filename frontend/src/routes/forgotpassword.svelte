<script>
    import { isAuthenticated, contentLoading } from '$lib/store';
    import { Auth } from 'aws-amplify';

    let emailInput = "";
    let loginError = false;

    async function forgotPassword() {
        contentLoading.set(true);
        try {
            await Auth.forgotPassword(emailInput)
            window.location.href = "/forgotpasswordverify";
        } catch {
            loginError = true;
        }
        contentLoading.set(false);
    }

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
    <h1>Forgot Password</h1>
    <p>Please enter your email and you recieve a confirmation code.</p>
    {#if loginError}
        <div class="small-centered err">Incorrect username or password.</div>
    {/if}
    <form on:submit|preventDefault={forgotPassword}>
        <label for="email">Email</label>
        <input class="form-input" type="email" bind:value={emailInput} id="email" required/>
        <button class="form-btn clickable rounded">Submit</button>
    </form>
</div>
