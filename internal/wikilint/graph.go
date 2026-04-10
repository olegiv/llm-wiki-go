package wikilint

import (
	"sort"
	"strings"
)

// Graph indexes all pages in a wiki so reachability and wikilink resolution
// can be evaluated.
type Graph struct {
	// Pages maps wiki-root-relative slash paths to parsed pages.
	Pages map[string]*Page
	// Stems maps a lowercased filename stem (e.g. "foo") to the sorted list
	// of pages whose filename matches.
	Stems map[string][]string
}

// NewGraph builds a Graph from a set of parsed pages.
func NewGraph(pages map[string]*Page) *Graph {
	g := &Graph{
		Pages: pages,
		Stems: map[string][]string{},
	}
	for rel := range pages {
		stem := stemOf(rel)
		g.Stems[stem] = append(g.Stems[stem], rel)
	}
	for k := range g.Stems {
		sort.Strings(g.Stems[k])
	}
	return g
}

// stemOf returns the lowercase filename stem of a wiki-root-relative path.
// For "entities/foo.md" it returns "foo".
func stemOf(relPath string) string {
	base := relPath
	if i := strings.LastIndex(base, "/"); i >= 0 {
		base = base[i+1:]
	}
	base = strings.TrimSuffix(base, ".md")
	base = strings.TrimSuffix(base, ".markdown")
	return strings.ToLower(base)
}

// ResolveWikilink returns the wiki-root-relative paths that match the given
// wikilink target. Zero results means the link is broken; more than one means
// it is ambiguous.
func (g *Graph) ResolveWikilink(target string) []string {
	t := strings.TrimSpace(target)
	if i := strings.Index(t, "|"); i >= 0 {
		t = t[:i]
	}
	if i := strings.Index(t, "#"); i >= 0 {
		t = t[:i]
	}
	t = strings.TrimSpace(t)
	if t == "" {
		return nil
	}

	if strings.Contains(t, "/") {
		candidate := t
		lower := strings.ToLower(candidate)
		if !strings.HasSuffix(lower, ".md") && !strings.HasSuffix(lower, ".markdown") {
			candidate += ".md"
		}
		if _, ok := g.Pages[candidate]; ok {
			return []string{candidate}
		}
		return nil
	}

	stem := strings.ToLower(strings.TrimSuffix(strings.TrimSuffix(t, ".md"), ".markdown"))
	if matches, ok := g.Stems[stem]; ok {
		out := make([]string, len(matches))
		copy(out, matches)
		return out
	}
	return nil
}

// Reachable returns the set of wiki-root-relative paths reachable from entry
// via Markdown links and wikilinks.
func (g *Graph) Reachable(entry string) map[string]bool {
	visited := map[string]bool{}
	if _, ok := g.Pages[entry]; !ok {
		return visited
	}
	queue := []string{entry}
	visited[entry] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		page := g.Pages[cur]
		if page == nil {
			continue
		}
		for _, l := range page.Links {
			if classifyLink(l.Target) != linkMarkdown {
				continue
			}
			resolved, ok := resolveMarkdownLink(cur, l.Target)
			if !ok {
				continue
			}
			if _, exists := g.Pages[resolved]; exists && !visited[resolved] {
				visited[resolved] = true
				queue = append(queue, resolved)
			}
		}
		for _, w := range page.Wikilinks {
			for _, m := range g.ResolveWikilink(w.Target) {
				if !visited[m] {
					visited[m] = true
					queue = append(queue, m)
				}
			}
		}
	}
	return visited
}
