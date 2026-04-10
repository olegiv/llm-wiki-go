package wikilint

import "testing"

func TestCheckH1(t *testing.T) {
	cases := []struct {
		name    string
		count   int
		wantMsg string
	}{
		{"no H1", 0, "missing H1 title"},
		{"one H1", 1, ""},
		{"two H1", 2, "multiple H1 titles"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := checkH1(&Page{H1Count: tc.count})
			if got != tc.wantMsg {
				t.Errorf("got %q, want %q", got, tc.wantMsg)
			}
		})
	}
}

func TestCheckRequiredSources(t *testing.T) {
	cases := []struct {
		name    string
		rel     string
		hasSrc  bool
		wantMsg string
	}{
		{"index exempt", "index.md", false, ""},
		{"log exempt", "log.md", false, ""},
		{"substantive missing", "entities/foo.md", false, "missing required `## Sources` section"},
		{"substantive present", "entities/foo.md", true, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := &Page{RelPath: tc.rel, Sections: map[string]string{}}
			if tc.hasSrc {
				p.Sections["Sources"] = "- a source"
			}
			got := checkRequiredSources(p)
			if got != tc.wantMsg {
				t.Errorf("got %q, want %q", got, tc.wantMsg)
			}
		})
	}
}

func TestCheckEmptySections(t *testing.T) {
	cases := []struct {
		name     string
		sections map[string]string
		want     []string
	}{
		{
			name:     "all filled",
			sections: map[string]string{"Sources": "- a", "Contradictions": "none", "Open questions": "q"},
			want:     nil,
		},
		{
			name:     "empty sources",
			sections: map[string]string{"Sources": ""},
			want:     []string{"`## Sources` section is empty"},
		},
		{
			name:     "empty contradictions and open questions",
			sections: map[string]string{"Sources": "- a", "Contradictions": "", "Open questions": ""},
			want: []string{
				"`## Contradictions` section is empty",
				"`## Open questions` section is empty",
			},
		},
		{
			name:     "only untracked empty sections",
			sections: map[string]string{"Summary": ""},
			want:     nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := checkEmptySections(&Page{Sections: tc.sections})
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i, msg := range tc.want {
				if got[i] != msg {
					t.Errorf("got[%d] = %q, want %q", i, got[i], msg)
				}
			}
		})
	}
}
