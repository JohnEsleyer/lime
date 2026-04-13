<script>
    export let selectedClip = null;
</script>

<div class="sidebar-right">
    <div class="section-header">
        <h3>Inspector</h3>
    </div>

    <div class="inspector-content">
        {#if selectedClip}
            <div class="property-group">
                <label for="clip-name">Clip Name</label>
                <input
                    id="clip-name"
                    type="text"
                    bind:value={selectedClip.name}
                />
            </div>

            <div class="property-group">
                <label for="clip-start">Start (s)</label>
                <div class="input-with-unit">
                    <input
                        id="clip-start"
                        type="number"
                        step="0.1"
                        bind:value={selectedClip.start}
                    />
                    <span>s</span>
                </div>
            </div>

            <div class="property-group">
                <label for="clip-duration">Duration (s)</label>
                <div class="input-with-unit">
                    <input
                        id="clip-duration"
                        type="number"
                        step="0.1"
                        bind:value={selectedClip.duration}
                    />
                    <span>s</span>
                </div>
            </div>

            <div class="property-group">
                <label for="clip-color">Color</label>
                <div class="color-picker">
                    <input
                        id="clip-color"
                        type="color"
                        bind:value={selectedClip.color}
                    />
                    <span class="color-hex">{selectedClip.color}</span>
                </div>
            </div>

            <div class="divider"></div>

            <div class="property-group">
                <label for="clip-opacity">Opacity</label>
                <input
                    id="clip-opacity"
                    type="range"
                    min="0"
                    max="1"
                    step="0.01"
                    value="1"
                />
            </div>

            <div class="property-group">
                <label for="clip-scale">Scale</label>
                <input
                    id="clip-scale"
                    type="range"
                    min="0.1"
                    max="5"
                    step="0.1"
                    value="1"
                />
            </div>

            <div class="prototype-actions">
                <button class="proto-btn">Split Clip</button>
                <button class="proto-btn danger">Delete</button>
            </div>
        {:else}
            <div class="empty-state">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="48"
                    height="48"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="1"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><circle cx="12" cy="12" r="10"></circle><line
                        x1="12"
                        y1="16"
                        x2="12"
                        y2="12"
                    ></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg
                >
                <p>
                    Select a clip in the timeline to view and edit its
                    properties.
                </p>
            </div>
        {/if}
    </div>
</div>

<style>
    .sidebar-right {
        width: 280px;
        background: #111111;
        border-left: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    .section-header {
        padding: 12px 16px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
    }

    .section-header h3 {
        margin: 0;
        font-size: 0.85rem;
        text-transform: uppercase;
        letter-spacing: 0.05em;
        color: #a3a3a3;
    }

    .inspector-content {
        flex: 1;
        overflow-y: auto;
        padding: 20px 16px;
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .property-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .property-group label {
        font-size: 0.75rem;
        color: #525252;
        font-weight: 600;
        text-transform: uppercase;
    }

    .property-group input[type="text"],
    .property-group input[type="number"] {
        background: #000000;
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 4px;
        padding: 6px 10px;
        color: white;
        font-size: 0.85rem;
    }

    .input-with-unit {
        position: relative;
        display: flex;
        align-items: center;
    }

    .input-with-unit input {
        width: 100%;
    }

    .input-with-unit span {
        position: absolute;
        right: 10px;
        font-size: 0.75rem;
        color: #404040;
    }

    .color-picker {
        display: flex;
        align-items: center;
        gap: 10px;
        background: #000000;
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 4px;
        padding: 4px 8px;
    }

    .color-picker input {
        background: none;
        border: none;
        width: 30px;
        height: 24px;
        cursor: pointer;
    }

    .color-hex {
        font-family: monospace;
        font-size: 0.85rem;
        color: #d4d4d4;
    }

    .divider {
        height: 1px;
        background: rgba(255, 255, 255, 0.05);
        margin: 8px 0;
    }

    .prototype-actions {
        margin-top: 20px;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    .proto-btn {
        padding: 8px;
        background: rgba(255, 255, 255, 0.05);
        border: 1px solid rgba(255, 255, 255, 0.1);
        color: white;
        font-size: 0.8rem;
        border-radius: 4px;
        cursor: pointer;
    }

    .proto-btn:hover {
        background: rgba(255, 255, 255, 0.1);
    }

    .proto-btn.danger {
        color: #ef4444;
    }

    .empty-state {
        text-align: center;
        padding: 60px 20px;
        color: #404040;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 16px;
    }

    .empty-state p {
        font-size: 0.85rem;
        line-height: 1.5;
    }
</style>
