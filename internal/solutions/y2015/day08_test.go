package y2015_test

import (
	"strings"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

func TestDay08Part1(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := TestCase{
		{
			desc:           "empty string",
			input:          strings.Split(`"""abc""aaa\\"aaa""\\x27"`, "\n"),
			expectedResult: 12,
		},
	}
	registry := registry.NewRegistry()
	solver, ok := registry.GetSolver(2015, 8)

	if !ok {
		t.Errorf("failed to get solver")
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf("expected: %v, got: %v, errors: %v", tC.expectedResult, result, err)
			}
		})
	}
}

// func TestDay08Part2(t *testing.T) {

// 	type TestCase = []struct {
// 		desc           string
// 		input          []string
// 		expectedResult int
// 	}

// 	testCases := TestCase{
// 		{},
// 	}
// 	solver, ok := solutions.GetSolver(2015, 8)

// 	if !ok {
// 		t.Errorf("failed to get solver")
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			result, err := solver.SolvePart2(tC.input)
// 			if result != tC.expectedResult || err != nil {
// 				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)
// 			}
// 		})
// 	}
// }
