package y2015_test

import "testing"

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
				"123 -> a", // fake assignment to 'a' to test overall functionality
			},
			expectedResult: 123,
		},
	}
	solver := mustSolver(t, 0o7)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}

func TestDay07Part2(t *testing.T) {
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
				"123 -> a", // fake assignment to 'a' to test overall functionality
			},
			expectedResult: 123, // fake assignment to 'a' to test overall functionality

		},
	}
	solver := mustSolver(t, 0o7)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}
