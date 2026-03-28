package lime

import (
	"fmt"
	"log"
)

// Config holds the main settings for a Video process.
type Config struct {
	Width   int
	Height  int
	FPS     int
	Bitrate int // Used for JPEG compression quality (1-100)
}

// Video represents a complete video project to be rendered.
type Video struct {
	Config   Config
	Timeline *Timeline
}

// New creates a new video project.
func New(config Config) *Video {
	if config.FPS == 0 {
		config.FPS = 30
	}
	if config.Bitrate == 0 {
		config.Bitrate = 85
	}
	return &Video{
		Config:   config,
		Timeline: NewTimeline(),
	}
}

// Export renders the video timeline to a given output path (e.g., "output.avi").
func (v *Video) Export(outputPath string) error {
	log.Printf("Starting export to %s", outputPath)
	encoder, err := NewAVIEncoder(outputPath, v.Config.Width, v.Config.Height, v.Config.FPS, v.Config.Bitrate)
	if err != nil {
		return fmt.Errorf("failed to initialize AVI encoder: %w", err)
	}
	defer encoder.Close()

	totalDuration := v.Timeline.Duration()
	totalFrames := int(totalDuration * float64(v.Config.FPS))

	for frame := 0; frame <= totalFrames; frame++ {
		timeSec := float64(frame) / float64(v.Config.FPS)
		
		// Render frame
		img := v.Timeline.RenderFrame(timeSec, v.Config.Width, v.Config.Height)
		
		// Encode and write frame
		if err := encoder.AddFrame(img); err != nil {
			return fmt.Errorf("failed to encode frame %d: %w", frame, err)
		}
		
		if frame%10 == 0 || frame == totalFrames {
			log.Printf("Rendered frame %d/%d (%.2fs/%.2fs)", frame, totalFrames, timeSec, totalDuration)
		}
	}
	log.Printf("Export complete!")
	return nil
}
