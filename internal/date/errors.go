package date

// Error is the error type.
type Error uint8

const (
	// ErrNoCaptureGroupDay is the error returned when the regex capture group is not found in the date string.
	ErrNoCaptureGroupDay Error = iota + 1
	// ErrParsingDayAsInt is the error returned when the extracted day value cannot be parsed as an int.
	ErrParsingDayAsInt
)

// Error returns the message string for the given error.
func (e Error) Error() string {
	switch e {
	case ErrNoCaptureGroupDay:
		return "capture group 'day' not found in string"

	case ErrParsingDayAsInt:
		return "error parsing day as int"

	default:
		return "unknown error"
	}
}
