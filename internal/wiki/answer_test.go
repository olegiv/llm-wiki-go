package wiki

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

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
