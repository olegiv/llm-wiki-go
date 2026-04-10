package wiki

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestListRawFiles(t *testing.T) {
	dir := t.TempDir()
	files := []string{
		"alpha.txt",
		"nested/beta.log",
		"nested/deep/gamma.md",
		".gitkeep",
	}
	for _, f := range files {
		full := filepath.Join(dir, f)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("content"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	got, err := ListRawFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"alpha.txt", "nested/beta.log", "nested/deep/gamma.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestListRawFiles_MissingDir(t *testing.T) {
	_, err := ListRawFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for missing dir")
	}
}

func TestSourcePageTemplate(t *testing.T) {
	got := SourcePageTemplate("Example Source", "raw/example.txt")
	musts := []string{
		"# Example Source",
		"## Summary",
		"## Sources",
		"`raw/example.txt`",
	}
	for _, m := range musts {
		if !strings.Contains(got, m) {
			t.Errorf("template missing %q in:\n%s", m, got)
		}
	}
}
