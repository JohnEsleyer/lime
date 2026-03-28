package main

import (
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/user/lime"
)

type Ball struct {
	X, Y   float64
	VX, VY float64
	Radius float64
	Color  string
}

func main() {
	video := lime.New(lime.Config{Width: 800, Height: 600, FPS: 60, Bitrate: 90})

	track := video.Timeline.AddTrack()

	balls := []Ball{
		{X: 100, Y: 100, VX: 300, VY: 200, Radius: 30, Color: "#f38ba8"},
		{X: 400, Y: 300, VX: -250, VY: 400, Radius: 50, Color: "#89b4fa"},
		{X: 700, Y: 500, VX: -400, VY: -300, Radius: 40, Color: "#a6e3a1"},
	}

	duration := 10.0

	track.AddClip(0, lime.NewCanvasClip(duration, func(dc *gg.Context, localTime float64) {
		// Draw background
		dc.SetHexColor("#11111b")
		dc.Clear()

		// Calculate physics step (since we query purely by time, we can calculate position exactly)
		// For perfect bouncing simulation based on time without a state loop:
		for i := range balls {
			b := balls[i]
			
			// x(t) = x0 + vx*t, modulo walls
			x := b.X + b.VX*localTime
			y := b.Y + b.VY*localTime

			// Calculate bounces using triangle waves
			ww := 800.0 - b.Radius*2
			hh := 600.0 - b.Radius*2

			// Map x to bounded range [b.Radius, 800-b.Radius]
			x = math.Mod(math.Abs(x), ww*2)
			if x > ww {
				x = ww*2 - x
			}
			x += b.Radius

			y = math.Mod(math.Abs(y), hh*2)
			if y > hh {
				y = hh*2 - y
			}
			y += b.Radius

			// Draw ball
			dc.DrawCircle(x, y, b.Radius)
			dc.SetHexColor(b.Color)
			dc.Fill()
		}
	}))

	err := video.Export("bouncing_balls.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
