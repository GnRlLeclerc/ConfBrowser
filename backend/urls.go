package backend

// Main page urls (contain paper metadata in a noscript tag)
var (
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
