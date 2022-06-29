<script>
    import { userAPI } from "$lib/requests.js";
    import { isAuthenticated, contentLoading, settings } from "$lib/store.js";
    import { Auth } from "aws-amplify";

    let verificationCode = "";
    let resentTextVisible = false;
    let errorMessage = ""
    let errorVisible = false;
    let email = window.location.href.split("=")[1];
    console.log("page loaded");


    async function verify() {
        resentTextVisible = false;
        contentLoading.set(true);
        try {
            errorVisible = false;
            console.log("started confirmSignup");
            await Auth.confirmSignUp(email, verificationCode);
            console.log("passed confirmSignup");

            // Create User
            // TODO user should be created with current settings
            await userAPI("post", "", $settings);

            isAuthenticated.set(true);

        } catch (error) {
            errorMessage = "Error confirming sign up. Please try again.";
            errorVisible = true;
            console.log('error confirming sign up', error);
        }
        contentLoading.set(false);
    }

    async function resendConfirmationCode() {
        errorVisible = false;
        try {
            await Auth.resendSignUp(email);
            resentTextVisible = true;
            console.log('code resent successfully');
        } catch (err) {
            console.log('error resending code: ', err);
            errorMessage = "Error resending code.";
            errorVisible = true;
        }
    }
</script>

<div class="login-register-container">
    <h1>Verify Your Account</h1>
    {#if resentTextVisible}
        <div class="small-centered">Verification code resent.</div>
    {/if}

    {#if errorVisible}
        <div class="small-centered err">{errorMessage}</div>
    {/if}

    <form on:submit|preventDefault={verify}>
        <span>
            A code has been sent to your email.<br>
            Please enter your code below.
        </span>
        <input class="form-input" type="text" bind:value={verificationCode}>
        <button class="form-btn clickable rounded" on:click|preventDefault={verify}>Submit</button>
        <div class="small-centered">
            Didn't recieve the code?
            <span class="link" on:click|preventDefault={resendConfirmationCode}>Resend Code</span>
        </div>
    </form>
</div>