<script>
    import {
    isAuthenticated,
    } from "./store.js";
    import { Auth } from "aws-amplify";

    async function logout() {
        try {
            await Auth.signOut();
            isAuthenticated.set(false);
            window.location.href = "/";
        } catch (error) {
            console.log('error signing out: ', error);
        }
    }
</script>

<nav class="nav-container">
    <span class="logo clickable"><a href="/">POMO.CAFE</a></span>
    
    <ul>
        {#if $isAuthenticated}
        <li><a href="/stats" class="clickable">STATS</a></li>
        <li><a href="/settings" class="clickable">SETTINGS</a></li>
        <li><span class="clickable" on:click={logout}>LOG OUT</span></li>
        {:else}
        <li><a href="/settings" class="clickable">SETTINGS</a></li>
        <li><a href="/login" class="clickable">LOGIN</a></li>
        <li><a href="/register" class="clickable">REGISTER</a></li>
        {/if}
    </ul>
</nav>

<style>
    .nav-container {
        display: flex;
        justify-content: space-between;
    }

    .logo {
        font-size: 2em;
        height: 40px;
        line-height: 40px;
        font-weight: bold;
    }

    a {
        text-decoration: none;
        color: var(--text-on-primary);
    }

    ul {
        list-style-type: none;
        margin: 0;
        padding: 0 1em;
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        gap: 1em;
    }

    span {
        height: 40px;
        line-height: 40px;
        padding: 10px 10px;
        cursor: pointer;
    }

</style>