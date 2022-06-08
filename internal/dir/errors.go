package dir

// Error is the error type.
type Error uint

const (
	// ErrNoSubDirectories is the error thrown when no sub directories exist in
	// the specified parent directory.
	ErrNoSubDirectories Error = iota
)

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrNoSubDirectories:
		return "no sub directories found in parent"

	default:
		return "unknown error"
	}
}
