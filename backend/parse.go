package backend

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParsePapers parses the HTML contents of a conference papers page
// by extracting the list of posters.
//
// This function works for NeurIPS, ICLR, ICML and CVPR papers pages.
func parsePapers(html string) ([]Paper, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	var papers []Paper

	// Find the noscript tag that hides the papers list
	doc.Find("noscript").Each(func(i int, s *goquery.Selection) {
		doc, err = goquery.NewDocumentFromReader(strings.NewReader(s.Text()))
		if err != nil {
			return
		}

		// Parse all paper links
		doc.Find("a").Each(func(j int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if !exists {
				return
			}

			parts := strings.Split(href, "/")
			id := parts[len(parts)-1]
			title := strings.TrimSpace(link.Text())

			papers = append(papers, Paper{
				ID:    id,
				Title: title,
			})
		})
	})

	return papers, nil
}
