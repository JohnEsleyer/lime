package lime

import (
	"image"
	"image/color"
	"sync"

	"github.com/fogleman/gg"
)

// CanvasClip is a generic clip that allows the user to write an arbitrary drawing function using an HTML5 Canvas-like API.
type CanvasClip struct {
	Dur    float64
	DrawFn func(ctx *gg.Context, localTime float64)
}

func (c *CanvasClip) Duration() float64 {
	return c.Dur
}

func (c *CanvasClip) Draw(ctx *gg.Context, localTime float64) {
	if c.DrawFn != nil {
		c.DrawFn(ctx, localTime)
	}
}

// NewCanvasClip is a helper to easily create custom drawn clips.
func NewCanvasClip(duration float64, drawFn func(ctx *gg.Context, localTime float64)) *CanvasClip {
	return &CanvasClip{
		Dur:    duration,
		DrawFn: drawFn,
	}
}

// ShaderPixelFn defines how a single pixel's color is calculated for a given coordinate (x,y) at time t.
// Time is normalized (0.0 to duration). For infinite effects, just use localTime.
type ShaderPixelFn func(x, y int, width, height int, localTime float64) color.Color

// ShaderClip generates an image mathematically pixel-by-pixel, efficiently running the calculation concurrently.
type ShaderClip struct {
	Dur    float64
	Shader ShaderPixelFn
}

func (c *ShaderClip) Duration() float64 {
	return c.Dur
}

func (c *ShaderClip) Draw(ctx *gg.Context, localTime float64) {
	if c.Shader == nil {
		return
	}

	bounds := ctx.Image().Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	
	// Create an intermediate RGBA to draw directly pixel-by-pixel
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Determine how many workers based on logical cores. 
	// For simplicity, we chunk by rows.
	var wg sync.WaitGroup
	// To avoid overloading goroutines per scanline, chunk the image.
	chunkY := 32
	
	for yStart := 0; yStart < h; yStart += chunkY {
		yEnd := yStart + chunkY
		if yEnd > h {
			yEnd = h
		}
		
		wg.Add(1)
		go func(ys, ye int) {
			defer wg.Done()
			for y := ys; y < ye; y++ {
				for x := 0; x < w; x++ {
					c := c.Shader(x, y, w, h, localTime)
					img.Set(x, y, c)
				}
			}
		}(yStart, yEnd)
	}
	wg.Wait()

	// Draw the result onto the gg Context
	ctx.DrawImage(img, 0, 0)
}

// NewShaderClip is a helper to easily create a video clip driven purely by a pixel shader function.
func NewShaderClip(duration float64, shaderFn ShaderPixelFn) *ShaderClip {
	return &ShaderClip{
		Dur:    duration,
		Shader: shaderFn,
	}
}
