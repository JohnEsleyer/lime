package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/johnesleyer/lime"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx           context.Context
	video         *lime.Video
	workspacePath string
	streamer      *Streamer
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Initialize a default video project for the editor
	a.video = lime.New(lime.Config{
		Width:   1920,
		Height:  1080,
		FPS:     30,
		Bitrate: 85,
	})

	// Start local HTTP Server for Live Previews
	if s, err := startStreamer(a.video); err == nil {
		a.streamer = s
	}
}

// GetStreamPort tells the frontend where the image stream is bound natively
func (a *App) GetStreamPort() int {
	if a.streamer != nil {
		return a.streamer.port
	}
	return 0
}

// OpenWorkspace opens a folder selection dialog for the user to select their project directory.
func (a *App) OpenWorkspace() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Workspace Folder",
	})
	if err != nil || dir == "" {
		return "", err
	}

	a.workspacePath = dir
	
	// Read or initialize project.json
	proj, err := EnsureDefaultProject(dir)
	if err != nil {
		return "", err
	}
	
	val, err := json.Marshal(proj)
	return string(val), err
}

// SaveWorkspace ensures the updated layout config string is persisted to the workspace project.json
func (a *App) SaveWorkspace(jsonData string) error {
	if a.workspacePath == "" {
		return nil // No workspace yet to save to
	}
	
	projPath := filepath.Join(a.workspacePath, "project.json")
	return os.WriteFile(projPath, []byte(jsonData), 0644)
}

// OpenFileDialog opens the native file dialog to pick an image or video file
func (a *App) OpenFileDialog() string {
	filepath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Media File",
	})
	if err != nil {
		return ""
	}
	return filepath
}

// RenderFramePreview renders a frame at the given timestamp and returns base64 PNG
// This enables the Svelte frontend preview window to show lime.Timeline directly.
func (a *App) RenderFramePreview(timeSec float64, width, height int) (string, error) {
	if a.video == nil || a.video.Timeline == nil {
		return "", nil
	}
	// Note: Wails frontend will frequently request this during Timeline scrubbing. 
	img := a.video.Timeline.RenderFrame(timeSec, width, height)
	
	// Encode frame to PNG
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return "", err
	}
	
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return "data:image/png;base64," + b64, nil
}

// ExportVideo triggers the video export mechanism
func (a *App) ExportVideo(outputPath string) error {
	if a.video == nil {
		return nil
	}
	// Note: You can add `runtime.EventsEmit` later for progress reporting
	return a.video.Export(outputPath)
}

// ImportAsset opens a dialog to select a media file and copies it to the workspace.
func (a *App) ImportAsset() (*MediaAsset, error) {
	if a.workspacePath == "" {
		return nil, fmt.Errorf("no workspace open")
	}

	filepathSelected, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Media File",
	})
	if err != nil || filepathSelected == "" {
		return nil, err
	}

	// Copy to workspacePath/assets/
	assetsDir := filepath.Join(a.workspacePath, "assets")
	os.MkdirAll(assetsDir, 0755)

	baseName := filepath.Base(filepathSelected)
	destPath := filepath.Join(assetsDir, baseName)

	in, err := os.Open(filepathSelected)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	io.Copy(out, in)

	// Thumbnail generation can be optimized later. For now, provide the path structure.
	return &MediaAsset{
		ID:   fmt.Sprintf("%d", time.Now().Unix()),
		Path: "assets/" + baseName,
		Thumbnail: "", 
	}, nil
}
