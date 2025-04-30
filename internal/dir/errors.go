package dir

// Error is the error type.
type Error uint

const (
	// ErrNoSubDirectories is the error returned when no sub directories exist in
	// the specified parent directory.
	ErrNoSubDirectories Error = iota
	// ErrCreatingParentDirectories is the error returned when something goes wrong creating parent directories.
	ErrCreatingParentDirectories
)

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrNoSubDirectories:
		return "no sub directories found in parent"

	case ErrCreatingParentDirectories:
		return "error creating parent directory for file"

	default:
		return "unknown error"
	}
}
