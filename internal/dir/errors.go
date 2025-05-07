package dir

// Error is the error type.
type Error uint

const (
	// ErrNoSubDirectories is the error returned when no sub directories exist in
	// the specified parent directory.
	ErrNoSubDirectories Error = iota
	// ErrCreatingParentDirectories is the error returned when something goes wrong creating parent directories.
	ErrCreatingParentDirectories
	// ErrReadingDirectory is the error returned when something dodes wrong reading the directory.
	ErrReadingDirectory
	// ErrInvalidDirectoryPath is the error returned when you pass a path that doesn't exist
	// to GetAllFilesInDirectory.
	ErrInvalidDirectoryPath
)

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrNoSubDirectories:
		return "no sub directories found in parent"

	case ErrCreatingParentDirectories:
		return "error creating parent directory for file"

	case ErrReadingDirectory:
		return "error reading directory"

	case ErrInvalidDirectoryPath:
		return "invalid directory path"

	default:
		return "unknown error"
	}
}
