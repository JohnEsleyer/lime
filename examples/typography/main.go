package main

import (
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/johnesleyer/lime"
	"golang.org/x/image/font/gofont/gobold"
)

func main() {
	video := lime.New(lime.Config{Width: 1280, Height: 720, FPS: 30, Bitrate: 90})

	// Parse built-in Go font (so we don't need external TTF files)
	font, err := truetype.Parse(gobold.TTF)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}
	// Create a large 120pt font face
	face := truetype.NewFace(font, &truetype.Options{Size: 120})

	// Add background
	bg := video.Timeline.AddTrack()
	bg.AddClip(0, lime.NewCanvasClip(10.0, func(dc *gg.Context, localTime float64) {
		dc.SetHexColor("#282a36") // Dark Dracula theme background
		dc.Clear()
	}))

	track := video.Timeline.AddTrack()

	// 1. Fade-in Transition (0s - 3s)
	track.AddClip(0.0, lime.NewCanvasClip(3.0, func(dc *gg.Context, t float64) {
		dc.SetFontFace(face)
		text := "Fade In"

		// Calculate opacity (fade in over 1 second, hold for 1 sec, fade out for 1 sec)
		alpha := 1.0
		if t < 1.0 {
			alpha = t // fade in linearly
		} else if t > 2.0 {
			alpha = 3.0 - t // fade out linearly
		}

		dc.SetRGBA(1, 1, 1, alpha)
		dc.DrawStringAnchored(text, 640, 360, 0.5, 0.5)
	}))

	// 2. Slide Transition + Drop Shadow (3s - 6s)
	track.AddClip(3.0, lime.NewCanvasClip(3.0, func(dc *gg.Context, t float64) {
		dc.SetFontFace(face)
		text := "Slide & Shadow"

		// Easing function (ease out cubic)
		progress := t / 3.0
		ease := 1.0 - math.Pow(1.0-progress, 3)
		x := -400.0 + (1040.0 * ease) // Slides perfectly to 640 (center)

		// Drop Shadow (draw slightly offset with opacity)
		dc.SetRGBA(0, 0, 0, 0.5)
		dc.DrawStringAnchored(text, x+12, 360+12, 0.5, 0.5)

		// Main Text
		dc.SetHexColor("#ffb86c") 
		dc.DrawStringAnchored(text, x, 360, 0.5, 0.5)
	}))

	// 3. Scaling & Rotation & Outline styling (6s - 10s)
	track.AddClip(6.0, lime.NewCanvasClip(4.0, func(dc *gg.Context, t float64) {
		dc.SetFontFace(face)
		text := "LIME VIDEO"

		// Scaling pulse via math.Sin
		scale := 1.0 + 0.15*math.Sin(t*math.Pi*2) 

		// Transform the context
		dc.Translate(640, 360)
		dc.Scale(scale, scale)

		// Rotation wobble
		dc.Rotate(math.Sin(t) * 0.1)

		// Text Outline (Stroke) Hack: 
		// Draw the text multiple times in a circle to create a thick outline
		dc.SetHexColor("#ff5555") // Red outline color
		outlineThickness := 6.0
		for angle := 0.0; angle < math.Pi*2; angle += math.Pi / 4 {
			ox := math.Cos(angle) * outlineThickness
			oy := math.Sin(angle) * outlineThickness
			dc.DrawStringAnchored(text, ox, oy, 0.5, 0.5)
		}

		// Text Inner Fill
		dc.SetHexColor("#f8f8f2")
		dc.DrawStringAnchored(text, 0, 0, 0.5, 0.5)
	}))

	err = video.Export("typography.avi")
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}
}
