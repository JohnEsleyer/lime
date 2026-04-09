package main

import (
	"bytes"
	"image/jpeg"
	"net"
	"net/http"
	"strconv"

	"github.com/johnesleyer/lime"
)

type Streamer struct {
	port  int
	video *lime.Video // reference to current video
}

func startStreamer(video *lime.Video) (*Streamer, error) {
	// Find an open port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	
	port := listener.Addr().(*net.TCPAddr).Port
	s := &Streamer{port: port, video: video}

	// Setup endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/frame", s.handleFrame)
	
	go http.Serve(listener, mux) // Run the server concurrently
	
	return s, nil
}

func (s *Streamer) handleFrame(w http.ResponseWriter, r *http.Request) {
	tStr := r.URL.Query().Get("t")
	if tStr == "" {
		http.Error(w, "missing parameter t", http.StatusBadRequest)
		return
	}
	tSec, err := strconv.ParseFloat(tStr, 64)
	if err != nil {
		http.Error(w, "invalid t parameter", http.StatusBadRequest)
		return
	}

	if s.video == nil || s.video.Timeline == nil {
		http.Error(w, "no video loaded", http.StatusNotFound)
		return
	}

	// Always allow CORS for local frontend requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")

	img := s.video.Timeline.RenderFrame(tSec, 800, 450)
	
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 60})
	
	w.Write(buf.Bytes())
}
