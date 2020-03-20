package cerrors // import "google.golang.org/grpc/codes"

import (
	"fmt"
	"strconv"
)

type Code uint32

const (
	StatusOK                 Code = 0
	StatusCanceled           Code = 1
	StatusUnknown            Code = 2
	StatusInvalidArgument    Code = 3
	StatusDeadlineExceeded   Code = 4
	StatusNotFound           Code = 5
	StatusAlreadyExists      Code = 6
	StatusPermissionDenied   Code = 7
	StatusResourceExhausted  Code = 8
	StatusFailedPrecondition Code = 9
	StatusAborted            Code = 10
	StatusOutOfRange         Code = 11
	StatusUnimplemented      Code = 12
	StatusInternal           Code = 13
	StatusUnavailable        Code = 14
	StatusDataLoss           Code = 15
	StatusUnauthenticated    Code = 16
	_maxCode                      = 17
)

var strToCode = map[string]Code{
	`"OK"`:                  StatusOK,
	`"CANCELLED"`:           StatusCanceled,
	`"UNKNOWN"`:             StatusUnknown,
	`"INVALID_ARGUMENT"`:    StatusInvalidArgument,
	`"DEADLINE_EXCEEDED"`:   StatusDeadlineExceeded,
	`"NOT_FOUND"`:           StatusNotFound,
	`"ALREADY_EXISTS"`:      StatusAlreadyExists,
	`"PERMISSION_DENIED"`:   StatusPermissionDenied,
	`"RESOURCE_EXHAUSTED"`:  StatusResourceExhausted,
	`"FAILED_PRECONDITION"`: StatusFailedPrecondition,
	`"ABORTED"`:             StatusAborted,
	`"OUT_OF_RANGE"`:        StatusOutOfRange,
	`"UNIMPLEMENTED"`:       StatusUnimplemented,
	`"INTERNAL"`:            StatusInternal,
	`"UNAVAILABLE"`:         StatusUnavailable,
	`"DATA_LOSS"`:           StatusDataLoss,
	`"UNAUTHENTICATED"`:     StatusUnauthenticated,
}

func (c *Code) UnmarshalJSON(b []byte) error {
	// From json.Unmarshaler: By convention, to approximate the behavior of
	// Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as
	// a no-op.
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 32); err == nil {
		if ci >= _maxCode {
			return fmt.Errorf("invalid code: %q", ci)
		}

		*c = Code(ci)
		return nil
	}

	if jc, ok := strToCode[string(b)]; ok {
		*c = jc
		return nil
	}
	return fmt.Errorf("invalid code: %q", string(b))
}
