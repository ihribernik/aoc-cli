package solutions_test

import (
	"strings"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

func TestRegisterYear(t *testing.T) {
	t.Run("registers supported year", func(t *testing.T) {
		r := registry.NewRegistry()
		if err := solutions.RegisterYear(r, 2015); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		solver, ok := r.GetSolver(2015, 1)
		if !ok || solver == nil {
			t.Fatalf("expected solver for 2015 day 1 after registration")
		}
	})

	t.Run("fails for unsupported year", func(t *testing.T) {
		r := registry.NewRegistry()
		err := solutions.RegisterYear(r, 2099)
		if err == nil {
			t.Fatalf("expected error for unsupported year")
		}
		if !strings.Contains(err.Error(), "unsupported year: 2099") {
			t.Fatalf("unexpected error message: %v", err)
		}
	})

	t.Run("fails for nil registry", func(t *testing.T) {
		err := solutions.RegisterYear(nil, 2015)
		if err == nil {
			t.Fatalf("expected error for nil registry")
		}
		if !strings.Contains(err.Error(), "nil registry for year 2015") {
			t.Fatalf("unexpected error message: %v", err)
		}
	})

	t.Run("wraps year registration errors", func(t *testing.T) {
		r := registry.NewRegistry()
		if err := solutions.RegisterYear(r, 2015); err != nil {
			t.Fatalf("unexpected setup error: %v", err)
		}

		err := solutions.RegisterYear(r, 2015)
		if err == nil {
			t.Fatalf("expected duplicate registration error")
		}
		if !strings.Contains(err.Error(), "register year 2015") {
			t.Fatalf("expected year context in error, got: %v", err)
		}
		if !strings.Contains(err.Error(), "solver already registered") {
			t.Fatalf("expected wrapped cause in error, got: %v", err)
		}
	})
}
