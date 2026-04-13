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

  // Components
  import ProjectManager from "./components/ProjectManager.svelte";
  import SidebarLeft from "./components/SidebarLeft.svelte";
  import SidebarRight from "./components/SidebarRight.svelte";
  import Player from "./components/Player.svelte";
  import Timeline from "./components/Timeline.svelte";

  let currentTime = 0;
  let isPlaying = false;
  let playInterval;

  let workspaceOpen = false;
  let projectConfig = null;
  let streamPort = 0;

  // Track state
  let clips = [];
  let selectedClip = null;

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
      // Ensure we preserve the tracks structure
      if (!projectConfig.tracks) projectConfig.tracks = [{}];
      projectConfig.tracks[0].clips = clips;
      await SaveWorkspace(JSON.stringify(projectConfig));
    }
  }

  async function handleImport() {
    if (!workspaceOpen) return;
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
        currentTime += 0.05; // Finer precision for visual smoothness
        if (currentTime > 15) {
          currentTime = 0;
        }
      }, 50);
    } else {
      clearInterval(playInterval);
    }
  }

  function handleSelectClip(clip) {
    selectedClip = clip;
  }

  function handleCloseProject() {
    workspaceOpen = false;
    projectConfig = null;
    clips = [];
    selectedClip = null;
  }

  onMount(async () => {
    streamPort = await GetStreamPort();
  });

  onDestroy(() => {
    if (playInterval) clearInterval(playInterval);
  });
</script>

<div class="app-container">
  {#if !workspaceOpen}
    <ProjectManager onOpenWorkspace={handleOpenWorkspace} />
  {:else}
    <!-- Top Bar -->
    <div class="topbar">
      <div class="title">
        🍋 {projectConfig?.name?.toUpperCase() || "LIME EDITOR"}
      </div>
      <div style="display: flex; gap: 10px">
        <button class="btn" on:click={handleCloseProject}>Close Project</button>
        <button class="btn btn-primary" on:click={handleExport}
          >Export Video</button
        >
      </div>
    </div>

    <!-- Main Editor -->
    <div class="main-editor">
      <div class="upper-panels">
        <SidebarLeft {projectConfig} {handleImport} />

        <Player
          {workspaceOpen}
          {streamPort}
          bind:currentTime
          {isPlaying}
          {togglePlay}
        />

        <SidebarRight {selectedClip} />
      </div>

      <!-- Timeline -->
      <Timeline {clips} {currentTime} onSelectClip={handleSelectClip} />
    </div>
  {/if}
</div>

<style>
  /* Local overrides if necessary, but most is in style.css */
</style>
