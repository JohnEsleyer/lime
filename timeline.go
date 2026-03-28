package lime

import (
	"image"

	"github.com/fogleman/gg"
)

// Clip interface represents anything that can draw itself on a frame context at a specific local time.
type Clip interface {
	Duration() float64
	Draw(ctx *gg.Context, localTime float64)
}

// TrackClip holds a Clip and its start time on the track.
type TrackClip struct {
	StartTime float64
	Clip      Clip
}

// Track is a collection of clips. 
type Track struct {
	Clips []TrackClip
}

// AddClip adds a clip to the track at a given start time.
func (t *Track) AddClip(startTime float64, clip Clip) {
	t.Clips = append(t.Clips, TrackClip{
		StartTime: startTime,
		Clip:      clip,
	})
}

// Duration returns the total duration of the track based on its furthest clip.
func (t *Track) Duration() float64 {
	max := 0.0
	for _, tc := range t.Clips {
		end := tc.StartTime + tc.Clip.Duration()
		if end > max {
			max = end
		}
	}
	return max
}

// Timeline organizes tracks and renders them back-to-front.
type Timeline struct {
	Tracks []*Track
}

// NewTimeline creates an empty timeline.
func NewTimeline() *Timeline {
	return &Timeline{}
}

// AddTrack adds an empty track to the timeline and returns it.
func (t *Timeline) AddTrack() *Track {
	track := &Track{}
	t.Tracks = append(t.Tracks, track)
	return track
}

// Duration returns the total duration of all tracks.
func (t *Timeline) Duration() float64 {
	max := 0.0
	for _, tr := range t.Tracks {
		dur := tr.Duration()
		if dur > max {
			max = dur
		}
	}
	return max
}

// RenderFrame renders all clips active at timeSec to a new image.
func (t *Timeline) RenderFrame(timeSec float64, width, height int) image.Image {
	dc := gg.NewContext(width, height)
	// Base transparent background
	dc.SetRGBA(0, 0, 0, 0)
	dc.Clear()

	// Draw tracks layer by layer (first track is bottom-most)
	for _, tr := range t.Tracks {
		for _, tc := range tr.Clips {
			if timeSec >= tc.StartTime && timeSec <= tc.StartTime+tc.Clip.Duration() {
				localTime := timeSec - tc.StartTime
				
				// Push state so clip transformations don't affect others
				dc.Push()
				tc.Clip.Draw(dc, localTime)
				dc.Pop()
			}
		}
	}

	return dc.Image()
}
