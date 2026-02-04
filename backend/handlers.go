package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// writeJSON writes JSON data to an http response
func writeJSON[T any](w http.ResponseWriter, data T) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}

// ExitHandler exits the backend when called.
// This is used to close the server when the user closes the webapp tab
func ExitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Shutting down server...")
	os.Exit(0)
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
		writeJSON(w, papers)
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
	writeJSON(w, papers)
}
