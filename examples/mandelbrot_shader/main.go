package main

import (
	"image/color"
	"log"

	"github.com/johnesleyer/lime"
)

func main() {
	video := lime.New(lime.Config{Width: 640, Height: 480, FPS: 30, Bitrate: 90})

	track := video.Timeline.AddTrack()

	// High performance concurrency Mandelbrot zoom
	track.AddClip(0, lime.NewShaderClip(6.0, func(x, y int, w, h int, localTime float64) color.Color {
		// Calculate zoom level based on time
		zoom := 1.0 + (localTime * localTime * 2.0)
		
		// Map pixel coordinate to complex plane
		// Center on an interesting spot
		cx := -0.743643887037151
		cy := 0.131825904205330
		
		aspect := float64(w) / float64(h)
		scale := 2.5 / zoom
		
		u := (float64(x)/float64(w) - 0.5) * scale * aspect
		v := (float64(y)/float64(h) - 0.5) * scale
		
		zx, zy := 0.0, 0.0
		x0, y0 := u+cx, v+cy

		iteration := 0
		maxIter := 50 + int(zoom*10)

		for zx*zx+zy*zy <= 4 && iteration < maxIter {
			xtemp := zx*zx - zy*zy + x0
			zy = 2*zx*zy + y0
			zx = xtemp
			iteration++
		}

		if iteration == maxIter {
			return color.RGBA{0, 0, 0, 255}
		}

		// Colorize based on iterations
		smooth := float64(iteration) / float64(maxIter)
		
		r := uint8(255 * smooth)
		g := uint8(255 * (smooth * smooth))
		b := uint8(255 * (smooth * smooth * smooth))

		return color.RGBA{r, g, b, 255}
	}))

	err := video.Export("mandelbrot_zoom.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
