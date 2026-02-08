package backend

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	gap "github.com/muesli/go-app-paths"
)

// papersCache helps read and write cached papers list and data
type papersCache struct {
	dir string
}

var cache papersCache

func init() {
	scope := gap.NewScope(gap.User, "conf-browser")

	dir, err := scope.CacheDir()
	if err != nil {
		log.Fatalf("Failed to get cache directory: %v", err)
	}

	cache = papersCache{dir: dir}
}

// GetPapers retrieves the list of papers for a given conference and year from the cache.
func (pc *papersCache) GetPapers(conf string, year string) ([]PaperMetadata, error) {
	p := filepath.Join(pc.dir, conf, year, "papers.jsonl")
	return readJSONL[PaperMetadata](p)
}

// WritePapers writes the list of papers for a given conference and year to the cache.
func (pc *papersCache) WritePapers(conf string, year string, papers []PaperMetadata) error {
	dir := filepath.Join(pc.dir, conf, year)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}
	p := filepath.Join(dir, "papers.jsonl")
	return writeJSONL(p, papers)
}

// GetPaper retrieves the details of a single paper from the cache.
func (pc *papersCache) GetPaper(conf string, year string, id string) (Paper, error) {
	p := filepath.Join(pc.dir, conf, year, "papers", id+".json")
	return readJSON[Paper](p)
}

// WritePaper writes the details of a single paper to the cache.
func (pc *papersCache) WritePaper(conf string, year string, id string, paper Paper) error {
	dir := filepath.Join(pc.dir, conf, year, "papers")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}
	p := filepath.Join(dir, id+".json")
	return writeJSON(p, paper)
}

// ************************************************************************* //
//                                  JSON HELPERS                             //
// ************************************************************************* //

// readJSON reads a JSON file into a single object
func readJSON[T any](path string) (T, error) {
	var obj T

	f, err := os.Open(path)
	if err != nil {
		return obj, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&obj); err != nil {
		return obj, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return obj, nil
}

// writeJSON writes a single object to a JSON file, in a pretty-printed format
func writeJSON[T any](path string, obj T) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(obj); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return nil
}

// ************************************************************************* //
//                                 JSONL HELPERS                             //
// ************************************************************************* //

// readJSONL reads a JSONL file into a slice of objects.
func readJSONL[T any](path string) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	objects := make([]T, 0)

	for scanner.Scan() {
		var obj T
		if err := json.Unmarshal(scanner.Bytes(), &obj); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSONL line: %w", err)
		}
		objects = append(objects, obj)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return objects, nil
}

// writeJSONL writes a slice of objects to a JSONL file.
func writeJSONL[T any](path string, objects []T) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, obj := range objects {
		data, err := json.Marshal(obj)
		if err != nil {
			return fmt.Errorf("failed to marshal JSONL object: %w", err)
		}
		if _, err := w.Write(data); err != nil {
			return err
		}
		if err := w.WriteByte('\n'); err != nil {
			return err
		}
	}
	return w.Flush()
}
