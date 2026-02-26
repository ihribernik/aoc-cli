package run

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

type stubSolver struct {
	part1    int
	part2    int
	part1Err error
	part2Err error
}

func (s stubSolver) SolvePart1(input []string) (int, error) {
	if s.part1Err != nil {
		return 0, s.part1Err
	}
	return s.part1, nil
}

func (s stubSolver) SolvePart2(input []string) (int, error) {
	if s.part2Err != nil {
		return 0, s.part2Err
	}
	return s.part2, nil
}

func TestRunnerExecuteSuccess(t *testing.T) {
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error {
			return reg.Register(year, 6, stubSolver{part1: 123, part2: 456})
		},
		func(year int, day int) ([]string, error) {
			return []string{"input"}, nil
		},
	)

	result, err := r.Execute(2015, 6)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result.Part1 != 123 || result.Part2 != 456 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestRunnerExecuteNilRunner(t *testing.T) {
	var r *Runner

	_, err := r.Execute(2015, 1)
	if !errors.Is(err, ErrNilRunner) {
		t.Fatalf("expected ErrNilRunner, got %v", err)
	}
}

func TestRunnerExecuteNotConfigured(t *testing.T) {
	r := &Runner{}

	_, err := r.Execute(2015, 1)
	if !errors.Is(err, ErrRunnerNotConfigured) {
		t.Fatalf("expected ErrRunnerNotConfigured, got %v", err)
	}
}

func TestRunnerExecuteNilRegistry(t *testing.T) {
	r := NewRunner(
		func() *registry.Registry { return nil },
		func(reg *registry.Registry, year int) error { return nil },
		func(year int, day int) ([]string, error) { return []string{}, nil },
	)

	_, err := r.Execute(2015, 1)
	if !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("expected ErrNilRegistry, got %v", err)
	}
}

func TestRunnerExecuteRegisterYearError(t *testing.T) {
	wantErr := errors.New("unsupported year")
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error { return wantErr },
		func(year int, day int) ([]string, error) { return []string{}, nil },
	)

	_, err := r.Execute(2099, 1)
	var registerErr *ErrRegisterYear
	if !errors.As(err, &registerErr) {
		t.Fatalf("expected ErrRegisterYear, got %v", err)
	}
	if registerErr.Year != 2099 {
		t.Fatalf("expected year 2099, got %d", registerErr.Year)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}

func TestRunnerExecuteSolverNotFound(t *testing.T) {
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error { return nil },
		func(year int, day int) ([]string, error) { return []string{}, nil },
	)

	_, err := r.Execute(2015, 2)
	var solverErr *ErrSolverNotFound
	if !errors.As(err, &solverErr) {
		t.Fatalf("expected ErrSolverNotFound, got %v", err)
	}
	if solverErr.Year != 2015 || solverErr.Day != 2 {
		t.Fatalf("unexpected solver error payload: %+v", solverErr)
	}
}

func TestRunnerExecuteGetInputError(t *testing.T) {
	wantErr := errors.New("input missing")
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error {
			return reg.Register(year, 1, stubSolver{})
		},
		func(year int, day int) ([]string, error) { return nil, wantErr },
	)

	_, err := r.Execute(2015, 1)
	var inputErr *ErrGetInput
	if !errors.As(err, &inputErr) {
		t.Fatalf("expected ErrGetInput, got %v", err)
	}
	if inputErr.Year != 2015 || inputErr.Day != 1 {
		t.Fatalf("unexpected input error payload: %+v", inputErr)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}

func TestRunnerExecuteSolvePart1Error(t *testing.T) {
	wantErr := errors.New("part1 failed")
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error {
			return reg.Register(year, 1, stubSolver{part1Err: wantErr})
		},
		func(year int, day int) ([]string, error) { return []string{"ok"}, nil },
	)

	_, err := r.Execute(2015, 1)
	var solveErr *ErrSolvePart
	if !errors.As(err, &solveErr) {
		t.Fatalf("expected ErrSolvePart, got %v", err)
	}
	if solveErr.Part != 1 {
		t.Fatalf("expected part 1, got %d", solveErr.Part)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}

func TestRunnerExecuteSolvePart2Error(t *testing.T) {
	wantErr := errors.New("part2 failed")
	r := NewRunner(
		registry.NewRegistry,
		func(reg *registry.Registry, year int) error {
			return reg.Register(year, 1, stubSolver{part1: 10, part2Err: wantErr})
		},
		func(year int, day int) ([]string, error) { return []string{"ok"}, nil },
	)

	_, err := r.Execute(2015, 1)
	var solveErr *ErrSolvePart
	if !errors.As(err, &solveErr) {
		t.Fatalf("expected ErrSolvePart, got %v", err)
	}
	if solveErr.Part != 2 {
		t.Fatalf("expected part 2, got %d", solveErr.Part)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}
