<svelte:head>
    <title>Stats</title>
</svelte:head>

<script>
    import { userAPI } from "$lib/requests.js"
    import StatSquare from "$lib/StatSquare.svelte";
    import RadioBtns from "$lib/RadioBtns.svelte";
    import Visualization from "$lib/Visualization.svelte";
    import { pomodoros, contentLoading } from "$lib/store.js";

    // Retrieve data from database.
    async function getData() {
        contentLoading.set(true);

        const resp = await userAPI("get", "/stats", null);
        console.log(resp);
        // pomodoros.set(resp);


        contentLoading.set(false);

        return resp;
    }

</script>

<div>
    <h1>Stats</h1>
    <p>
        Keep track of how many pomodoros you have completed. Every completed
        pomodoro counts towards your stats.
    </p>

    {#await getData() then resp}

    <!-- {#if $pomodoros.length > 0} -->

    <h3>Period of Time</h3>
    <RadioBtns
        values={["7 days", "30 days", "90 days"]}
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

    <Visualization data={resp.Bar7Day}/>

    <div class="grid">
        <!-- TODO convert to readable format such as 1d 2h 30m -->
        <StatSquare label="Today" stat="{resp.Today}m"/>
        <StatSquare label="Last Week" stat="{resp.Week}m"/>
        <StatSquare label="Last Month" stat="{resp.Month}m"/>
        <StatSquare label="Last Year" stat="{resp.Year}m"/>
        <StatSquare label="All Time" stat="{resp.All}m"/>
        <StatSquare label="Avg" stat="{resp.AllAvg}m"/>
    </div>

    <!-- {:else}
    <p>No data yet. Complete pomodoros to see some stats!</p>
    {/if} -->

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