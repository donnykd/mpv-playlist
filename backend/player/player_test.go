package player

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	t.Parallel()

	t.Run("Test relative to abs path works", func(t *testing.T) {
		t.Parallel()

		result, err := normalizePath("../test.mp4")
		if err != nil {
			t.Fatal(err)
		}
		if !filepath.IsAbs(result) {
			t.Errorf("expected absolute path, got %v", result)
		}
	})

	t.Run("Test if path is cleaned", func(t *testing.T) {
		t.Parallel()

		result, err := normalizePath("../denmark/./test.mkv")
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(result, "/./") {
			t.Errorf("expected path cleaned, got %v", result)
		}
	})

	t.Run("Test empty string", func(t *testing.T) {
		t.Parallel()

		if _, err := normalizePath(""); err == nil {
			t.Fatal("Expected error")
		}
	})
}
