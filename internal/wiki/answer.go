package wiki

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

// SearchWiki walks wikiDir and returns the slash-separated relative paths of
// Markdown pages whose content contains every supplied term (case-insensitive).
// Results are sorted for determinism. An empty terms slice returns all pages.
func SearchWiki(wikiDir string, terms []string) ([]string, error) {
	resolved, err := filepath.EvalSymlinks(wikiDir)
	if err != nil {
		return nil, fmt.Errorf("resolve wiki dir: %w", err)
	}
	var matches []string
	err = filepath.WalkDir(resolved, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(d.Name()), ".md") {
			return nil
		}
		data, ok, rerr := readRegularFileLimited(path, maxFileBytes)
		if rerr != nil {
			return rerr
		}
		if !ok {
			return nil
		}
		lower := strings.ToLower(string(data))
		for _, term := range terms {
			if !strings.Contains(lower, strings.ToLower(term)) {
				return nil
			}
		}
		rel, rerr := filepath.Rel(resolved, path)
		if rerr != nil {
			return rerr
		}
		matches = append(matches, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(matches)
	return matches, nil
}
