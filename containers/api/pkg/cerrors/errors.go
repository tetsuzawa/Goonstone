package cerrors

import "fmt"

var (
	ErrCanceled           = New(StatusCanceled)
	ErrUnknown            = New(StatusUnknown)
	ErrInvalidArgument    = New(StatusInvalidArgument)
	ErrDeadlineExceeded   = New(StatusDeadlineExceeded)
	ErrNotFound           = New(StatusNotFound)
	ErrAlreadyExists      = New(StatusAlreadyExists)
	ErrPermissionDenied   = New(StatusPermissionDenied)
	ErrResourceExhausted  = New(StatusResourceExhausted)
	ErrFailedPrecondition = New(StatusFailedPrecondition)
	ErrAborted            = New(StatusAborted)
	ErrOutOfRange         = New(StatusOutOfRange)
	ErrUnimplemented      = New(StatusUnimplemented)
	ErrInternal           = New(StatusInternal)
	ErrUnavailable        = New(StatusUnavailable)
	ErrDataLoss           = New(StatusDataLoss)
	ErrUnauthenticated    = New(StatusUnauthenticated)
)

func New(c Code, message ...interface{}) error {
	ce := CommonError{
		Code:    c,
		Message: c.String(),
	}
	if len(message) > 0 {
		ce.Message = message[0]
	}
	return ce
}

type CommonError struct {
	Code     Code
	Message  interface{}
	Internal error
}

func (e CommonError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
