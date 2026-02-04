package registry_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

func TestRegistry_Register(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		year    int
		day     int
		solver  registry.Solver
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := registry.NewRegistry()
			gotErr := r.Register(tt.year, tt.day, tt.solver)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Register() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Register() succeeded unexpectedly")
			}
		})
	}
}

func TestRegistry_GetSolver(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		year  int
		day   int
		want  registry.Solver
		want2 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := registry.NewRegistry()
			got, got2 := r.GetSolver(tt.year, tt.day)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetSolver() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("GetSolver() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestNewRegistry(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want *registry.Registry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.NewRegistry()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}
