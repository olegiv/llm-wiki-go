package wikilint

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPages_DoesNotFollowRootSymlink(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	outside := filepath.Join(tmp, "outside")
	if err := os.Mkdir(outside, 0o755); err != nil {
		t.Fatalf("mkdir outside: %v", err)
	}
	outsidePage := filepath.Join(outside, "outside.md")
	if err := os.WriteFile(outsidePage, []byte("# outside\n\n## Summary\n"), 0o644); err != nil {
		t.Fatalf("write outside page: %v", err)
	}

	wikiLink := filepath.Join(tmp, "wiki")
	if err := os.Symlink(outside, wikiLink); err != nil {
		t.Fatalf("symlink wiki: %v", err)
	}

	pages, err := loadPages(wikiLink)
	if err != nil {
		t.Fatalf("loadPages: %v", err)
	}
	if len(pages) != 0 {
		t.Fatalf("got %d pages, want 0", len(pages))
	}
}
