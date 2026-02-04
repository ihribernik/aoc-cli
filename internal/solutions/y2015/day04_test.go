package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

func TestDay04Part01(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := []TestCase{
		{
			desc: "answer is 609043...",
			input: []string{
				"abcdef",
			},
			expectedResult: 609043,
		},
		{
			desc: "anwser is 1048970...",
			input: []string{
				"pqrstuv",
			},
			expectedResult: 1048970,
		},
	}

	registry := registry.NewRegistry()
	solver, ok := registry.GetSolver(2015, 0o4)

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
