package solutions

import (
	"fmt"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions/y2015"
)

// RegisterYear registers all solvers for a specific year.
// It returns an error if the year is unsupported or if any
// solver registration fails.
func RegisterYear(r *registry.Registry, year int) error {
	if r == nil {
		return fmt.Errorf("nil registry for the year %d", year)
	}

	switch year {
	case 2015:
		return y2015.Register(r)
	default:
		return fmt.Errorf("unsupported year: %d", year)
	}
}
