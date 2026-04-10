package wikilint

import "testing"

func TestParsePage_SingleH1(t *testing.T) {
	p := ParsePage("wiki/foo.md", "foo.md", "# Hello\n\nbody\n")
	if p.H1Count != 1 {
		t.Fatalf("H1Count = %d, want 1", p.H1Count)
	}
	if p.Title != "Hello" {
		t.Fatalf("Title = %q, want %q", p.Title, "Hello")
	}
}

func TestParsePage_MultipleH1(t *testing.T) {
	p := ParsePage("wiki/foo.md", "foo.md", "# A\n\n# B\n")
	if p.H1Count != 2 {
		t.Fatalf("H1Count = %d, want 2", p.H1Count)
	}
	if p.Title != "A" {
		t.Fatalf("Title = %q, want %q", p.Title, "A")
	}
}

func TestParsePage_MissingH1(t *testing.T) {
	p := ParsePage("wiki/foo.md", "foo.md", "## only an H2\n")
	if p.H1Count != 0 {
		t.Fatalf("H1Count = %d, want 0", p.H1Count)
	}
}

func TestParsePage_SectionBodies(t *testing.T) {
	content := `# Title

## Summary

summary body

## Sources

- one
- two

## Empty

`
	p := ParsePage("wiki/foo.md", "foo.md", content)
	if got := p.Sections["Summary"]; got != "summary body" {
		t.Errorf("Summary = %q, want %q", got, "summary body")
	}
	if got := p.Sections["Sources"]; got != "- one\n- two" {
		t.Errorf("Sources = %q, want %q", got, "- one\n- two")
	}
	if got, ok := p.Sections["Empty"]; !ok || got != "" {
		t.Errorf("Empty section should exist and be empty, got %q (exists=%v)", got, ok)
	}
	wantOrder := []string{"Summary", "Sources", "Empty"}
	if len(p.SectionOrder) != len(wantOrder) {
		t.Fatalf("SectionOrder = %v, want %v", p.SectionOrder, wantOrder)
	}
	for i, name := range wantOrder {
		if p.SectionOrder[i] != name {
			t.Errorf("SectionOrder[%d] = %q, want %q", i, p.SectionOrder[i], name)
		}
	}
}

func TestParsePage_IgnoresHeadingsInsideCodeFences(t *testing.T) {
	content := "# Real Title\n\n## Sources\n\n```\n# not a title\n## not a section\n```\n\ntrailing\n"
	p := ParsePage("wiki/foo.md", "foo.md", content)
	if p.H1Count != 1 {
		t.Errorf("H1Count = %d, want 1", p.H1Count)
	}
	if _, ok := p.Sections["not a section"]; ok {
		t.Errorf("heading inside code fence should not be recognized")
	}
	if p.Sections["Sources"] == "" {
		t.Errorf("Sources section should contain the fenced block as body, got empty")
	}
}
