package inputs_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/inputs"
)

func TestGetInput(t *testing.T) {
	t.Run("returns error when file does not exist", func(t *testing.T) {
		tempDir := t.TempDir()
		chdirForTest(t, tempDir)

		got, err := inputs.GetInput(2099, 1)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if len(got) != 0 {
			t.Fatalf("expected empty slice on error, got %v", got)
		}
	})

	t.Run("normalizes line endings and trims trailing newline", func(t *testing.T) {
		tempDir := t.TempDir()
		chdirForTest(t, tempDir)
		mustWriteInput(t, tempDir, 2026, 1, "line1\r\nline2\r\n")

		got, err := inputs.GetInput(2026, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := []string{"line1", "line2"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("GetInput() = %v, want %v", got, want)
		}
	})

	t.Run("preserves internal blank lines", func(t *testing.T) {
		tempDir := t.TempDir()
		chdirForTest(t, tempDir)
		mustWriteInput(t, tempDir, 2026, 2, "a\n\nb\n")

		got, err := inputs.GetInput(2026, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := []string{"a", "", "b"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("GetInput() = %v, want %v", got, want)
		}
	})

	t.Run("returns empty slice for empty file", func(t *testing.T) {
		tempDir := t.TempDir()
		chdirForTest(t, tempDir)
		mustWriteInput(t, tempDir, 2026, 3, "")

		got, err := inputs.GetInput(2026, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got) != 0 {
			t.Fatalf("expected empty slice, got %v", got)
		}
	})
}

func chdirForTest(t *testing.T, dir string) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
}

func mustWriteInput(t *testing.T, baseDir string, year int, day int, content string) {
	t.Helper()
	dir := filepath.Join(baseDir, "input", fmt.Sprintf("%d", year))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir input dir: %v", err)
	}
	path := filepath.Join(dir, fmt.Sprintf("day%02d.txt", day))
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write input file: %v", err)
	}
}
