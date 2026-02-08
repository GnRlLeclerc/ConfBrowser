package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

// writeJSON writes JSON data to an http response
func writeJSONResponse[T any](w http.ResponseWriter, data T) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}

// writeOk writes a default Ok response for endpoints that do not need to return anything.
// It returns an empty JSON object because the frontend expects JSON responses
func writeOk(w http.ResponseWriter) {
	writeJSONResponse(w, map[string]any{})
}

var tabCount atomic.Int64

// MountHandler registers a new opened tab connected to the backend.
// The backend auto closes once no tab is left.
func MountHandler(w http.ResponseWriter, r *http.Request) {
	tabCount.Add(1)

	writeOk(w)
}

// ExitHandler exits the backend when called.
// This is used to close the server when the user closes the webapp tab
func ExitHandler(w http.ResponseWriter, r *http.Request) {
	// Do not go to negative values (happens when reloading a tab that was disconnected)
	if tabCount.Load() <= 0 {
		w.WriteHeader(http.StatusOK)
		writeOk(w)
		return
	}

	if tabCount.Add(-1) <= 0 {
		go func() {
			// Wait 10 second for a reconnection
			time.Sleep(10 * time.Second)

			if tabCount.Load() > 0 {
				return
			}

			fmt.Printf("Shutting down server...")
			os.Exit(0)
		}()
	}

	writeOk(w)
}

// GetPapersHandler scrapes the papers list for a given conference
// and year.
func GetPapersHandler(w http.ResponseWriter, r *http.Request) {
	conf := r.PathValue("conf")
	year := r.PathValue("year")

	// Try to get the papers from the file cache first
	papers, err := cache.GetPapers(conf, year)
	if err != nil {
		log.Printf("Cache miss for %s %s: %v", conf, year, err)
	} else {
		writeJSONResponse(w, papers)
		return
	}

	// If not found in cache, scrape and cache them
	papers, err = scrapePapers(conf, year)
	if err != nil {
		http.Error(w, "Failed to get papers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cache.WritePapers(conf, year, papers); err != nil {
		log.Printf("Failed to write papers to cache: %v", err)
	}
	writeJSONResponse(w, papers)
}

func GetPaperHandler(w http.ResponseWriter, r *http.Request) {
	conf := r.PathValue("conf")
	year := r.PathValue("year")
	id := r.PathValue("id")

	// Try to get the paper from the file cache first
	paper, err := cache.GetPaper(conf, year, id)
	if err != nil {
		log.Printf("Cache miss for paper %s %s %s: %v", conf, year, id, err)
	} else {
		writeJSONResponse(w, paper)
		return
	}

	// If not found in cache, scrape and cache it
	paper, err = scrapePaper(conf, year, id)
	if err != nil {
		http.Error(w, "Failed to get paper: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cache.WritePaper(conf, year, id, paper); err != nil {
		log.Printf("Failed to write paper to cache: %v", err)
	}

	writeJSONResponse(w, paper)
}
