<script>
    export let clips;
    export let currentTime;
    export let onSelectClip;

    let timelineWidth;
    const pixelsPerSecond = 40;

    // Selected clip index tracking
    let selectedIdx = -1;

    function handleClipClick(idx, clip) {
        selectedIdx = idx;
        onSelectClip(clip);
    }
</script>

<div class="timeline-container" bind:clientWidth={timelineWidth}>
    <div class="timeline-toolbar">
        <div class="time-readout">
            <span>{currentTime.toFixed(2)}</span>
        </div>
        <div class="timeline-actions">
            <button class="t-btn" title="Add Track">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><line x1="12" y1="5" x2="12" y2="19"></line><line
                        x1="5"
                        y1="12"
                        x2="19"
                        y2="12"
                    ></line></svg
                >
            </button>
            <button class="t-btn" title="Snap to Grid">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><rect x="3" y="3" width="18" height="18" rx="2" ry="2"
                    ></rect><line x1="3" y1="9" x2="21" y2="9"></line><line
                        x1="3"
                        y1="15"
                        x2="21"
                        y2="15"
                    ></line><line x1="9" y1="3" x2="9" y2="21"></line><line
                        x1="15"
                        y1="3"
                        x2="15"
                        y2="21"
                    ></line></svg
                >
            </button>
        </div>
    </div>

    <div class="timeline-workarea">
        <div class="track-headers">
            <div class="track-header">
                <div class="track-title">Video 1</div>
                <div class="track-controls">
                    <button class="ctrl-btn">V</button>
                    <button class="ctrl-btn">L</button>
                </div>
            </div>
            <div class="track-header">
                <div class="track-title">Audio 1</div>
                <div class="track-controls">
                    <button class="ctrl-btn">M</button>
                    <button class="ctrl-btn">S</button>
                </div>
            </div>
        </div>

        <div class="tracks-container">
            <div
                class="playhead"
                style="left: {currentTime * pixelsPerSecond}px"
            ></div>

            <div class="track">
                {#each clips as clip, i}
                    <div
                        class="clip {selectedIdx === i ? 'selected' : ''}"
                        style="left: {clip.start *
                            pixelsPerSecond}px; width: {clip.duration *
                            pixelsPerSecond}px; background: {clip.color}; --glow: {clip.color}80"
                        on:click={() => handleClipClick(i, clip)}
                    >
                        <span class="clip-label">{clip.name}</span>
                        <div class="handle l"></div>
                        <div class="handle r"></div>
                    </div>
                {/each}
            </div>

            <div class="track"></div>

            <div class="ruler">
                {#each Array(20) as _, i}
                    <div
                        class="ruler-tick"
                        style="left: {i * pixelsPerSecond}px"
                    >
                        <span>{i}s</span>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</div>

<style>
    .timeline-container {
        height: 250px;
        background: #000000;
        border-top: 1px solid rgba(255, 255, 255, 0.1);
        display: flex;
        flex-direction: column;
        user-select: none;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
    }

    .timeline-toolbar {
        height: 36px;
        background: #111111;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        align-items: center;
        padding: 0 16px;
        gap: 20px;
    }

    .time-readout {
        font-family: "JetBrains Mono", monospace;
        font-size: 0.8rem;
        color: #a3e635;
        background: #000000;
        padding: 2px 8px;
        border-radius: 4px;
        min-width: 60px;
        text-align: center;
    }

    .timeline-actions {
        display: flex;
        gap: 8px;
    }

    .t-btn {
        background: transparent;
        border: none;
        color: #525252;
        cursor: pointer;
        padding: 4px;
        border-radius: 4px;
    }

    .t-btn:hover {
        background: rgba(255, 255, 255, 0.05);
        color: white;
    }

    .timeline-workarea {
        flex: 1;
        display: flex;
        overflow: hidden;
    }

    .track-headers {
        width: 150px;
        background: #111111;
        border-right: 1px solid rgba(255, 255, 255, 0.05);
        z-index: 10;
    }

    .track-header {
        height: 60px;
        padding: 8px 12px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }

    .track-title {
        font-size: 0.75rem;
        font-weight: 600;
        color: #a3a3a3;
    }

    .track-controls {
        display: flex;
        gap: 4px;
    }

    .ctrl-btn {
        background: #000000;
        border: 1px solid rgba(255, 255, 255, 0.1);
        color: #404040;
        font-size: 0.6rem;
        padding: 2px 4px;
        border-radius: 2px;
        cursor: pointer;
    }

    .ctrl-btn:hover {
        color: white;
        border-color: rgba(255, 255, 255, 0.2);
    }

    .tracks-container {
        flex: 1;
        overflow-x: auto;
        overflow-y: hidden;
        position: relative;
        background-image: linear-gradient(
                rgba(255, 255, 255, 0.02) 1px,
                transparent 1px
            ),
            linear-gradient(
                90deg,
                rgba(255, 255, 255, 0.02) 1px,
                transparent 1px
            );
        background-size:
            100% 60px,
            40px 100%;
    }

    .playhead {
        position: absolute;
        top: 0;
        bottom: 0;
        width: 2px;
        background: #ef4444;
        z-index: 20;
        pointer-events: none;
    }

    .playhead::after {
        content: "";
        position: absolute;
        top: 0;
        left: -5px;
        border-left: 6px solid transparent;
        border-right: 6px solid transparent;
        border-top: 10px solid #ef4444;
    }

    .track {
        height: 60px;
        position: relative;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
    }

    .clip {
        position: absolute;
        top: 4px;
        bottom: 4px;
        border-radius: 4px;
        display: flex;
        align-items: center;
        padding: 0 10px;
        cursor: pointer;
        border: 1px solid rgba(0, 0, 0, 0.2);
        transition: box-shadow 0.2s;
    }

    .clip.selected {
        border: 2px solid white;
        box-shadow: 0 0 15px var(--glow);
        z-index: 5;
    }

    .clip-label {
        font-size: 0.75rem;
        font-weight: 600;
        color: white;
        text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .handle {
        position: absolute;
        top: 0;
        bottom: 0;
        width: 6px;
        background: rgba(0, 0, 0, 0.2);
        cursor: ew-resize;
    }

    .handle.l {
        left: 0;
    }
    .handle.r {
        right: 0;
    }

    .ruler {
        height: 20px;
        position: relative;
        background: #000000;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
    }

    .ruler-tick {
        position: absolute;
        top: 0;
        height: 100%;
        border-left: 1px solid rgba(255, 255, 255, 0.1);
        padding-left: 4px;
    }

    .ruler-tick span {
        font-size: 0.6rem;
        color: #404040;
    }

    /* Custom Scrollbar */
    .tracks-container::-webkit-scrollbar {
        height: 8px;
    }
    .tracks-container::-webkit-scrollbar-track {
        background: #000000;
    }
    .tracks-container::-webkit-scrollbar-thumb {
        background: #262626;
        border-radius: 4px;
    }
    .tracks-container::-webkit-scrollbar-thumb:hover {
        background: #404040;
    }
</style>
