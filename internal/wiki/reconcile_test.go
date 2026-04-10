package wiki

import (
	"strings"
	"testing"
)

func TestFormatContradiction(t *testing.T) {
	got := FormatContradiction("Disputed start year", []ContradictionEntry{
		{Source: "sources/alice.md", Claim: "Founded in 1999"},
		{Source: "sources/bob.md", Claim: "Founded in 2001"},
	})
	musts := []string{
		"### Disputed start year",
		"- sources/alice.md: Founded in 1999",
		"- sources/bob.md: Founded in 2001",
	}
	for _, m := range musts {
		if !strings.Contains(got, m) {
			t.Errorf("output missing %q in:\n%s", m, got)
		}
	}
}

func TestFormatContradiction_DefaultTopic(t *testing.T) {
	got := FormatContradiction("", []ContradictionEntry{
		{Source: "sources/a.md", Claim: "X"},
	})
	if !strings.Contains(got, "### Unresolved contradiction") {
		t.Errorf("expected default heading, got:\n%s", got)
	}
}
