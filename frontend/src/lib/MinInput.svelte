<script>
    export let value;
</script>

<div>
    <input id="pomo-time" type="text" bind:value={value}> <label for="pomo-time">Minutes</label>
</div>

<style>

    div {
        display: grid;
        grid-template-columns: 1fr 3fr;
        gap: 1em;
        align-items: end;
        padding: 0;
    }

    div label {
        margin-top: 0;
    }

    input[type=text] {
        width: 100%;
    }
</style>