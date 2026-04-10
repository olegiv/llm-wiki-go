package wikilint

import (
	"reflect"
	"testing"
)

func TestClassifyLink(t *testing.T) {
	cases := []struct {
		target string
		want   linkKind
	}{
		{"http://example.com", linkExternal},
		{"https://example.com/path", linkExternal},
		{"HTTPS://Example.com", linkExternal},
		{"mailto:a@b.com", linkExternal},
		{"tel:+15551234", linkExternal},
		{"ftp://files", linkExternal},
		{"#section", linkAnchor},
		{"", linkAnchor},
		{"foo.md", linkMarkdown},
		{"../topics/bar.md", linkMarkdown},
		{"foo.md#anchor", linkMarkdown},
		{"image.png", linkNonMarkdown},
		{"../data/file.json", linkNonMarkdown},
	}
	for _, tc := range cases {
		t.Run(tc.target, func(t *testing.T) {
			if got := classifyLink(tc.target); got != tc.want {
				t.Errorf("classifyLink(%q) = %d, want %d", tc.target, got, tc.want)
			}
		})
	}
}

func TestExtractLinks_IgnoresImages(t *testing.T) {
	content := "See ![alt text](image.png) and [link](page.md).\n"
	links, _ := extractLinks(content)
	if len(links) != 1 {
		t.Fatalf("want 1 link, got %d: %+v", len(links), links)
	}
	if links[0].Target != "page.md" {
		t.Errorf("target = %q, want %q", links[0].Target, "page.md")
	}
}

func TestExtractLinks_IgnoresCodeFences(t *testing.T) {
	content := "Outside [real](real.md).\n\n```\n[hidden](hidden.md)\n```\n\nInline `[also hidden](also.md)` text.\n"
	links, _ := extractLinks(content)
	if len(links) != 1 {
		t.Fatalf("want 1 link, got %d: %+v", len(links), links)
	}
	if links[0].Target != "real.md" {
		t.Errorf("target = %q, want %q", links[0].Target, "real.md")
	}
}

func TestExtractLinks_Wikilinks(t *testing.T) {
	content := "See [[foo]] and [[topics/bar]] and [[baz|Display]].\n"
	_, wikis := extractLinks(content)
	wantTargets := []string{"foo", "topics/bar", "baz|Display"}
	if len(wikis) != len(wantTargets) {
		t.Fatalf("want %d wikilinks, got %d: %+v", len(wantTargets), len(wikis), wikis)
	}
	for i, w := range wikis {
		if w.Target != wantTargets[i] {
			t.Errorf("wikis[%d] = %q, want %q", i, w.Target, wantTargets[i])
		}
	}
}

func TestResolveMarkdownLink(t *testing.T) {
	cases := []struct {
		name   string
		from   string
		target string
		want   string
		wantOK bool
	}{
		{"same dir", "entities/foo.md", "bar.md", "entities/bar.md", true},
		{"parent dir", "entities/foo.md", "../topics/baz.md", "topics/baz.md", true},
		{"root from subdir", "entities/foo.md", "../index.md", "index.md", true},
		{"strip anchor", "entities/foo.md", "bar.md#section", "entities/bar.md", true},
		{"escape above root", "foo.md", "../outside.md", "", false},
		{"empty target", "foo.md", "", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := resolveMarkdownLink(tc.from, tc.target)
			if ok != tc.wantOK || got != tc.want {
				t.Errorf("resolveMarkdownLink(%q, %q) = (%q, %v), want (%q, %v)",
					tc.from, tc.target, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

func TestExtractLinks_WithTitleAttribute(t *testing.T) {
	content := `See [foo](foo.md "A title") and [bar](bar.md).`
	links, _ := extractLinks(content)
	got := []string{}
	for _, l := range links {
		got = append(got, l.Target)
	}
	want := []string{"foo.md", "bar.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("targets = %v, want %v", got, want)
	}
}
