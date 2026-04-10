package wikilint

import (
	"reflect"
	"testing"
)

func TestNormalizeSlug(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Sample Entity", "sample-entity"},
		{"  Whitespace   Padding  ", "whitespace-padding"},
		{"Punctuation!! & Stuff?", "punctuation-stuff"},
		{"Mixed CASE Name", "mixed-case-name"},
		{"___under___scores", "under-scores"},
		{"", ""},
		{"!!!", ""},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got := normalizeSlug(tc.in)
			if got != tc.want {
				t.Errorf("normalizeSlug(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestFindDuplicates_NoDuplicates(t *testing.T) {
	pages := map[string]*Page{
		"a.md": {Title: "Alpha"},
		"b.md": {Title: "Beta"},
	}
	if got := findDuplicates(pages); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestFindDuplicates_TwoPagesSameSlug(t *testing.T) {
	pages := map[string]*Page{
		"entities/foo.md": {Title: "Foo Bar"},
		"topics/foo.md":   {Title: "foo bar"},
		"entities/baz.md": {Title: "Baz"},
	}
	got := findDuplicates(pages)
	want := map[string][]string{
		"foo-bar": {"entities/foo.md", "topics/foo.md"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindDuplicates_IgnoresEmptyTitles(t *testing.T) {
	pages := map[string]*Page{
		"a.md": {Title: ""},
		"b.md": {Title: ""},
	}
	if got := findDuplicates(pages); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}
