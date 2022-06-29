<script>
    import { contentLoading } from '$lib/store.js';
    import { Auth } from 'aws-amplify';

    // TODO add google login option
    // TODO add AWS SES so that emails aren't limited for registration

    let emailInput = "";
    let passwordInput = "";
    let confirmPasswordInput = "";
    let validationChecks = new Array();

    function validate() {
        validationChecks = new Array();
        // Must at least 1 lowercase letter
        if (!/[a-z]+/.test(passwordInput)) {
            validationChecks.push("Password must contain at least 1 lowercase character.");
        }
        // Must at least 1 uppercase letter
        if (!/[A-Z]+/.test(passwordInput)) {
            validationChecks.push("Password must contain at least 1 uppercase character.");
        }
        // Must at least 1 numeric character
        if (!/[0-9]+/.test(passwordInput)) {
            validationChecks.push("Password must contain at least 1 numeric character.");
        }
        // Must contain at least 1 special character (^$*.[]{}()?-"!@#%&/\,><':;|_~`+=)
        if (!/[\^$*.\[\]{}\(\)?\-"!@#%&\/,><\’:;|_~`]+/.test(passwordInput)) {
            validationChecks.push("Password must contain at least special character.");
        }
        // Must contain at least 8 characters long 
        if (!/[\^$*.\[\]{}\(\)?\-"!@#%&\/,><\’:;|_~`A-Za-z0-9]{8}/.test(passwordInput)) {
            validationChecks.push("Password must contain at least 8 characters.");
        }

        // Check if passwords match
        if (passwordInput !=  confirmPasswordInput) {
            validationChecks.push("Passwords must match");
        }

    }

    async function signUp() {
        validate();
        if (validationChecks.length > 0) {
            return;
        }

        // signup with cognito
        contentLoading.set(true);
        try {
            await Auth.signUp({
                username: emailInput,
                password: passwordInput,
                attributes: {
                    email: emailInput,
                }
            });
        } catch (error) {
            console.log('error signing up:', error);
        }

        window.location.href = "/verify?user=" + emailInput;
        contentLoading.set(false);
    }

</script>

<div class="login-register-container">
    <h1>Register</h1>
    {#if validationChecks.length > 0}
        <div class="small-centered err">
            {#each validationChecks as v}
        <span>{v}</span><br>
            {/each}
        </div>
    {/if}

    <form on:submit|preventDefault={signUp}>
        <label for="email">Email</label>
        <input class="form-input" type="email" bind:value={emailInput} id="email" required/>
        <label for="password">Password</label>
        <input class="form-input" type="password" bind:value={passwordInput} id="password" required/>
        <label for="confirmpassword">Confirm Password</label>
        <input class="form-input" type="password" bind:value={confirmPasswordInput} id="confirmpassword" required/>
        <button class="form-btn clickable rounded">Sign Up</button>
        <div class="small-centered"><span>Already have an account? <a href="/login" class="link">Login</a></span></div>
    </form>
</div>