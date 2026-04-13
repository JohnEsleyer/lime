<script>
    export let projectConfig;
    export let handleImport;

    $: media = projectConfig?.media || [];
</script>

<div class="sidebar-left">
    <div class="section-header">
        <h3>Library</h3>
        <button class="icon-btn" title="New Folder">
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
                ><path
                    d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"
                ></path><line x1="12" y1="11" x2="12" y2="17"></line><line
                    x1="9"
                    y1="14"
                    x2="15"
                    y2="14"
                ></line></svg
            >
        </button>
    </div>

    <div class="search-bar">
        <input type="text" placeholder="Search media..." />
    </div>

    <div class="media-container">
        <button class="import-btn" on:click={handleImport}>
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
                ><line x1="12" y1="5" x2="12" y2="19"></line><line
                    x1="5"
                    y1="12"
                    x2="19"
                    y2="12"
                ></line></svg
            >
            Import Media
        </button>

        <div class="media-grid">
            {#each media as asset}
                <div class="media-item">
                    <div class="thumbnail">
                        {#if asset.thumbnail}
                            <img src={asset.thumbnail} alt="thumb" />
                        {:else}
                            <div class="file-placeholder">
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="24"
                                    height="24"
                                    viewBox="0 0 24 24"
                                    fill="none"
                                    stroke="currentColor"
                                    stroke-width="2"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    ><path
                                        d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"
                                    ></path><polyline points="14 2 14 8 20 8"
                                    ></polyline><line
                                        x1="12"
                                        y1="18"
                                        x2="12"
                                        y2="12"
                                    ></line><line x1="9" y1="15" x2="15" y2="15"
                                    ></line></svg
                                >
                            </div>
                        {/if}
                    </div>
                    <div class="name">{asset.path.split("/").pop()}</div>
                </div>
            {:else}
                <div class="empty-state">No media yet.</div>
            {/each}
        </div>
    </div>

    <div class="sidebar-tabs">
        <button class="tab active">Media</button>
        <button class="tab">Effects</button>
        <button class="tab">Transitions</button>
    </div>
</div>

<style>
    .sidebar-left {
        width: 300px;
        background: #111111;
        border-right: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    .section-header {
        padding: 12px 16px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
    }

    .section-header h3 {
        margin: 0;
        font-size: 0.85rem;
        text-transform: uppercase;
        letter-spacing: 0.05em;
        color: #a3a3a3;
    }

    .icon-btn {
        background: transparent;
        border: none;
        color: #525252;
        cursor: pointer;
        padding: 4px;
        border-radius: 4px;
    }

    .icon-btn:hover {
        background: rgba(255, 255, 255, 0.05);
        color: white;
    }

    .search-bar {
        padding: 12px 16px;
    }

    .search-bar input {
        width: 100%;
        background: #000000;
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 6px;
        padding: 8px 12px;
        color: white;
        font-size: 0.85rem;
    }

    .media-container {
        flex: 1;
        overflow-y: auto;
        padding: 0 16px;
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .import-btn {
        width: 100%;
        padding: 12px;
        background: rgba(163, 230, 53, 0.1);
        border: 1px dashed rgba(163, 230, 53, 0.3);
        border-radius: 8px;
        color: #a3e635;
        font-weight: 600;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 8px;
        transition: all 0.2s;
    }

    .import-btn:hover {
        background: rgba(163, 230, 53, 0.2);
        border-color: #a3e635;
    }

    .media-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 12px;
    }

    .media-item {
        cursor: pointer;
        transition: transform 0.1s;
    }

    .media-item:hover {
        transform: scale(1.02);
    }

    .thumbnail {
        aspect-ratio: 16/9;
        background: #000000;
        border-radius: 6px;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        border: 1px solid rgba(255, 255, 255, 0.05);
    }

    .thumbnail img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .file-placeholder {
        color: #262626;
    }

    .name {
        font-size: 0.75rem;
        color: #d4d4d4;
        margin-top: 6px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .empty-state {
        grid-column: span 2;
        text-align: center;
        padding: 40px 0;
        color: #404040;
        font-size: 0.85rem;
    }

    .sidebar-tabs {
        display: flex;
        background: #111111;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
    }

    .tab {
        flex: 1;
        padding: 12px;
        background: transparent;
        border: none;
        color: #525252;
        font-size: 0.75rem;
        font-weight: 600;
        cursor: pointer;
        border-bottom: 2px solid transparent;
    }

    .tab.active {
        color: #a3e635;
        border-bottom-color: #a3e635;
    }
</style>
