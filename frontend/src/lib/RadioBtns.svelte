<script>
    export let values;
    export let name;
    export let selected;
    export let nPerRow;

</script>

<div class="wrapper" style="grid-template-columns: repeat({nPerRow}, 1fr);">
    {#each values as value}
        <label for={value}>
            <input type="radio" {name} id={value} {value} bind:group={selected}>
            <span class="rounded">{value}</span>
        </label>
    {/each}
</div>

<style>
    .wrapper {
        display: grid;
        gap: 1em;
        width: 100%;
        align-items: stretch;
    }

    label{
        display: block;
        float: none;
        margin: 0;
    }

    input[type=radio] {
        display: none;
    }

    input[type=radio]:checked ~ span{
        background-color: var(--primary);
        color: var(--text-on-primary);
    }

    span {
        display: table;
        border: 2px solid var(--primary);
        line-height: 2.5em;
        text-align: center;
        width: 100%;
        cursor: pointer;
        color: var(--primary);
    }

    span:hover {
        background-color: var(--hover);
    }

    span:active {
        background-color: var(--press);
        color: var(--text-on-primary)
    }
</style>
