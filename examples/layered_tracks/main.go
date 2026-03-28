package main

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/johnesleyer/lime"
)

func main() {
	video := lime.New(lime.Config{Width: 1280, Height: 720, FPS: 30, Bitrate: 90})

	// Background track (bottom layer)
	bg := video.Timeline.AddTrack()
	bg.AddClip(0, lime.NewCanvasClip(10.0, func(dc *gg.Context, localTime float64) {
		dc.SetHexColor("#1e1e2e")
		dc.Clear()
	}))

	// Middle layer (shapes appearing and disappearing)
	mid := video.Timeline.AddTrack()
	// Box 1 (0-4 seconds)
	mid.AddClip(0.0, lime.NewCanvasClip(4.0, func(dc *gg.Context, t float64) {
		dc.DrawRectangle(100+t*50, 200, 200, 200)
		dc.SetHexColor("#f9e2af")
		dc.Fill()
	}))

	// Box 2 (3-7 seconds)
	mid.AddClip(3.0, lime.NewCanvasClip(4.0, func(dc *gg.Context, t float64) {
		dc.DrawRectangle(400, 400-t*50, 200, 200)
		dc.SetHexColor("#f38ba8")
		dc.Fill()
	}))
	
	// Box 3 (6-10 seconds)
	mid.AddClip(6.0, lime.NewCanvasClip(4.0, func(dc *gg.Context, t float64) {
		dc.DrawRectangle(800-t*50, 200, 200, 200)
		dc.SetHexColor("#89b4fa")
		dc.Fill()
	}))

	// Foreground layer (global overlay)
	fg := video.Timeline.AddTrack()
	fg.AddClip(0, lime.NewCanvasClip(10.0, func(dc *gg.Context, t float64) {
		// Draw a static vignette overlay
		grad := gg.NewRadialGradient(640, 360, 0, 640, 360, 800)
		grad.AddColorStop(0, gg.Color{R: 0, G: 0, B: 0, A: 0})
		grad.AddColorStop(1, gg.Color{R: 0, G: 0, B: 0, A: 1})
		
		dc.SetFillStyle(grad)
		dc.DrawRectangle(0, 0, 1280, 720)
		dc.Fill()
	}))

	err := video.Export("layered_tracks.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
