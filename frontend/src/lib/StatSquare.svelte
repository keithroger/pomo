<script>
    export let label;
    export let stat;
</script>
<div>
    <span id="label">{label}</span>
    <span id="stat">{stat}</span>
</div>

<style>
    div {
        border: 2px solid var(--primary);
        display: grid;
        border-radius: 8px;
    }
    
    span {
        text-align: center;
    }

    #label {
        color: var(--primary);
        justify-items: stretch;
        line-height: 2em;
    }

    #stat {
        color: var(--primary);
        line-height: 2em;
        height: 2.5em;
        /* font-family: "Roboto", sans-serif; */
        font-weight: bold;
        font-size: larger;
    }
</style>
