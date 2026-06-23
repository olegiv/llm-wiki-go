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

func TestListRawFiles_Symlink(t *testing.T) {
	real := t.TempDir()
	if err := os.WriteFile(filepath.Join(real, "note.txt"), []byte("content"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(t.TempDir(), "linked")
	if err := os.Symlink(real, link); err != nil {
		t.Fatal(err)
	}
	got, err := ListRawFiles(link)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"note.txt"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadRawFile_TooLarge(t *testing.T) {
	big := filepath.Join(t.TempDir(), "huge.bin")
	if err := os.WriteFile(big, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Truncate(big, 11<<20); err != nil {
		t.Fatal(err)
	}
	_, err := ReadRawFile(big)
	if err == nil {
		t.Fatal("expected error for oversized file")
	}
	if !strings.Contains(err.Error(), "file too large") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestReadRawFile_Symlink(t *testing.T) {
	target := filepath.Join(t.TempDir(), "target.txt")
	if err := os.WriteFile(target, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(t.TempDir(), "linked.txt")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	_, err := ReadRawFile(link)
	if err == nil {
		t.Fatal("expected error for symlink")
	}
	if !strings.Contains(err.Error(), "symlink") {
		t.Errorf("unexpected error: %v", err)
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
