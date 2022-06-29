<script>
	import { isAuthenticated, settings } from "./store.js";
    import { userAPI } from "./requests.js";
    import RadioBtns from "./RadioBtns.svelte";

	let timeTotal;
    let timeStr;
    let timeRemaining;

    // Update timer based on mode selected
    let selected = "Pomodoro";
    $: selected, selectionUpdate();

    function selectionUpdate() {
        if (selected == "Pomodoro") {
            handleReset();
            timeTotal = $settings.pomodoro * 60 * 1000;
        } else if (selected == "Short Break") {
            handleReset();
            timeTotal = $settings.shortBreak * 60 * 1000;
        } else if (selected == "Long Break") {
            handleReset();
            timeTotal = $settings.longBreak * 60 * 1000;
        }
        timeRemaining = timeTotal;
        timeStr = msToString(timeRemaining)
    }


    // Update the timer if it was changed in the settings

	const circumference = 282.743339;
	let circleLength = 0;
    let interval;
    let initialTime = null;

    // Start timer countdown
    function handleStart() {
        // Save initial time for stats
        if (!initialTime) {
            initialTime = new Date();
        }

        // Check if countdown is already finished
        if (timeRemaining < 0) {
            return;
        }

        let end = Date.now() + timeRemaining;
        clearInterval(interval);
        interval = setInterval(() => {
            timeRemaining = end - Date.now();
            timeStr = msToString(timeRemaining);
			circleLength = circumference * (timeRemaining/timeTotal);
			if (timeRemaining < 0) {
                pomoComplete();
			}
        }, 100);
    }

    // Pause timer countdown
    function handlePause() {
        clearInterval(interval);
    }

    // Reset the timer countdown
    function handleReset() {
        initialTime = null;
        circleLength = 0;
        clearInterval(interval);
        timeRemaining = timeTotal;
        timeStr = msToString(timeRemaining);
    }

    // Convert milliseconds to a string for displaying to the user
    function msToString(ms) {
        ms /= 1000
        let s = Math.floor(ms % 60);
        let m = Math.floor(ms / 60 % 60);

        return (m < 10 ? "0" : "") + m + ":" + (s < 10 ? "0" : "") + s;
    }

    // Run when pomodoro is completed
    function pomoComplete() {
        clearInterval(interval);
        timeStr = "00:00";
        circleLength = 0;
        
        if ($isAuthenticated && selected == "Pomodoro") {
            postToDB();
        }
    }

    // post a finished pomodoro
    async function postToDB() {
        const timestamp = initialTime.toISOString();
        userAPI("post", "/stats", {timestamp, minutes: $settings.pomodoro})
        .then(resp => console.log(resp))
        .catch(error => console.log(error));
    }


</script>

<div class="timer-container">
    <!-- TODO add a startup animation -->
    <div class="btn-row">
        <RadioBtns
            values={["Pomodoro", "Short Break", "Long Break"]}
            name="timer-top-btns"
            nPerRow=3
            bind:selected={selected}
        />
    </div>
    <div class="timer">
		<svg version="1.1" viewBox="0 0 100 100">
			<circle id="background-circle" cx="50%" cy="50%" r="45"/>
			<circle id="animated-circle" style="stroke-dashoffset:{circleLength}" cx="50%" cy="50%" r="45"/>
            <text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle">{timeStr}</text>
        </svg>
    </div>
    <div class="btn-row ">
        <button class="clickable rounded" on:click={handleStart}>Start</button> 
        <button class="clickable rounded" on:click={handlePause}>Pause</button> 
        <button class="clickable rounded" on:click={handleReset}>Reset</button> 
    </div>
</div>

<style>
    .timer-container {
        max-width: 400px;
        margin: auto;
    }

    .btn-row {
        margin: 0 auto;
        display: flex;
        gap: 1em;
        width: 100%;
    }

    button {
        flex-grow: 1;
        padding: 0 1em;
	}

    .timer {
		display: grid;
		align-content: center;
		margin: auto;
		padding: 1em 0;
		width: 100%;
    }

    svg {
        width: 100%;
        height: 100%;
    }

	svg circle {
		fill: transparent;
		stroke-width: 0.35rem;
		transform-origin: center;
		transform: rotate(-90deg);
	}

	#background-circle {
		stroke: var(--bg-circle);
	}

	#animated-circle {
		stroke: var(--fg-circle);
		stroke-dasharray: 282.743339;
	}

    svg text {
        fill: var(--primary);
    }
</style>
