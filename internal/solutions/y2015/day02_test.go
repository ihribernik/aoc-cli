package y2015_test

import "testing"

func TestDay02Part1(t *testing.T) {
	type TestCase struct {
		name     string
		input    []string
		expected int
		wantErr  bool
	}

	testCases := []TestCase{
		{name: "example 2x3x4", input: []string{"2x3x4"}, expected: 58},
		{name: "example 1x1x10", input: []string{"1x1x10"}, expected: 43},
		{name: "multiple boxes", input: []string{"2x3x4", "1x1x10"}, expected: 101},
		{name: "invalid dimension", input: []string{"2x3xa"}, wantErr: true},
		{name: "missing dimension", input: []string{"2x3"}, wantErr: true},
		{name: "extra dimension", input: []string{"2x3x4x5"}, wantErr: true},
		{name: "empty dimension", input: []string{""}, wantErr: true},
		{name: "empty input", input: []string{}, expected: 0},
	}

	solver := mustSolver(t, 2)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := solver.SolvePart1(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestDay02Part2(t *testing.T) {
	type TestCase struct {
		name     string
		input    []string
		expected int
		wantErr  bool
	}

	testCases := []TestCase{
		{name: "example 2x3x4", input: []string{"2x3x4"}, expected: 34},
		{name: "example 1x1x10", input: []string{"1x1x10"}, expected: 14},
		{name: "multiple boxes", input: []string{"2x3x4", "1x1x10"}, expected: 48},
		{name: "invalid dimension", input: []string{"2x3xa"}, wantErr: true},
		{name: "missing dimension", input: []string{"2x3"}, wantErr: true},
		{name: "extra dimension", input: []string{"2x3x4x5"}, wantErr: true},
		{name: "empty dimension", input: []string{""}, wantErr: true},
		{name: "empty input", input: []string{}, expected: 0},
	}

	solver := mustSolver(t, 2)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := solver.SolvePart2(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
