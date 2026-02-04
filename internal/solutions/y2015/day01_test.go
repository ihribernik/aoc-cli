package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

func TestDay01_SolvePart1(t *testing.T) {
	type TestCase struct {
		input          []string
		expectedResult int
	}

	testsCases := []TestCase{
		{[]string{"("}, 1},
		{[]string{"(", ")"}, 0},
		{[]string{"(", "(", ")"}, 1},
		{[]string{"(", ")", "(", ")"}, 0},
		{[]string{"(", "(", ")", "(", ")"}, 1},
	}

	registry := registry.NewRegistry()
	err := solutions.RegisterYear(registry, 2015)
	if err != nil {
		t.Errorf("failed to load registry for the year %d, err: %d", 2015, err)
	}

	solver, ok := registry.GetSolver(2015, 1)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for i, testCase := range testsCases {
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart1(testCase.input)
			if err != nil {
				if result != testCase.expectedResult {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expectedResult, result)
				}
			}
		})
	}
}

func TestDay01_SolvePart2(t *testing.T) {
	type TestCase struct {
		input    []string
		expected int
	}

	testsCases := []TestCase{
		{[]string{"("}, 1},
		{[]string{"(", ")"}, 0},
		{[]string{"(", "(", ")"}, 1},
		{[]string{"(", ")", "(", ")"}, 0},
		{[]string{"(", "(", ")", "(", ")"}, 1},
	}

	registry := registry.NewRegistry()
	err := solutions.RegisterYear(registry, 2015)
	if err != nil {
		t.Errorf("failed to load registry for the year %d, err: %d", 2015, err)
	}

	solver, ok := registry.GetSolver(2015, 1)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for i, testCase := range testsCases {
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart2(testCase.input)
			if err != nil {
				if result != testCase.expected {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expected, result)
				}
			}
		})
	}
}
