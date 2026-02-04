package backend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// fallbackFS is a filesystem that defaults to index.html
type fallbackFS struct {
	fs fs.FS
}

func (f fallbackFS) Open(name string) (fs.File, error) {
	file, err := f.fs.Open(name)
	if err != nil {
		return f.fs.Open("index.html")
	}
	return file, nil
}

// MakeFileserverHandler creates a file server that serves files
// from an embedded filesystem and redirects unknown paths to index.html
func MakeFileserverHandler(filesystem embed.FS, subfolder string) http.Handler {
	subFS, err := fs.Sub(filesystem, subfolder)
	if err != nil {
		log.Fatal(err)
	}
	fallbackFS := fallbackFS{fs: subFS}

	return http.FileServer(http.FS(fallbackFS))
}
