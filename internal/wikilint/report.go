package wikilint

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Issue is a single problem found by Lint.
type Issue struct {
	Path    string
	Message string
}

// Report collects issues found while linting a wiki directory.
type Report struct {
	Issues []Issue
}

// Add appends an issue to the report.
func (r *Report) Add(path, msg string) {
	r.Issues = append(r.Issues, Issue{Path: path, Message: msg})
}

// OK reports whether the wiki passed validation.
func (r *Report) OK() bool {
	return len(r.Issues) == 0
}

// Sort puts issues into a deterministic order by path, then message.
func (r *Report) Sort() {
	sort.SliceStable(r.Issues, func(i, j int) bool {
		if r.Issues[i].Path != r.Issues[j].Path {
			return r.Issues[i].Path < r.Issues[j].Path
		}
		return r.Issues[i].Message < r.Issues[j].Message
	})
}

// Lint walks wikiDir, parses every .md file, and runs all validations.
// It returns the collected Report. An error is only returned for I/O or
// configuration problems (e.g. wikiDir does not exist).
func Lint(wikiDir string) (*Report, error) {
	info, err := os.Stat(wikiDir)
	if err != nil {
		return nil, fmt.Errorf("stat wiki dir: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("wiki path %q is not a directory", wikiDir)
	}

	pages, err := loadPages(wikiDir)
	if err != nil {
		return nil, err
	}

	report := &Report{}
	graph := NewGraph(pages)

	for _, rel := range sortedPageKeys(pages) {
		page := pages[rel]
		if msg := checkH1(page); msg != "" {
			report.Add(page.RepoPath, msg)
		}
		if msg := checkRequiredSources(page); msg != "" {
			report.Add(page.RepoPath, msg)
		}
		for _, msg := range checkEmptySections(page) {
			report.Add(page.RepoPath, msg)
		}
		checkLinks(page, graph, report)
	}

	for slug, rels := range findDuplicates(pages) {
		for _, rel := range rels {
			others := otherPaths(rel, rels)
			page := pages[rel]
			msg := fmt.Sprintf("duplicate title slug %q also at: %s", slug, strings.Join(others, ", "))
			report.Add(page.RepoPath, msg)
		}
	}

	const entry = "index.md"
	if _, ok := pages[entry]; ok {
		reachable := graph.Reachable(entry)
		for _, rel := range sortedPageKeys(pages) {
			if rel == entry {
				continue
			}
			if !reachable[rel] {
				report.Add(pages[rel].RepoPath, "orphan page not reachable from index.md")
			}
		}
	} else {
		report.Add(filepath.ToSlash(filepath.Join(wikiDir, entry)), "missing wiki entry point index.md")
	}

	report.Sort()
	return report, nil
}

// loadPages walks wikiDir and parses all .md files into Page values keyed by
// their wiki-root-relative slash-separated path (e.g. "entities/foo.md").
func loadPages(wikiDir string) (map[string]*Page, error) {
	pages := map[string]*Page{}
	err := filepath.WalkDir(wikiDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(d.Name()), ".md") {
			return nil
		}
		data, rerr := os.ReadFile(path)
		if rerr != nil {
			return rerr
		}
		rel, rerr := filepath.Rel(wikiDir, path)
		if rerr != nil {
			return rerr
		}
		relSlash := filepath.ToSlash(rel)
		repoPath := filepath.ToSlash(filepath.Join(wikiDir, rel))
		pages[relSlash] = ParsePage(repoPath, relSlash, string(data))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func sortedPageKeys(pages map[string]*Page) []string {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func otherPaths(self string, all []string) []string {
	out := make([]string, 0, len(all)-1)
	for _, p := range all {
		if p != self {
			out = append(out, p)
		}
	}
	return out
}

func checkLinks(page *Page, graph *Graph, report *Report) {
	for _, l := range page.Links {
		switch classifyLink(l.Target) {
		case linkMarkdown:
			resolved, ok := resolveMarkdownLink(page.RelPath, l.Target)
			if !ok {
				report.Add(page.RepoPath, fmt.Sprintf("broken markdown link %q", l.Target))
				continue
			}
			if _, exists := graph.Pages[resolved]; !exists {
				report.Add(page.RepoPath, fmt.Sprintf("broken markdown link %q", l.Target))
			}
		default:
			// external, anchor, or non-markdown links are ignored
		}
	}
	for _, w := range page.Wikilinks {
		matches := graph.ResolveWikilink(w.Target)
		switch len(matches) {
		case 0:
			report.Add(page.RepoPath, fmt.Sprintf("broken wikilink [[%s]]", w.Target))
		case 1:
			// ok
		default:
			report.Add(page.RepoPath, fmt.Sprintf("ambiguous wikilink [[%s]] matches: %s", w.Target, strings.Join(matches, ", ")))
		}
	}
}
