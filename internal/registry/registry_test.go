package registry_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

type testSolver struct{}

func (s testSolver) SolvePart1(input []string) (int, error) { return len(input), nil }
func (s testSolver) SolvePart2(input []string) (int, error) { return len(input), nil }

func TestNewRegistry(t *testing.T) {
	r := registry.NewRegistry()
	if r == nil {
		t.Fatalf("expected non-nil registry")
	}

	if _, ok := r.GetSolver(2015, 1); ok {
		t.Fatalf("expected empty registry")
	}
}

func TestRegistryRegisterAndGetSolver(t *testing.T) {
	r := registry.NewRegistry()
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}

	got, ok := r.GetSolver(2015, 1)
	if !ok {
		t.Fatalf("expected solver to exist")
	}
	if got == nil {
		t.Fatalf("expected non-nil solver")
	}
}

func TestRegistryRegisterNilSolver(t *testing.T) {
	r := registry.NewRegistry()

	if err := r.Register(2015, 1, nil); err == nil {
		t.Fatalf("expected error for nil solver")
	}
}

func TestRegistryRegisterDuplicate(t *testing.T) {
	r := registry.NewRegistry()
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}
	if err := r.Register(2015, 1, s); err == nil {
		t.Fatalf("expected duplicate register error")
	}
}

func TestRegistryRegisterOnNilReceiver(t *testing.T) {
	var r *registry.Registry

	if err := r.Register(2015, 1, testSolver{}); err == nil {
		t.Fatalf("expected error for nil receiver")
	}
}

func TestRegistryGetSolverOnNilReceiver(t *testing.T) {
	var r *registry.Registry

	if got, ok := r.GetSolver(2015, 1); ok || got != nil {
		t.Fatalf("expected nil,false for nil receiver, got %v,%v", got, ok)
	}
}

func TestRegistryRegisterInitializesNilMap(t *testing.T) {
	r := &registry.Registry{}
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}

	if got, ok := r.GetSolver(2015, 1); !ok || got == nil {
		t.Fatalf("expected solver after lazy map initialization")
	}
}

func TestRegistryGetSolverOnNilMap(t *testing.T) {
	r := &registry.Registry{}

	if got, ok := r.GetSolver(2015, 1); ok || got != nil {
		t.Fatalf("expected nil,false for nil map, got %v,%v", got, ok)
	}
}
