<script>
    import { userAPI } from "$lib/requests.js"
    import StatSquare from "$lib/StatSquare.svelte";
    import RadioBtns from "$lib/RadioBtns.svelte";
    import Visualization from "$lib/Visualization.svelte";
    import { pomodoros, contentLoading } from "$lib/store.js";
    import { getGraphData } from "$lib/statsprocessing.js";

    // TODO if no data, then display "No data to display."
    // TODO make data update sections

    // Retrieve data from database.
    async function getData() {
        contentLoading.set(true);

        const resp = await userAPI("get", "/stats", null);
        pomodoros.set(resp);

        const data = getGraphData($pomodoros);

        contentLoading.set(false);

        return data;
    }

</script>

<div>
    <h1>Stats</h1>
    <p>
        Keep track of how many pomodoros you have completed. Every completed
        pomodoro counts towards your stats.
    </p>

    {#await getData() then graphData}

    {#if $pomodoros.length > 0}

    <h3>Period of Time</h3>
    <RadioBtns
        values={["Week", "Month", "Year"]}
        name="period"
        selected="all"
        nPerRow=3
    />

    <!-- bar graph -->
    <!-- Just graph time spent doing pomodoros in given period -->
    <!-- Ranges: Last Week, Last Month, Last Year(by Month), All Time(by Month) -->
    <!-- Hours on the y axis -->

    <!-- Avg Study Time by Weekday -->

    <!-- Avg  study time by hour of the day -->

    <Visualization data={graphData.threemonths}/>

    <div class="grid">
        <StatSquare label="Today" stat="3"/>
        <StatSquare label="Last Week" stat="30"/>
        <StatSquare label="Last Month" stat="70"/>
        <StatSquare label="Last Year" stat="700"/>
        <StatSquare label="All Time" stat="2222"/>
        <StatSquare label="Most Time Day" stat="20h"/>
        <StatSquare label="Total Time" stat="1d 2h 10m"/>
        <StatSquare label="Avg Time" stat="30m"/>
    </div>

    {:else}
    <p>No data yet. Complete pomodoros to see some stats!</p>
    {/if}

    {:catch error}

    {error}

    {/await}


</div>

<style>

    .grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1em;
        padding:0;
    }

</style>