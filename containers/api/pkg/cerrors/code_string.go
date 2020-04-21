package cerrors

import "strconv"

func (c Code) String() string {
	switch c {
	case StatusOK:
		return "OK"
	case StatusCanceled:
		return "Canceled"
	case StatusUnknown:
		return "Unknown"
	case StatusInvalidArgument:
		return "InvalidArgument"
	case StatusDeadlineExceeded:
		return "DeadlineExceeded"
	case StatusNotFound:
		return "NotFound"
	case StatusAlreadyExists:
		return "AlreadyExists"
	case StatusPermissionDenied:
		return "PermissionDenied"
	case StatusResourceExhausted:
		return "ResourceExhausted"
	case StatusFailedPrecondition:
		return "FailedPrecondition"
	case StatusAborted:
		return "Aborted"
	case StatusOutOfRange:
		return "OutOfRange"
	case StatusUnimplemented:
		return "Unimplemented"
	case StatusInternal:
		return "Internal"
	case StatusUnavailable:
		return "Unavailable"
	case StatusDataLoss:
		return "DataLoss"
	case StatusUnauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
