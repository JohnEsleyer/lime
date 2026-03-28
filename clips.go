package lime

import (
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
