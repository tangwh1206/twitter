package faults

import "fmt"

type Code int

type Fault interface {
	Code() Code
	Message() string
	Error() string
}

type fault struct {
	code    Code
	message string
	err     error
}

func New(code Code, message string, err error) *fault {
	return &fault{
		code:    code,
		message: message,
		err:     err,
	}
}

func (f *fault) Code() Code {
	return f.code
}

func (f *fault) Message() string {
	return f.message
}

func (f *fault) Error() string {
	return fmt.Sprintf("FAULT-%d: %s", f.code, f.message)
}
