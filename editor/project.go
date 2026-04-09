package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ProjectConfig struct {
	Name       string       `json:"name"`
	Version    string       `json:"version"`
	Resolution struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"resolution"`
	FPS    int          `json:"fps"`
	Tracks []TrackConfig `json:"tracks"`
	Media  []MediaAsset  `json:"media"`
}

type MediaAsset struct {
	ID        string `json:"id"`
	Path      string `json:"path"` // relative to workspace/assets/
	Thumbnail string `json:"thumbnail"`
}

type TrackConfig struct {
	ID    int          `json:"id"`
	Clips []ClipConfig `json:"clips"`
}

type ClipConfig struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Start    float64 `json:"start"`
	Duration float64 `json:"duration"`
	Color    string  `json:"color"`
}

// EnsureDefaultProject creates or returns an active project.json from a given workspace path.
func EnsureDefaultProject(workspacePath string) (*ProjectConfig, error) {
	projPath := filepath.Join(workspacePath, "project.json")

	if _, err := os.Stat(projPath); os.IsNotExist(err) {
		// Create default empty project
		cfg := &ProjectConfig{
			Name:    filepath.Base(workspacePath),
			Version: "1.0",
			FPS:     30,
		}
		cfg.Resolution.Width = 1920
		cfg.Resolution.Height = 1080
		cfg.Tracks = []TrackConfig{
			{ID: 1, Clips: []ClipConfig{}},
		}

		b, _ := json.MarshalIndent(cfg, "", "  ")
		if err := os.WriteFile(projPath, b, 0644); err != nil {
			return nil, err
		}
		return cfg, nil
	}

	// Read existing project config
	b, err := os.ReadFile(projPath)
	if err != nil {
		return nil, err
	}
	var cfg ProjectConfig
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
