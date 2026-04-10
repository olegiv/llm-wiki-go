// Package wiki contains small helpers used by the Claude Code skills and
// agents that operate on the wiki: ingesting sources, answering from the
// wiki, and reconciling contradictions. The helpers deliberately stay
// minimal — they are scaffold hooks, not a full pipeline.
package wiki

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ListRawFiles walks rawDir read-only and returns the slash-separated
// relative paths of regular files in sorted order. It never modifies the
// raw directory.
func ListRawFiles(rawDir string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(rawDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() == ".gitkeep" {
			return nil
		}
		rel, rerr := filepath.Rel(rawDir, path)
		if rerr != nil {
			return rerr
		}
		paths = append(paths, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	return paths, nil
}

// ReadRawFile returns the contents of a raw source file. It is read-only.
func ReadRawFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// SourcePageTemplate produces a bootstrap wiki page body for an ingested
// raw source. The output contains an H1 title, a `## Summary` placeholder,
// and a `## Sources` section that records the origin.
func SourcePageTemplate(title, origin string) string {
	var b strings.Builder
	b.WriteString("# ")
	b.WriteString(title)
	b.WriteString("\n\n")
	b.WriteString("## Summary\n\n")
	b.WriteString("TODO: summarize this source after ingesting it.\n\n")
	b.WriteString("## Sources\n\n")
	b.WriteString("- Origin: `")
	b.WriteString(origin)
	b.WriteString("`\n")
	return b.String()
}
