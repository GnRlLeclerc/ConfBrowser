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
func (pc *papersCache) GetPapers(conf string, year string) ([]Paper, error) {
	p := filepath.Join(pc.dir, conf, year, "papers.jsonl")
	return readJSONL[Paper](p)
}

// WritePapers writes the list of papers for a given conference and year to the cache.
func (pc *papersCache) WritePapers(conf string, year string, papers []Paper) error {
	dir := filepath.Join(pc.dir, conf, year)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}
	p := filepath.Join(dir, "papers.jsonl")
	return writeJSONL(p, papers)
}

// ************************************************************************* //
//                                 JSONL HELPERS 												     //
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
