package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

func TestDay03Part01(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := []TestCase{
		{
			desc: "delivers presents to 2 houses: one at the starting location, and one to the east",
			input: []string{
				">",
			},
			expectedResult: 2,
		},
		{
			desc: "delivers presents to 4 houses in a square, including twice to the house at his starting/ending location",
			input: []string{
				"^>v<",
			},
			expectedResult: 4,
		},
		{
			desc: "delivers a bunch of presents to some very lucky children at only 2 houses",
			input: []string{
				"^v^v^v^v^v",
			},
			expectedResult: 2,
		},
	}

	registry := registry.NewRegistry()
	solver, ok := registry.GetSolver(2015, 0o3)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}

func TestDay03Part02(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := []TestCase{
		{
			desc: "delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.",
			input: []string{
				"^v",
			},
			expectedResult: 3,
		},
		{
			desc: "now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.",
			input: []string{
				"^>v<",
			},
			expectedResult: 3,
		},
		{
			desc: "delivers a bunch of presents to some very lucky children at only 2 houses",
			input: []string{
				"^v^v^v^v^v",
			},
			expectedResult: 11,
		},
	}

	registry := registry.NewRegistry()
	solver, ok := registry.GetSolver(2015, 0o3)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}
