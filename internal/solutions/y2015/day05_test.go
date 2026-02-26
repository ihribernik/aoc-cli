package y2015_test

import "testing"

func TestDay05Part01(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}
	testCases := TestCase{
		{desc: "one long nice string", input: []string{"ugknbfddgicrmopn"}, expectedResult: 1},
		{desc: "one short nice string", input: []string{"aaa"}, expectedResult: 1},
		{desc: "one long naughty string", input: []string{"jchzalrnumimnmhp"}, expectedResult: 0},
		{desc: "second long naughty string", input: []string{"haegwjzuvuyypxyu"}, expectedResult: 0},
		{desc: "another long naughty string", input: []string{"dvszwmarrgswjxmb"}, expectedResult: 0},
	}

	solver := mustSolver(t, 0o5)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}

func TestDay05Part02(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := TestCase{
		{desc: "one long nice string", input: []string{"qjhvhtzxzqqjkmpb"}, expectedResult: 1},
		{desc: "one short nice string", input: []string{"xxyxx"}, expectedResult: 1},
		{desc: "one long naughty string", input: []string{"uurcxstgmygtbstg"}, expectedResult: 0},
		{desc: "second long naughty string", input: []string{"ieodomkazucvgmuy"}, expectedResult: 0},
	}

	solver := mustSolver(t, 0o5)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}
