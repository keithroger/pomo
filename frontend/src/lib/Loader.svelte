<script>

</script>

<div class="bg">
    <div class="animation-container">
        <!-- <div style="color: var(--text-on-bg);">Loading...</div> -->
		<svg version="1.1" viewBox="0 0 100 100">
			<circle id="background-circle" cx="50%" cy="50%" r="45"/>
        </svg>
    </div>
</div>

<style>
    .bg {
        background-color: var(--bg);
        position: fixed;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
    }

    .animation-container {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    svg {
        width: 10em;
        height: 10em;
        animation: rotation 2s infinite linear;
    }

    circle {
        fill: transparent;
        stroke-linecap: round;
        stroke-dashoffset: 190;
        stroke-dasharray: 283;
        stroke: var(--primary);
        stroke-width: 10;
    }

    @keyframes rotation {
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(359deg);
        }
    }
</style>