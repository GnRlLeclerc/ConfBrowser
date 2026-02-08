package backend

// PaperMetadata represents a conference paper with an ID and Title,
// as it appears in the main papers list.
type PaperMetadata struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Paper represents detailed information about a paper
type Paper struct {
	Title    string   `json:"title"`
	Authors  []string `json:"authors"`
	Abstract string   `json:"abstract"`
	PDF      string   `json:"pdf"`
	Poster   string   `json:"poster,omitempty"`
}
