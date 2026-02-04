package backend

// Paper represents a conference paper with an ID and Title,
// as it appears in the main papers list.
type Paper struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
