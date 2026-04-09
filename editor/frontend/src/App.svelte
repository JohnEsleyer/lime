<script>
  import { onMount, onDestroy } from "svelte";
  import {
    GetStreamPort,
    ImportAsset,
    ExportVideo,
    OpenWorkspace,
    SaveWorkspace,
  } from "../wailsjs/go/main/App.js";
  import "./style.css";

  let currentTime = 0;
  let isPlaying = false;
  let playInterval;

  let workspaceOpen = false;
  let projectConfig = null;
  let streamPort = 0;

  // Internal generic state mapping to projectConfig
  let clips = [];

  async function handleOpenWorkspace() {
    let result = await OpenWorkspace();
    if (result) {
      projectConfig = JSON.parse(result);
      projectConfig.media = projectConfig.media || [];
      if (
        projectConfig &&
        projectConfig.tracks &&
        projectConfig.tracks.length > 0
      ) {
        clips = projectConfig.tracks[0].clips || [];
      }
      workspaceOpen = true;
    }
  }

  async function syncSave() {
    if (workspaceOpen && projectConfig) {
      projectConfig.tracks[0].clips = clips;
      await SaveWorkspace(JSON.stringify(projectConfig));
    }
  }

  async function handleImport() {
    if (!workspaceOpen) {
      alert("Open a workspace first!");
      return;
    }
    let asset = await ImportAsset();
    if (asset) {
      projectConfig.media = [...projectConfig.media, asset];
      await syncSave();
    }
  }

  async function handleExport() {
    await ExportVideo("output.avi");
    alert("Export complete!");
  }

  function togglePlay() {
    isPlaying = !isPlaying;
    if (isPlaying) {
      playInterval = setInterval(() => {
        currentTime += 0.1;
        if (currentTime > 15) {
          // Reset loop placeholder
          currentTime = 0;
        }
      }, 100); // Note: natively driven by interval updates the timestamp img tag Reactively!
    } else {
      clearInterval(playInterval);
    }
  }

  onMount(async () => {
    streamPort = await GetStreamPort();
  });

  onDestroy(() => {
    if (playInterval) clearInterval(playInterval);
  });
</script>

<div class="topbar">
  <div class="title">
    🍋 {projectConfig?.name ? projectConfig.name.toUpperCase() : "LIME EDITOR"}
  </div>
  <div style="display: flex; gap: 10px">
    <button class="btn" on:click={handleOpenWorkspace}>
      {workspaceOpen ? "Change Workspace" : "Open Workspace"}
    </button>
    <button class="btn" on:click={handleExport}>Export Video</button>
  </div>
</div>

<div class="workspace">
  <div class="panel media-bin">
    <h3>Media Bin</h3>
    <button
      class="btn"
      on:click={handleImport}
      style="width: 100%; margin-bottom: 10px;">+ Import</button
    >
    <div
      style="font-size: 0.8rem; color: #94a3b8; text-align: left; padding: 10px 0; display: flex; flex-direction: column; gap: 8px;"
    >
      {#if !workspaceOpen}
        <div style="text-align: center; padding: 20px;">
          Please Open Workspace
        </div>
      {:else if projectConfig && projectConfig.media.length > 0}
        {#each projectConfig.media as asset}
          <div
            style="background: rgba(255,255,255,0.05); padding: 8px; border-radius: 4px; display: flex; align-items: center; gap: 10px; cursor: pointer;"
          >
            <div
              style="width: 40px; height: 40px; background: #000; border-radius: 4px; display: flex; align-items: center; justify-content: center; font-size: 0.7rem; overflow: hidden;"
            >
              {#if asset.thumbnail}
                <img
                  src={asset.thumbnail}
                  alt="thumb"
                  style="width: 100%; height: 100%; object-fit: cover;"
                />
              {:else}
                <span style="opacity: 0.5;">FILE</span>
              {/if}
            </div>
            <div
              style="flex: 1; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;"
              title={asset.path}
            >
              {asset.path.split("/").pop()}
            </div>
          </div>
        {/each}
      {:else}
        <div style="text-align: center; padding: 20px;">Drag assets here</div>
      {/if}
    </div>
  </div>

  <div class="player">
    <div class="preview-container">
      {#if workspaceOpen && streamPort}
        <img
          src={`http://127.0.0.1:${streamPort}/frame?t=${currentTime.toFixed(2)}&rnd=${currentTime}`}
          alt="Preview"
        />
      {:else}
        <div
          style="width: 100%; height: 100%; display: flex; align-items: center; justify-content: center; color: #475569"
        >
          {workspaceOpen
            ? "Connecting to Render Engine..."
            : "No Project Opened"}
        </div>
      {/if}
    </div>
    <div
      style="margin-top: 20px; display: flex; gap: 10px; align-items: center;"
    >
      <button class="btn" on:click={togglePlay} style="width: 80px;"
        >{isPlaying ? "Pause" : "Play"}</button
      >
      <input
        type="range"
        min="0"
        max="15"
        step="0.1"
        bind:value={currentTime}
        style="width: 300px;"
      />
      <span style="font-family: monospace;"
        >Time: {currentTime.toFixed(1)}s</span
      >
    </div>
  </div>

  <div class="panel inspector">
    <h3>Properties</h3>
    <p style="font-size: 0.8rem; color: #94a3b8;">
      Select an item to view properties.
    </p>
  </div>
</div>

<div class="timeline">
  <div class="timeline-header">Timeline</div>
  <div class="tracks">
    <div class="track">
      {#each clips as clip}
        <div
          class="clip"
          style="left: {clip.start * 30}px; width: {clip.duration *
            30}px; background: {clip.color}; box-shadow: 0 2px 10px {clip.color}40"
        >
          {clip.name}
        </div>
      {/each}
    </div>
    <div class="track"></div>
    <div class="track"></div>
  </div>
</div>
