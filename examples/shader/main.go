package main

import (
	"image/color"
	"log"
	"math"

	"github.com/johnesleyer/lime"
)

func main() {
	// 1. Create a 720p 30fps video project
	video := lime.New(lime.Config{
		Width:   1280,
		Height:  720,
		FPS:     30,
		Bitrate: 90,
	})

	// 2. Add an animation track
	animTrack := video.Timeline.AddTrack()
	
	// Complex animated gradient shader
	animTrack.AddClip(0, lime.NewShaderClip(5.0, func(x, y int, w, h int, localTime float64) color.Color {
		// Normalize coordinates from 0 to 1
		u := float64(x) / float64(w)
		v := float64(y) / float64(h)

		// Create shifting waves based on localTime
		r := 0.5 + 0.5*math.Sin(u*10.0+localTime*4.0)
		g := 0.5 + 0.5*math.Cos(v*8.0-localTime*3.0)
		b := 0.5 + 0.5*math.Sin((u+v)*5.0+localTime*2.0)

		return color.RGBA{
			R: uint8(r * 255),
			G: uint8(g * 255),
			B: uint8(b * 255),
			A: 255,
		}
	}))

	// 3. Export the video
	err := video.Export("shader_output.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
