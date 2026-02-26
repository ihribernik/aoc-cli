package run

import "fmt"

var (
	ErrNilRunner           = fmt.Errorf("nil runner")
	ErrRunnerNotConfigured = fmt.Errorf("runner dependencies are not configured")
	ErrNilRegistry         = fmt.Errorf("registry factory returned nil")
)

type ErrRegisterYear struct {
	Year int
	Err  error
}

func (e *ErrRegisterYear) Error() string {
	return fmt.Sprintf("register year %d: %v", e.Year, e.Err)
}

func (e *ErrRegisterYear) Unwrap() error { return e.Err }

type ErrSolverNotFound struct {
	Year int
	Day  int
	Err  error
}

func (e *ErrSolverNotFound) Error() string {
	return fmt.Sprintf("cannot find a solution for year %d day %d", e.Year, e.Day)
}

func (e *ErrSolverNotFound) Unwrap() error { return e.Err }

type ErrGetInput struct {
	Year int
	Day  int
	Err  error
}

func (e *ErrGetInput) Error() string {
	return fmt.Sprintf("get input for year %d day %d: %v", e.Year, e.Day, e.Err)
}

func (e *ErrGetInput) Unwrap() error { return e.Err }

type ErrSolvePart struct {
	Part int
	Err  error
}

func (e *ErrSolvePart) Error() string {
	return fmt.Sprintf("solve part %d: %v", e.Part, e.Err)
}

func (e *ErrSolvePart) Unwrap() error { return e.Err }
