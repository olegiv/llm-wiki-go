package wikilint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLint_SkipsSymlinkedMarkdownFile(t *testing.T) {
	wikiDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(wikiDir, "index.md"), []byte("# Index\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// The target would fail linting and exceeds maxPageBytes; a symlink's
	// DirEntry size must not let it slip past the size guard.
	outside := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(outside, []byte("## no heading\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Truncate(outside, 11<<20); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(wikiDir, "linked.md")); err != nil {
		t.Fatal(err)
	}

	report, err := Lint(wikiDir)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK() {
		t.Fatalf("unexpected issues: %#v", report.Issues)
	}
}

func TestLint_OversizedPage(t *testing.T) {
	wikiDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(wikiDir, "index.md"), []byte("# Index\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	big := filepath.Join(wikiDir, "huge.md")
	if err := os.WriteFile(big, []byte("# Huge\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Truncate(big, 11<<20); err != nil {
		t.Fatal(err)
	}
	_, err := Lint(wikiDir)
	if err == nil {
		t.Fatal("expected error for oversized page")
	}
	if !strings.Contains(err.Error(), "file too large") {
		t.Errorf("unexpected error: %v", err)
	}
}
