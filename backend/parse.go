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
func parsePapers(html string) ([]PaperMetadata, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	var papers []PaperMetadata

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

			papers = append(papers, PaperMetadata{
				ID:    id,
				Title: title,
			})
		})
	})

	return papers, nil
}

func parsePoster(doc *goquery.Document, conf string) string {
	url, exists := doc.Find("a[title='Poster']").First().Attr("href")

	if !exists {
		return ""
	}

	return baseURL[conf] + url
}

// parsePaper parses the page of a single paper.
func parsePaper(html string, conf string) (Paper, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return Paper{}, fmt.Errorf("error parsing HTML: %v", err)
	}

	return Paper{
		Title:    doc.Find("h1.event-title").First().Text(),
		Authors:  strings.Split(strings.TrimSpace(doc.Find("div.event-organizers").First().Text()), " · "),
		Abstract: doc.Find("div.abstract-text-inner").First().Text(),
		PDF:      strings.Replace(doc.Find("a[title='OpenReview']").First().AttrOr("href", ""), "forum", "pdf", 1),
		Poster:   parsePoster(doc, conf),
	}, nil

}
