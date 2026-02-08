package backend

// Base urls for the conferences

const (
	baseNeurIPS = "https://neurips.cc"
	baseICML    = "https://icml.cc"
	baseICLR    = "https://iclr.cc"
	baseCVPR    = "https://cvpr.thecvf.com"
)

var (
	baseURL = map[string]string{
		"NeurIPS": baseNeurIPS,
		"ICML":    baseICML,
		"ICLR":    baseICLR,
		"CVPR":    baseCVPR,
	}
)

// Main page urls (contain paper metadata in a noscript tag)
const (
	papersNeurIPS = "https://neurips.cc/virtual/%s/papers.html"
	papersICML    = "https://icml.cc/virtual/%s/papers.html"
	papersICLR    = "https://iclr.cc/virtual/%s/papers.html"
	papersCVPR    = "https://cvpr.thecvf.com/virtual/%s/papers.html"
)

var (
	papersURL = map[string]string{
		"NeurIPS": papersNeurIPS,
		"ICML":    papersICML,
		"ICLR":    papersICLR,
		"CVPR":    papersCVPR,
	}
)

// Single poster urls (year, paper id)
const (
	posterNeurIPS = "https://neurips.cc/virtual/%s/poster/%s"
	posterICML    = "https://icml.cc/virtual/%s/poster/%s"
	posterICLR    = "https://iclr.cc/virtual/%s/poster/%s"
	posterCVPR    = "https://cvpr.thecvf.com/virtual/%s/poster/%s"
)

var (
	posterURL = map[string]string{
		"NeurIPS": posterNeurIPS,
		"ICML":    posterICML,
		"ICLR":    posterICLR,
		"CVPR":    posterCVPR,
	}
)
