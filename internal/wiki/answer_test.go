package wiki

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSearchWiki_Symlink(t *testing.T) {
	real := t.TempDir()
	sub := filepath.Join(real, "entities")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "fox.md"), []byte("# Fox\n\nquick brown fox\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(t.TempDir(), "linked")
	if err := os.Symlink(real, link); err != nil {
		t.Fatal(err)
	}
	got, err := SearchWiki(link, []string{"fox"})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"entities/fox.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSearchWiki_TooLarge(t *testing.T) {
	dir := t.TempDir()
	big := filepath.Join(dir, "huge.md")
	if err := os.WriteFile(big, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Truncate(big, 11<<20); err != nil {
		t.Fatal(err)
	}
	_, err := SearchWiki(dir, nil)
	if err == nil {
		t.Fatal("expected error for oversized file")
	}
	if !strings.Contains(err.Error(), "file too large") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSearchWiki_SkipsSymlinkedMarkdownFile(t *testing.T) {
	dir := t.TempDir()
	external := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(external, []byte("# Outside\n\nfox\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(external, filepath.Join(dir, "linked.md")); err != nil {
		t.Fatal(err)
	}

	got, err := SearchWiki(dir, []string{"fox"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("got %v, want no matches", got)
	}
}

func TestSearchWiki(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"index.md":        "# Index\n\n- [foo](entities/foo.md)\n",
		"entities/foo.md": "# Foo\n\n## Summary\n\nThe quick brown fox jumps over the lazy dog.\n",
		"entities/bar.md": "# Bar\n\n## Summary\n\nCats are also animals.\n",
		"topics/misc.md":  "# Misc\n\n## Summary\n\nfoxes and cats coexist.\n",
		"ignore.txt":      "fox cat",
	}
	for rel, body := range files {
		full := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	cases := []struct {
		name  string
		terms []string
		want  []string
	}{
		{"fox matches two", []string{"fox"}, []string{"entities/foo.md", "topics/misc.md"}},
		{"intersection", []string{"fox", "cat"}, []string{"topics/misc.md"}},
		{"no match", []string{"octopus"}, nil},
		{"case insensitive", []string{"FOX"}, []string{"entities/foo.md", "topics/misc.md"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SearchWiki(dir, tc.terms)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
