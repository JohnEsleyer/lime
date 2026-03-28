package main

import (
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/user/lime"
)

func main() {
	// 1. Create a 720p 30fps video project
	video := lime.New(lime.Config{
		Width:   1280,
		Height:  720,
		FPS:     30,
		Bitrate: 90,
	})

	// 2. Add a background track
	bgTrack := video.Timeline.AddTrack()
	bgTrack.AddClip(0, lime.NewCanvasClip(5.0, func(dc *gg.Context, localTime float64) {
		dc.SetHexColor("#1e1e2e")
		dc.Clear()
	}))

	// 3. Add an animation track
	animTrack := video.Timeline.AddTrack()
	
	// A spinning, moving rectangle
	animTrack.AddClip(0, lime.NewCanvasClip(5.0, func(dc *gg.Context, localTime float64) {
		// Calculate position based on time (moves left to right)
		progress := localTime / 5.0
		x := 100.0 + progress*(1280-200)
		y := 360.0

		// Translate out context to center of rectangle
		dc.Translate(x, y)
		
		// Rotate based on time
		dc.Rotate(localTime * math.Pi)

		// Draw rectangle around origin
		dc.DrawRectangle(-100, -100, 200, 200)
		
		// Fill it
		dc.SetHexColor("#cba6f7")
		dc.FillPreserve()
		
		// Stroke it
		dc.SetRGB(1, 1, 1)
		dc.SetLineWidth(10)
		dc.Stroke()
	}))

	// 4. Export the video
	err := video.Export("output.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
