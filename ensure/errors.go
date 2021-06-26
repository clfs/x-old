package ensure

import "fmt"

type (
	NoSumError       struct{}
	FailedReadError  struct{ Err error }
	FailedWriteError struct{ Err error }
	WrongSumError    struct{ Want, Got []byte }
)

func (e *NoSumError) Error() string {
	return "no sum provided"
}

func (e *FailedReadError) Error() string {
	return fmt.Sprintf("failed read: %v", e.Err)
}

func (e *FailedWriteError) Error() string {
	return fmt.Sprintf("failed write: %v", e.Err)
}

func (e *WrongSumError) Error() string {
	return fmt.Sprintf("wrong sum: want %x, got %x", e.Want, e.Got)
}
