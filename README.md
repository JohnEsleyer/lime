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
go get github.com/johnesleyer/lime
```

### Basic Example

```go
package main

import (
	"github.com/fogleman/gg"
	"github.com/johnesleyer/lime"
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

## 📑 API Cheatsheet

### Core Objects
| Type       | Description                                                          |
| :--------- | :------------------------------------------------------------------- |
| `Config`   | Video settings: `Width`, `Height`, `FPS`, `Bitrate` (quality 1-100). |
| `Video`    | Main container. Create with `lime.New(config)`.                      |
| `Timeline` | Manager for tracks. Accessed via `video.Timeline`.                   |
| `Track`    | A horizontal layer in the timeline.                                  |
| `Clip`     | Interface for drawable elements over a duration.                     |

### Main Functions
| Function                     | Description                                           |
| :--------------------------- | :---------------------------------------------------- |
| `lime.New(Config)`           | Creates a new Video project.                          |
| `video.Export(path)`         | Renders and encodes the project to an MJPEG AVI file. |
| `timeline.AddTrack()`        | Adds a new rendering layer.                           |
| `track.AddClip(start, clip)` | Places a clip on a track at a specific start time.    |

### Visual Elements
| Function                 | Description                                 |
| :----------------------- | :------------------------------------------ |
| `NewCanvasClip(dur, fn)` | Uses HTML5-style canvas API for drawing.    |
| `NewShaderClip(dur, fn)` | High-performance per-pixel software shader. |

### Common Drawing (via `gg.Context`)
| Function                    | Description                                             |
| :-------------------------- | :------------------------------------------------------ |
| `dc.SetHexColor(hex)`       | Sets drawing color from a hex string (e.g., "#FF0000"). |
| `dc.DrawRectangle(x,y,w,h)` | Defines a rectangular path.                             |
| `dc.DrawCircle(x,y,r)`      | Defines a circular path.                                |
| `dc.Fill()`                 | Fills the current path with the current color.          |
| `dc.Stroke()`               | Outlines the current path with the current color.       |
| `dc.DrawString(s,x,y)`      | Draws text at the specified location.                   |
| `dc.Clear()`                | Clears the entire canvas with the current color.        |

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
