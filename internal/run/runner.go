package run

import "github.com/ihribernik/aoc-cli/internal/registry"

type NewRegistryFunc func() *registry.Registry
type RegisterYearFunc func(*registry.Registry, int) error
type GetInputFunc func(int, int) ([]string, error)

type Result struct {
	Part1 int
	Part2 int
}

type Runner struct {
	newRegistry  NewRegistryFunc
	registerYear RegisterYearFunc
	getInput     GetInputFunc
}

func NewRunner(newRegistry NewRegistryFunc, registerYear RegisterYearFunc, getInput GetInputFunc) *Runner {
	return &Runner{
		newRegistry:  newRegistry,
		registerYear: registerYear,
		getInput:     getInput,
	}
}

func (r *Runner) Execute(year int, day int) (Result, error) {
	if r == nil {
		return Result{}, ErrNilRunner
	}
	if r.newRegistry == nil || r.registerYear == nil || r.getInput == nil {
		return Result{}, ErrRunnerNotConfigured
	}

	reg := r.newRegistry()
	if reg == nil {
		return Result{}, ErrNilRegistry
	}

	if err := r.registerYear(reg, year); err != nil {
		return Result{}, &ErrRegisterYear{Year: year, Err: err}
	}

	solver, ok := reg.GetSolver(year, day)
	if !ok {
		return Result{}, &ErrSolverNotFound{Year: year, Day: day}
	}

	input, err := r.getInput(year, day)
	if err != nil {
		return Result{}, &ErrGetInput{Year: year, Day: day, Err: err}
	}

	part1, err := solver.SolvePart1(input)
	if err != nil {
		return Result{}, &ErrSolvePart{Part: 1, Err: err}
	}

	part2, err := solver.SolvePart2(input)
	if err != nil {
		return Result{}, &ErrSolvePart{Part: 2, Err: err}
	}

	return Result{
		Part1: part1,
		Part2: part2,
	}, nil
}
