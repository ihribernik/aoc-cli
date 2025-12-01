package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/solutions"
)

func TestDay07Part1(t *testing.T) {

	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := TestCase{
		{
			desc: "test circuit",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			expectedResult: 1,
		},
	}
	solver, ok := solutions.GetSolver(2015, 07)

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
