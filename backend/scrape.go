package backend

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// scrapePapers scrapes the main papers list from a given conference and year
func scrapePapers(conf string, year string) ([]Paper, error) {
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
