package y2015_test

import "testing"

func TestDay01_SolvePart1(t *testing.T) {
	type TestCase struct {
		input          string
		expectedResult int
	}

	testsCases := []TestCase{
		{input: "(())", expectedResult: 0},
		{input: "()()", expectedResult: 0},
		{input: "(((", expectedResult: 3},
		{input: "(()(()(", expectedResult: 3},
		{input: "))(((((", expectedResult: 3},
		{input: "())", expectedResult: -1},
		{input: "))(", expectedResult: -1},
		{input: ")))", expectedResult: -3},
		{input: ")())())", expectedResult: -3},
	}

	solver := mustSolver(t, 1)

	for i, testCase := range testsCases {
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart1([]string{testCase.input})
			if err != nil || result != testCase.expectedResult {
				t.Fatalf("test %v: expected: %v, got: %v, err: %v", i+1, testCase.expectedResult, result, err)
			}
		})
	}
}

func TestDay01_SolvePart2(t *testing.T) {
	type TestCase struct {
		input    string
		expected int
	}

	testsCases := []TestCase{
		{input: ")", expected: 1},
		{input: "()())", expected: 5},
		{input: "(((((", expected: 0},
	}

	solver := mustSolver(t, 1)

	for i, testCase := range testsCases {
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart2([]string{testCase.input})
			if err != nil || result != testCase.expected {
				t.Fatalf("test %v: expected: %v, got: %v, err: %v", i+1, testCase.expected, result, err)
			}
		})
	}
}
