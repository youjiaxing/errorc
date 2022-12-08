package errorc

import (
	"fmt"
	"io"
)

type codeError struct {
	cause error
	msg   string
	code  int
	*stack
}

func (e *codeError) Unwrap() error {
	return e.cause
}

func (e *codeError) Cause() error {
	return e.cause
}

func (e *codeError) Error() string {
	//if e.msg == "" {
	//	if e.cause != nil {
	//		return e.cause.Error()
	//	}
	//}
	return e.msg
}

func (e *codeError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			if e.cause != nil {
				_, _ = fmt.Fprintf(s, "%+v\n", e.cause)
			}

			fmt.Fprintf(s, "code: %d, msg: %s", e.code, e.msg)

			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		if s.Flag('-') {
			if e.cause != nil {
				_, _ = fmt.Fprintf(s, "%-v\n", e.cause)
			}

			if e.code != 0 {
				fmt.Fprintf(s, "code: %d, msg: %s", e.code, e.msg)
			} else {
				io.WriteString(s, e.msg)
			}

			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

func New(format string, args ...interface{}) error {
	return &codeError{
		cause: nil,
		msg:   fmt.Sprintf(format, args...),
		code:  0,
		stack: callers(2),
	}
}

func NewC(code int, format string, args ...interface{}) error {
	return &codeError{
		cause: nil,
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		stack: callers(2),
	}
}

func Wrap(err error, format string, args ...interface{}) error {
	if cause, ok := err.(*codeError); ok {
		return &codeError{
			cause: cause,
			msg:   fmt.Sprintf(format, args...),
			code:  cause.code,
			stack: callers(2),
		}
	}
	return &codeError{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
		code:  0,
		stack: callers(2),
	}
}

func WrapC(err error, code int, format string, args ...interface{}) error {
	return &codeError{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		stack: callers(2),
	}
}

func WithCode(err error, code int, format string, args ...interface{}) error {
	if cause, ok := err.(*codeError); ok {
		return &codeError{
			cause: cause,
			msg:   fmt.Sprintf(format, args...),
			code:  code,
			stack: nil,
		}
	}
	return &codeError{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		stack: callers(2),
	}
}

func WithMessage(err error, format string, args ...interface{}) error {
	if cause, ok := err.(*codeError); ok {
		return &codeError{
			cause: cause,
			msg:   fmt.Sprintf(format, args...),
			code:  cause.code,
			stack: nil,
		}
	}
	return &codeError{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
		code:  0,
		stack: callers(2),
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err2 := cause.Cause()
		if err2 == nil {
			break
		}
		err = err2
	}
	return err
}
