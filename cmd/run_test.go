package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	runusecase "github.com/ihribernik/aoc-cli/internal/run"
)

func TestMapRunError(t *testing.T) {
	registerCause := errors.New("unsupported year")
	inputCause := errors.New("permission denied")
	solveCause := errors.New("boom")
	fallbackCause := errors.New("unexpected")

	tests := []struct {
		name        string
		year        int
		day         int
		err         error
		wantContain string
		wantIs      error
	}{
		{
			name:        "register year",
			year:        2099,
			day:         1,
			err:         &runusecase.ErrRegisterYear{Year: 2099, Err: registerCause},
			wantContain: "year 2099 is not available",
			wantIs:      registerCause,
		},
		{
			name:        "solver not found",
			year:        2015,
			day:         9,
			err:         &runusecase.ErrSolverNotFound{Year: 2015, Day: 9},
			wantContain: "no solver registered for year 2015 day 09",
		},
		{
			name:        "input file missing",
			year:        2015,
			day:         2,
			err:         &runusecase.ErrGetInput{Year: 2015, Day: 2, Err: fmt.Errorf("open x: %w", os.ErrNotExist)},
			wantContain: "input file not found for year 2015 day 02",
		},
		{
			name:        "input other error",
			year:        2015,
			day:         2,
			err:         &runusecase.ErrGetInput{Year: 2015, Day: 2, Err: inputCause},
			wantContain: "cannot load input for year 2015 day 02",
			wantIs:      inputCause,
		},
		{
			name:        "solve part",
			year:        2015,
			day:         2,
			err:         &runusecase.ErrSolvePart{Part: 2, Err: solveCause},
			wantContain: "failed while solving part 2",
			wantIs:      solveCause,
		},
		{
			name:        "fallback",
			year:        2015,
			day:         2,
			err:         fallbackCause,
			wantContain: "run failed for year 2015 day 02",
			wantIs:      fallbackCause,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mapRunError(tc.year, tc.day, tc.err)
			if got == nil {
				t.Fatalf("expected error, got nil")
			}
			if !strings.Contains(got.Error(), tc.wantContain) {
				t.Fatalf("expected message to contain %q, got %q", tc.wantContain, got.Error())
			}
			if tc.wantIs != nil && !errors.Is(got, tc.wantIs) {
				t.Fatalf("expected mapped error to wrap %v, got %v", tc.wantIs, got)
			}
		})
	}
}
