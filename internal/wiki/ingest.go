// Package wiki contains small helpers used by the Claude Code skills and
// agents that operate on the wiki: ingesting sources, answering from the
// wiki, and reconciling contradictions. The helpers deliberately stay
// minimal — they are scaffold hooks, not a full pipeline.
package wiki

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// maxFileBytes is the largest file the wiki helpers will read.
const maxFileBytes = 10 << 20

// ListRawFiles walks rawDir read-only and returns the slash-separated
// relative paths of regular files in sorted order. It never modifies the
// raw directory.
func ListRawFiles(rawDir string) ([]string, error) {
	resolved, err := filepath.EvalSymlinks(rawDir)
	if err != nil {
		return nil, fmt.Errorf("resolve raw dir: %w", err)
	}
	var paths []string
	err = filepath.WalkDir(resolved, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() == ".gitkeep" {
			return nil
		}
		rel, rerr := filepath.Rel(resolved, path)
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
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil, fmt.Errorf("%s: symlink not allowed", path)
	}
	data, ok, err := readRegularFileLimited(path, maxFileBytes)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("%s: not a regular file", path)
	}
	return data, nil
}

func readRegularFileLimited(path string, maxBytes int64) ([]byte, bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, false, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil, false, nil
	}
	if !info.Mode().IsRegular() {
		return nil, false, nil
	}
	if info.Size() > maxBytes {
		return nil, false, fmt.Errorf("%s: file too large (%d bytes, max %d)", path, info.Size(), maxBytes)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()

	limited := io.LimitReader(f, maxBytes+1)
	data, err := io.ReadAll(limited)
	if err != nil {
		return nil, false, err
	}
	if int64(len(data)) > maxBytes {
		return nil, false, fmt.Errorf("%s: file too large (over %d bytes)", path, maxBytes)
	}
	return data, true, nil
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
