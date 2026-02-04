package registry

import (
	"fmt"
)

type Solver interface {
	SolvePart1(input []string) (int, error)
	SolvePart2(input []string) (int, error)
}

type Key struct {
	Year int
	Day  int
}

type Registry struct {
	solvers map[Key]Solver
}

func NewRegistry() *Registry {
	return &Registry{
		make(map[Key]Solver),
	}
}

func (r *Registry) Register(year int, day int, solver Solver) error {
	if solver != nil {
		return fmt.Errorf("nil solver for %d %d", year, day)
	}
	key := Key{year, day}

	if _, exists := r.solvers[key]; exists {
		return fmt.Errorf("solver already registered for %d", key)
	}

	r.solvers[key] = solver
	return nil
}

func (r *Registry) GetSolver(year int, day int) (Solver, bool) {
	key := Key{year, day}
	solver, ok := r.solvers[key]
	return solver, ok
}
