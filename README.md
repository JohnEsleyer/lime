# 🍋 Lime

Lime is a high-performance, pure-Go video editing and generation library. Inspired by Remotion and Revideo, Lime allows developers to create complex video animations, motion graphics, and shader-driven visuals directly in Go code.

## ✨ Features

- **0 External Dependencies**: No FFmpeg, no CGO, no system-level libraries required. Just `go get` and you're ready.
- **Pure-Go MJPEG-AVI Encoder**: Built-in native Go implementation of an AVI muxer for instant video export.
- **Canvas-style API**: Powerful 2D drawing API (paths, arcs, beziers, transforms) powered by `fogleman/gg`.
- **Software Pixel Shaders**: Concurrent, CPU-based shader support for mathematical graphics and complex pixel-level animations.
- **Layered Timeline**: Organized track-based system for compositing multiple clips over time.

## 🚀 Getting Started

### Installation

```bash
go get github.com/user/lime
```

### Basic Example

```go
package main

import (
	"github.com/fogleman/gg"
	"github.com/user/lime"
)

func main() {
	// Create a 720p 30fps project
	video := lime.New(lime.Config{Width: 1280, Height: 720, FPS: 30})

	// Add a drawing track
	track := video.Timeline.AddTrack()

	// Add a 5-second animated clip
	track.AddClip(0, lime.NewCanvasClip(5.0, func(dc *gg.Context, localTime float64) {
		dc.SetHexColor("#ff0000")
		dc.DrawCircle(640, 360, localTime * 50)
		dc.Fill()
	}))

	// Export to video!
	video.Export("output.avi")
}
```

## 🧠 Why Pure Go?

Most video libraries are wrappers around FFmpeg or C-based engines. Lime is different:
- **Easy Deployment**: Binaries are completely static and portable.
- **Zero Configuration**: No need to worry about shared library versions or external tool paths.
- **Cross-Platform**: Compiles and runs anywhere Go does (Linux, macOS, Windows, WASM).

## 🛠 Advanced: CPU Shaders

Lime includes a `ShaderClip` that lets you write pixel shaders in Go. It runs concurrently across all CPU cores:

```go
track.AddClip(0, lime.NewShaderClip(5.0, func(x, y, w, h int, localTime float64) color.Color {
    u := float64(x) / float64(w)
    v := float64(y) / float64(h)
    return color.RGBA{uint8(u*255), uint8(v*255), uint8(math.Sin(localTime)*255), 255}
}))
```

## ⚖️ License

MIT
