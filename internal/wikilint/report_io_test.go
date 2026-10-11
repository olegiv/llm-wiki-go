package wikilint

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLint_SkipsSymlinkedMarkdownFile(t *testing.T) {
	wikiDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(wikiDir, "index.md"), []byte("# Index\n\n## Sources\n\n- test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(outside, []byte("## no heading\n"), 0o644); err != nil {
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
