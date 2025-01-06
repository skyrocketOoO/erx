package erx

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

const (
	ID        = "ID"
	CallStack = "CallStack"
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Wrap wraps the given error with a call stack. If the error is already an
// ErrorCtx, it appends additional context (texts) to it. Otherwise,
// it converts the error to an ErrorCtx and records the call stack.
func W(err error, texts ...string) error {
	if err == nil {
		return nil
	}

	var newErr error
	if len(texts) > 0 {
		errs := make([]error, len(texts))
		for i, text := range texts {
			errs[i] = errors.New(text)
		}
		newErr = errors.Join(errs...)
	}

	var errCStk *ErrorCtx
	if errors.As(err, &errCStk) {
		errCStk.OriginalErr = errors.Join(newErr, errCStk.OriginalErr)
	} else {
		errCStk = &ErrorCtx{
			Ctx: map[string]string{
				ID:        uuid.NewString(),
				CallStack: getCallStack(3),
			},
			OriginalErr: errors.Join(newErr, err),
		}
	}

	return errCStk
}

func New(text string) *ErrorCtx {
	return &ErrorCtx{
		Ctx: map[string]string{
			ID:        uuid.NewString(),
			CallStack: getCallStack(2),
		},
		OriginalErr: errors.New(text),
	}
}

func Errorf(format string, a ...any) error {
	return W(fmt.Errorf(format, a...))
}

// default use 2
func getCallStack(callerSkip ...int) (stkMsg string) {
	pc := make([]uintptr, 10)
	skip := 2
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}
	n := runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc[:n-2])

	stkMsg = "Call stack:\n"

	for {
		frame, more := frames.Next()
		stkMsg += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return stkMsg
}

func Cause(err error) error {
	if jErr, ok := err.(interface{ Unwrap() []error }); ok {
		joinErrs := jErr.Unwrap()

		if len(joinErrs) > 0 {
			return joinErrs[len(joinErrs)-1]
		}
	}

	return err
}

func GetCallStack(err error) string {
	var errCStk *ErrorCtx
	if !errors.As(err, &errCStk) {
		return ""
	}

	return errCStk.Ctx[CallStack]
}

func IsErrCtx(err error) bool {
	var errCStk *ErrorCtx
	if !errors.As(err, &errCStk) {
		return false
	}
	return true
}

type ClientMsg struct {
	ID  string
	Err string
}

func GetClientMsg(err error) (msg ClientMsg, ok bool) {
	var errCStk *ErrorCtx
	if !errors.As(err, &errCStk) {
		return ClientMsg{}, false
	}

	return ClientMsg{
		ID:  errCStk.Ctx[ID],
		Err: fPrintErr(errCStk),
	}, true
}

type FullMsg struct {
	ID        string
	Err       string
	CallStack string
}

func GetFullMsg(err error) (msg FullMsg, ok bool) {
	var errCStk *ErrorCtx
	if !errors.As(err, &errCStk) {
		return FullMsg{}, false
	}

	return FullMsg{
		ID:        errCStk.Ctx[ID],
		Err:       fPrintErr(errCStk),
		CallStack: errCStk.Ctx[CallStack],
	}, true
}

func fPrintErr(err *ErrorCtx) string {
	eg := err.OriginalErr.Error()
	return strings.ReplaceAll(eg, "\n", ": ")
}
