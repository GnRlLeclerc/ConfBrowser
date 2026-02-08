package backend

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// scrapePapers scrapes the main papers list from a given conference and year
func scrapePapers(conf string, year string) ([]PaperMetadata, error) {
	url := papersURL[conf]
	url = fmt.Sprintf(url, year)

	log.Println("Fetching papers from URL:", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	html := string(data)
	return parsePapers(html)
}

// scrapePaper scrapes a single paper's details from a given conference, year, and paper id
func scrapePaper(conf string, year string, id string) (Paper, error) {
	url := posterURL[conf]
	url = fmt.Sprintf(url, year, id)

	log.Println("Fetching paper from URL:", url)
	response, err := http.Get(url)
	if err != nil {
		return Paper{}, fmt.Errorf("HTTP request error: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Paper{}, fmt.Errorf("status error: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Paper{}, fmt.Errorf("read body error: %v", err)
	}

	html := string(data)
	return parsePaper(html, conf)
}
