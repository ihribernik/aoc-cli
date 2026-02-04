package solutions_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

func TestRegisterYear(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r       *registry.Registry
		year    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"not found year", registry.NewRegistry(), 2099, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := solutions.RegisterYear(tt.r, tt.year)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("RegisterYear() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("RegisterYear() succeeded unexpectedly")
			}
		})
	}
}
