package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const (
  DEFAULT_SOUND_FILE = "doorbell.wav"
)

func main() {
	// Play sound route
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get the sound file path from query parameter, or use default
		soundFile := r.URL.Query().Get("file")
		if soundFile == "" {
			soundFile = DEFAULT_SOUND_FILE
		}

		// Execute mpg321 to play the sound
		cmd := exec.Command("aplay", soundFile)
		err := cmd.Start()
		if err != nil {
			log.Printf("Error playing sound: %v", err)
			http.Error(w, "Failed to play sound", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Playing sound: %s", soundFile)
	})

	// Status route
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create status response with current date
		status := map[string]string{
			"date":    time.Now().Format(time.RFC3339),
			"status":  "running",
			"version": "1.0",
		}

		// Set content type and encode to JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(status); err != nil {
			http.Error(w, "Failed to generate status", http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
