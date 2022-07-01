<svelte:head>
    <title>Stats</title>
</svelte:head>

<script>
    import { userAPI } from "$lib/requests.js"
    import StatSquare from "$lib/StatSquare.svelte";
    import RadioBtns from "$lib/RadioBtns.svelte";
    import Visualization from "$lib/Visualization.svelte";
    import WeekdayViz from "$lib/WeekdayViz.svelte";
    import { pomodoros, contentLoading } from "$lib/store.js";

    // TODO make graph axis colors match text color so its visible on dark backgrounds

    // Time period selection for bar graph
    let period = "7 Days";

    // Retrieve data from database.
    async function getData() {
        contentLoading.set(true);

        const resp = await userAPI("get", "/stats", null);
        console.log(resp);

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

    {#if resp.errorMessage}
    <p>No data yet. Complete pomodoros to see some stats!</p>
    {:else}
    <!-- {#if $pomodoros.length > 0} -->

    <h3>Period of Time</h3>
    <p>View the number of minutes studied in a given period.</p>
    <RadioBtns
        values={["7 Days", "14 Days", "30 Days"]}
        name="period"
        bind:selected={period}
        nPerRow=3
    />

    {#key period}
    <Visualization data={resp[period]}/>
    {/key}

    <h3>Weekday Summary</h3>
    <p>The total minutes studied during the last 30 days, grouped by day of the week.</p>
    <WeekdayViz data={resp.WeekdayData}/>

    <div class="grid">
        <!-- TODO convert to readable format such as 1d 2h 30m -->
        <StatSquare label="Today" stat="{resp.Today}m"/>
        <StatSquare label="Last Week" stat="{resp.Week}m"/>
        <StatSquare label="Last Month" stat="{resp.Month}m"/>
        <StatSquare label="Last Year" stat="{resp.Year}m"/>
        <StatSquare label="All Time" stat="{resp.All}m"/>
        <StatSquare label="Avg" stat="{resp.AllAvg}m"/>
    </div>


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
