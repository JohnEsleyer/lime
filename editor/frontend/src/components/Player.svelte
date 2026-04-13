<script>
    export let workspaceOpen;
    export let streamPort;
    export let currentTime;
    export let isPlaying;
    export let togglePlay;

    let containerWidth;

    $: frameUrl =
        workspaceOpen && streamPort
            ? `http://127.0.0.1:${streamPort}/frame?t=${currentTime.toFixed(2)}&rnd=${currentTime}`
            : null;
</script>

<div class="player-area" bind:clientWidth={containerWidth}>
    <div class="toolbar-top">
        <div class="tool-group">
            <button class="tool-btn" title="Selection Tool">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><path d="m3 3 7.07 16.97 2.51-7.39 7.39-2.51L3 3z"
                    ></path><path d="m13 13 6 6"></path></svg
                >
            </button>
            <button class="tool-btn" title="Razor Tool">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><path d="M14 12V2"></path><path d="M4 12V2"></path><path
                        d="M14 20v-8H4v8a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2Z"
                    ></path><path d="M21 15v-5a2 2 0 0 0-2-2H4"></path><path
                        d="M21 7V5a2 2 0 0 0-2-2H4"
                    ></path></svg
                >
            </button>
        </div>

        <div class="viewer-tabs">
            <button class="v-tab active">Program</button>
            <button class="v-tab">Source</button>
        </div>

        <div class="tool-group">
            <select class="zoom-select">
                <option>Fit</option>
                <option>25%</option>
                <option>50%</option>
                <option>100%</option>
            </select>
        </div>
    </div>

    <div class="monitor">
        {#if frameUrl}
            <img src={frameUrl} alt="Preview" class="preview-frame" />
        {:else}
            <div class="render-placeholder">
                <div class="spinner"></div>
                <span
                    >{workspaceOpen
                        ? "Preparing renderer..."
                        : "No Project"}</span
                >
            </div>
        {/if}
    </div>

    <div class="controls">
        <div class="time-display current">{currentTime.toFixed(2)}</div>

        <div class="playback-btns">
            <button class="p-btn" title="Previous Frame">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><polygon points="19 20 9 12 19 4 19 20"></polygon><line
                        x1="5"
                        y1="19"
                        x2="5"
                        y2="5"
                    ></line></svg
                >
            </button>
            <button
                class="p-btn play-pause"
                on:click={togglePlay}
                title={isPlaying ? "Pause" : "Play"}
            >
                {#if isPlaying}
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="currentColor"
                        ><rect x="6" y="4" width="4" height="16"></rect><rect
                            x="14"
                            y="4"
                            width="4"
                            height="16"
                        ></rect></svg
                    >
                {:else}
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="currentColor"
                        ><polygon points="5 3 19 12 5 21 5 3"></polygon></svg
                    >
                {/if}
            </button>
            <button class="p-btn" title="Next Frame">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><polygon points="5 4 15 12 5 20 5 4"></polygon><line
                        x1="19"
                        y1="5"
                        x2="19"
                        y2="19"
                    ></line></svg
                >
            </button>
        </div>

        <div class="time-display duration">00:00:15:00</div>
    </div>

    <div class="transport-bar">
        <input
            type="range"
            min="0"
            max="15"
            step="0.01"
            bind:value={currentTime}
            class="timeline-slider"
        />
    </div>
</div>

<style>
    .player-area {
        flex: 1;
        background: #000000;
        display: flex;
        flex-direction: column;
        min-width: 0;
    }

    .toolbar-top {
        height: 48px;
        background: #111111;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 0 16px;
    }

    .tool-group {
        display: flex;
        gap: 8px;
    }

    .tool-btn {
        background: transparent;
        border: none;
        color: #94a3b8;
        padding: 6px;
        border-radius: 4px;
        cursor: pointer;
    }

    .tool-btn:hover {
        background: rgba(255, 255, 255, 0.05);
        color: white;
    }

    .viewer-tabs {
        display: flex;
        gap: 4px;
        background: #000000;
        padding: 2px;
        border-radius: 6px;
    }

    .v-tab {
        background: transparent;
        border: none;
        color: #64748b;
        font-size: 0.75rem;
        font-weight: 600;
        padding: 6px 16px;
        border-radius: 4px;
        cursor: pointer;
    }

    .v-tab.active {
        background: #111111;
        color: #a3e635;
    }

    .zoom-select {
        background: #000000;
        border: 1px solid rgba(255, 255, 255, 0.1);
        color: #94a3b8;
        font-size: 0.7rem;
        padding: 4px 8px;
        border-radius: 4px;
        outline: none;
    }

    .monitor {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
        overflow: hidden;
        position: relative;
        padding: 20px;
    }

    .preview-frame {
        max-width: 100%;
        max-height: 100%;
        box-shadow: 0 10px 40px rgba(0, 0, 0, 0.5);
        background: #000;
    }

    .render-placeholder {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 16px;
        color: #404040;
    }

    .spinner {
        width: 40px;
        height: 40px;
        border: 3px solid rgba(163, 230, 53, 0.1);
        border-top-color: #a3e635;
        border-radius: 50%;
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }

    .controls {
        height: 48px;
        padding: 0 20px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        background: #111111;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
    }

    .time-display {
        font-family: "JetBrains Mono", monospace;
        font-size: 0.85rem;
        color: #a3a3a3;
        background: #000000;
        padding: 4px 10px;
        border-radius: 4px;
        min-width: 80px;
        text-align: center;
    }

    .time-display.current {
        color: #a3e635;
    }

    .playback-btns {
        display: flex;
        align-items: center;
        gap: 16px;
    }

    .p-btn {
        background: transparent;
        border: none;
        color: #cbd5e1;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .p-btn:hover {
        color: white;
    }

    .p-btn.play-pause {
        width: 44px;
        height: 44px;
        background: #a3e635;
        color: #000000;
        border-radius: 50%;
        transition: transform 0.1s;
    }

    .p-btn.play-pause:hover {
        transform: scale(1.05);
        background: #bef264;
    }

    .transport-bar {
        padding: 0;
        background: #111111;
    }

    .timeline-slider {
        width: 100%;
        margin: 0;
        cursor: pointer;
        background: transparent;
        appearance: none;
        height: 4px;
        display: block;
    }

    .timeline-slider::-webkit-slider-runnable-track {
        background: #262626;
        height: 4px;
    }

    .timeline-slider::-webkit-slider-thumb {
        appearance: none;
        height: 12px;
        width: 12px;
        border-radius: 50%;
        background: #a3e635;
        margin-top: -4px;
        box-shadow: 0 0 10px rgba(163, 230, 53, 0.5);
    }
</style>
