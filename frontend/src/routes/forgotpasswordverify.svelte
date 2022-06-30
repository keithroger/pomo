<script>
    import { isAuthenticated, contentLoading } from '$lib/store';
    import { Auth } from 'aws-amplify';

    let emailInput = "";
    let newpassword = "";
    let code = "";
    let loginError = false;

    async function passwordForgotVerify() {
        contentLoading.set(true);
        try {
            await Auth.forgotPasswordSubmit(emailInput, code, newpassword);
            isAuthenticated.set(true);
            window.location.href = "/login?msg=verified";
        } catch (error) {
            console.log('error signing in', error);
            loginError = true;
        }
        contentLoading.set(false);
    }

</script>

<div class="login-register-container">
    <h1>Verify</h1>
    {#if loginError}
        <div class="small-centered err">Error please try again.</div>
    {/if}
    <form on:submit|preventDefault={passwordForgotVerify}>
        <label for="email">Email</label>
        <input class="form-input" type="email" bind:value={emailInput} id="email" required/>
        <label for="code">Verification Code</label>
        <input class="form-input" type="text" bind:value={code} id="email" required/>
        <label for="password">New Password</label>
        <input class="form-input" type="password" bind:value={newpassword} id="password" required/>
        <button class="form-btn clickable rounded">Submit</button>
    </form>
</div>
