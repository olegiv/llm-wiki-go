package wikilint

import (
	"reflect"
	"sort"
	"testing"
)

func mkPage(rel, title string, links []string, wikis []string) *Page {
	p := &Page{
		RepoPath: "wiki/" + rel,
		RelPath:  rel,
		Title:    title,
		H1Count:  1,
		Sections: map[string]string{"Sources": "- placeholder"},
	}
	for _, l := range links {
		p.Links = append(p.Links, Link{Target: l})
	}
	for _, w := range wikis {
		p.Wikilinks = append(p.Wikilinks, Wikilink{Target: w})
	}
	return p
}

func buildTestGraph() *Graph {
	pages := map[string]*Page{
		"index.md":                   mkPage("index.md", "Index", []string{"entities/foo.md", "topics/bar.md"}, nil),
		"entities/foo.md":            mkPage("entities/foo.md", "Foo", nil, []string{"bar"}),
		"topics/bar.md":              mkPage("topics/bar.md", "Bar", nil, nil),
		"sources/baz.md":             mkPage("sources/baz.md", "Baz", nil, nil),
		"entities/duplicate-name.md": mkPage("entities/duplicate-name.md", "Dup", nil, nil),
		"topics/duplicate-name.md":   mkPage("topics/duplicate-name.md", "DupTwin", nil, nil),
	}
	return NewGraph(pages)
}

func TestResolveWikilink_Stem(t *testing.T) {
	g := buildTestGraph()
	got := g.ResolveWikilink("foo")
	want := []string{"entities/foo.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestResolveWikilink_ExplicitPath(t *testing.T) {
	g := buildTestGraph()
	got := g.ResolveWikilink("topics/bar")
	want := []string{"topics/bar.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestResolveWikilink_WithDisplayAndAnchor(t *testing.T) {
	g := buildTestGraph()
	got := g.ResolveWikilink("foo#section|Display")
	want := []string{"entities/foo.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestResolveWikilink_Ambiguous(t *testing.T) {
	g := buildTestGraph()
	got := g.ResolveWikilink("duplicate-name")
	sort.Strings(got)
	want := []string{"entities/duplicate-name.md", "topics/duplicate-name.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestResolveWikilink_BrokenReturnsEmpty(t *testing.T) {
	g := buildTestGraph()
	if got := g.ResolveWikilink("nonexistent"); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
	if got := g.ResolveWikilink("topics/nonexistent"); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestReachable_FromIndex(t *testing.T) {
	g := buildTestGraph()
	got := g.Reachable("index.md")
	mustReach := []string{"index.md", "entities/foo.md", "topics/bar.md"}
	for _, p := range mustReach {
		if !got[p] {
			t.Errorf("%s should be reachable", p)
		}
	}
	// sources/baz.md is only linked from nothing -> orphan
	if got["sources/baz.md"] {
		t.Errorf("sources/baz.md should be orphan, was reached")
	}
}

func TestReachable_FollowsWikilinks(t *testing.T) {
	pages := map[string]*Page{
		"index.md":        mkPage("index.md", "Index", nil, []string{"foo"}),
		"entities/foo.md": mkPage("entities/foo.md", "Foo", nil, []string{"topics/bar"}),
		"topics/bar.md":   mkPage("topics/bar.md", "Bar", nil, nil),
	}
	g := NewGraph(pages)
	got := g.Reachable("index.md")
	for _, p := range []string{"index.md", "entities/foo.md", "topics/bar.md"} {
		if !got[p] {
			t.Errorf("%s should be reachable via wikilinks", p)
		}
	}
}
