<svelte:head>
    <title>Settings</title>
</svelte:head>

<script>
    import MinInput from "$lib/MinInput.svelte";
    import RadioBtns from "$lib/RadioBtns.svelte";
    import { userAPI } from "$lib/requests.js"
    import { settings, isAuthenticated} from "$lib/store.js";
    import { setCSS } from "$lib/colors.js";

    // TODO change the way variables are reactive

    let pomodoroInput = $settings.pomodoro.toString();
    let shortBreakInput = $settings.shortBreak.toString();
    let longBreakInput = $settings.longBreak.toString();

    const options = {
        theme: ["Default", "Shore", "Cold", "Slate", "Ocean", "Pumkin", "Pastel", "Night", "Pink", "Lettuce", "Purp", "Cream", "Little", "Cafe", "Contrast", "Eve"],
        sound: ["8bit", "Bike", "Bowl", "Cathedral", "Coin", "Correct", "Ding", "Dingdong", "Hibell", "Lowbell", "Notify", "Peaceful", "Ping", "Pling"],
        volume: ["Mute", "Low", "Medium", "High"],
    };

    function isPositiveInt(str) {
        const num = Number(str);

        if (Number.isInteger(num) && num > 0) {
            return true;
        }

        return false;
    }

    // autosave
    let interval = null;
    function autosave() {
        clearTimeout(interval);
        interval = setTimeout(() => {
            $settings.pomodoro = Number(pomodoroInput);
            $settings.shortBreak = Number(shortBreakInput);
            $settings.longBreak = Number(longBreakInput);

            if ($isAuthenticated) {
                userAPI("put", "/settings", $settings)
                .then(resp => console.log(resp))
                .catch(error => {
                    console.log(error);
                    console.log($settings);
                });
            }
        }, 500);
    }

    $: {
        setCSS($settings.theme);
    }



    const audio = new Map();
    options.sound.forEach(elem => audio.set(elem, new Audio("/sounds/" + elem + ".mp3")))
    const volume = new Map([
        ["Mute", 0],
        ["Low", 0.4],
        ["Medium", 0.7],
        ["High", 1.0],
    ]);

    // TODO make sounds play on click instead of on change
    function playSound() {
        audio.get($settings.sound).volume = volume.get($settings.volume);
        audio.get($settings.sound).currentTime = 0;
        audio.get($settings.sound).play();
    }

    let currentSound = $settings.sound
    $: if (currentSound != $settings.sound) {
        playSound();
        currentSound = $settings.sound;
    }

    let currentVol = $settings.volume
    $: if (currentVol != $settings.volume) {
        playSound();
        currentVol = $settings.volume;
    }


</script>

<div class="container">
<form on:change={autosave} on:submit|preventDefault>
    <h1>Settings</h1>

    <h2>Timer</h2>

    <h3>Pomodoro</h3>
    <p>Customize the length of your pomodoro session.</p>
    <MinInput bind:value={pomodoroInput}/>
    {#if !isPositiveInt(pomodoroInput)}
    <span class="err">Input must be a positive integer.</span>
    {/if}

    <h3>Short Break</h3>
    <p>Customize your short break time.</p>
    <MinInput bind:value={shortBreakInput}/>
    {#if !isPositiveInt(shortBreakInput)}
    <span class="err">Input must be a positive integer.</span>
    {/if}

    <h3>Long Break</h3>
    <p>Customize your long break time.</p>
    <MinInput bind:value={longBreakInput}/>
    {#if !isPositiveInt(longBreakInput)}
    <span class="err">Input must be a positive integer.</span>
    {/if}

    <h2>Style</h2>

    <h3>Theme</h3>
    <p>Pick a theme that fits your style.</p>
    <RadioBtns
        name="theme"
        values={options.theme}
        nPerRow=4
        bind:selected={$settings.theme}
    />

    <h2>Audio</h2>

    <h3>Sound</h3>
    <p>Pick a sound to play when you complete a pomodoro.</p>
    <RadioBtns
        name="sound-type"
        values={options.sound}
        nPerRow=4
        bind:selected={$settings.sound}
    />

    <h3>Volume</h3>
    <p>The volume of your sound effects.</p>
    <RadioBtns
        name="sound-volume"
        values={options.volume}
        nPerRow=4
        bind:selected={$settings.volume}
    />

    <!-- TODO test danger zone -->
    {#if $isAuthenticated }
    <h2>Danger Zone</h2>

    <h3>Erase Statistics</h3>
    <p>Erase all saved data, including task history.</p>
    <button class="danger-btn rounded">Erase</button>

    <h3>Reset Settings</h3>
    <p>Revert settings to the default state.</p>
    <button class="danger-btn rounded">Reset</button>

    <h3>Delete Account</h3>
    <p>Delete you account and erase all associated data.</p>
    <button
        class="danger-btn rounded"
        on:click={() => userAPI("delete", "", null)}>
        Delete
    </button>
    {/if}

</form>
</div>
<style>
 
	.danger-btn {
        border: 2px solid var(--error);
		color: var(--error);
        background-color: transparent;
		height: 2.5em;
		font-size: medium;
		transition: all 0.1s;
        padding: 0 1em;
        margin: 0 0;
	}

    .danger-btn:hover {
        background-color: var(--error-hover)
    }

	.danger-btn:active {
		background-color: var(--error);
		color: var(--text-on-main);
	}


  /* @media (min-width: 480px) {
    h1 {
      max-width: none;
    }

    p {
      max-width: none;
    }


</style>
