package lime

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/jpeg"
	"os"
)

// AVIEncoder writes an MJPEG AVI file using pure Go standard library.
type AVIEncoder struct {
	file       *os.File
	width      int
	height     int
	fps        int
	quality    int
	frameCount int
	moviStart  int64
	index      []aviIndexEntry
}

type aviIndexEntry struct {
	chunkId string
	flags   uint32
	offset  uint32
	size    uint32
}

// NewAVIEncoder creates a new AVI output file.
func NewAVIEncoder(path string, width, height, fps, quality int) (*AVIEncoder, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	enc := &AVIEncoder{
		file:    file,
		width:   width,
		height:  height,
		fps:     fps,
		quality: quality,
	}

	if err := enc.writeHeader(); err != nil {
		file.Close()
		return nil, err
	}

	return enc, nil
}

// AddFrame encodes a single image to JPEG and appends it to the AVI.
func (e *AVIEncoder) AddFrame(img image.Image) error {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: e.quality})
	if err != nil {
		return err
	}
	
	jpegData := buf.Bytes()
	// Pad to an even size
	pad := 0
	if len(jpegData)%2 != 0 {
		jpegData = append(jpegData, 0)
		pad = 1
	}

	chunkOffset, _ := e.file.Seek(0, 1)

	// Write 00dc chunk
	e.file.Write([]byte("00dc"))
	if err := binary.Write(e.file, binary.LittleEndian, uint32(len(jpegData)-pad)); err != nil {
		return err
	}
	e.file.Write(jpegData)

	// Keep track for index
	e.index = append(e.index, aviIndexEntry{
		chunkId: "00dc",
		flags:   0x10, // AVIIF_KEYFRAME
		offset:  uint32(chunkOffset - e.moviStart - 8),
		size:    uint32(len(jpegData) - pad),
	})

	e.frameCount++
	return nil
}

// Close finishes the AVI file by writing the index and fixing up header sizes.
func (e *AVIEncoder) Close() error {
	// Write idx1 chunk
	idx1Offset, _ := e.file.Seek(0, 1)
	e.file.Write([]byte("idx1"))
	idx1Size := uint32(len(e.index) * 16)
	binary.Write(e.file, binary.LittleEndian, idx1Size)

	for _, entry := range e.index {
		e.file.Write([]byte(entry.chunkId))
		binary.Write(e.file, binary.LittleEndian, entry.flags)
		binary.Write(e.file, binary.LittleEndian, entry.offset)
		binary.Write(e.file, binary.LittleEndian, entry.size)
	}

	fileSize, _ := e.file.Seek(0, 1)

	// Update RIFF size
	e.file.Seek(4, 0)
	binary.Write(e.file, binary.LittleEndian, uint32(fileSize-8))

	// Update Frames in main headers
	e.file.Seek(48, 0)
	binary.Write(e.file, binary.LittleEndian, uint32(e.frameCount))
	
	// Stream length in frames
	e.file.Seek(140, 0)
	binary.Write(e.file, binary.LittleEndian, uint32(e.frameCount))

	// movi list size
	e.file.Seek(e.moviStart+4, 0)
	binary.Write(e.file, binary.LittleEndian, uint32(idx1Offset-e.moviStart-8))

	return e.file.Close()
}

func (e *AVIEncoder) writeHeader() error {
	w := e.file
	// Reserve space for the header blocks, writing zeroes for frames/sizes that we fill in `Close()`.
	w.Write([]byte("RIFF"))
	binary.Write(w, binary.LittleEndian, uint32(0)) // RIFF size
	w.Write([]byte("AVI "))

	w.Write([]byte("LIST"))
	binary.Write(w, binary.LittleEndian, uint32(192)) // hdrl LIST size
	w.Write([]byte("hdrl"))

	// avih header
	w.Write([]byte("avih"))
	binary.Write(w, binary.LittleEndian, uint32(56)) // avih size
	usecPerFrame := uint32(1000000 / e.fps)
	binary.Write(w, binary.LittleEndian, usecPerFrame) // MicroSecPerFrame
	binary.Write(w, binary.LittleEndian, uint32(0))    // MaxBytesPerSec
	binary.Write(w, binary.LittleEndian, uint32(0))    // PaddingGranularity
	binary.Write(w, binary.LittleEndian, uint32(0x10)) // Flags (has index)
	binary.Write(w, binary.LittleEndian, uint32(0))    // TotalFrames (updated in Close)
	binary.Write(w, binary.LittleEndian, uint32(0))    // InitialFrames
	binary.Write(w, binary.LittleEndian, uint32(1))    // Streams
	binary.Write(w, binary.LittleEndian, uint32(0))    // SuggestedBufferSize
	binary.Write(w, binary.LittleEndian, uint32(e.width))
	binary.Write(w, binary.LittleEndian, uint32(e.height))
	// reserved
	binary.Write(w, binary.LittleEndian, []uint32{0, 0, 0, 0})

	// strl LIST
	w.Write([]byte("LIST"))
	binary.Write(w, binary.LittleEndian, uint32(116)) // strl LIST size
	w.Write([]byte("strl"))

	// strh header
	w.Write([]byte("strh"))
	binary.Write(w, binary.LittleEndian, uint32(56)) // strh size
	w.Write([]byte("vids"))
	w.Write([]byte("MJPG"))
	binary.Write(w, binary.LittleEndian, uint32(0)) // Flags
	binary.Write(w, binary.LittleEndian, uint16(0)) // Priority
	binary.Write(w, binary.LittleEndian, uint16(0)) // Language
	binary.Write(w, binary.LittleEndian, uint32(0)) // InitialFrames
	binary.Write(w, binary.LittleEndian, uint32(1)) // Scale
	binary.Write(w, binary.LittleEndian, uint32(e.fps)) // Rate
	binary.Write(w, binary.LittleEndian, uint32(0)) // Start
	binary.Write(w, binary.LittleEndian, uint32(0)) // Length (updated in Close)
	binary.Write(w, binary.LittleEndian, uint32(0)) // SuggestedBufferSize
	binary.Write(w, binary.LittleEndian, uint32(10000)) // Quality
	binary.Write(w, binary.LittleEndian, uint32(0))     // SampleSize
	binary.Write(w, binary.LittleEndian, []uint16{0, 0, 0, 0}) // Frame (2xRECT) 

	// strf header
	w.Write([]byte("strf"))
	binary.Write(w, binary.LittleEndian, uint32(40)) // strf size (BITMAPINFOHEADER)
	binary.Write(w, binary.LittleEndian, uint32(40)) // Size
	binary.Write(w, binary.LittleEndian, uint32(e.width))
	binary.Write(w, binary.LittleEndian, uint32(e.height))
	binary.Write(w, binary.LittleEndian, uint16(1))  // Planes
	binary.Write(w, binary.LittleEndian, uint16(24)) // BitCount
	w.Write([]byte("MJPG"))                          // Compression
	binary.Write(w, binary.LittleEndian, uint32(0))  // SizeImage
	binary.Write(w, binary.LittleEndian, uint32(0))  // XPelsPerMeter
	binary.Write(w, binary.LittleEndian, uint32(0))  // YPelsPerMeter
	binary.Write(w, binary.LittleEndian, uint32(0))  // ClrUsed
	binary.Write(w, binary.LittleEndian, uint32(0))  // ClrImportant

	// movi LIST
	moviStartOffset, _ := e.file.Seek(0, 1)
	e.moviStart = moviStartOffset

	w.Write([]byte("LIST"))
	binary.Write(w, binary.LittleEndian, uint32(0)) // movi size (updated in Close)
	w.Write([]byte("movi"))

	return nil
}
