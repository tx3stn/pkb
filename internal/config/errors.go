package config

const (
	// ErrConfigNotFound is the error returned when a config file cannot be found.
	ErrConfigNotFound Error = iota
	// ErrReadingConfigFile is the error returned when the config file cannot be parsed.
	ErrReadingConfigFile
	// ErrUnmashallingJSON is the error returned when the provided config file can't be unmarshalled.
	ErrUnmashallingJSON
	// ErrNoDirectory is the error returned when the directory is not defined in the config file.
	ErrNoDirectory
)

// Error is the error type.
type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrConfigNotFound:
		return "no config file found"

	case ErrReadingConfigFile:
		return "error reading config file"

	case ErrUnmashallingJSON:
		return "error unmarshalling JSON config file"

	case ErrNoDirectory:
		return "no directory defined in config file"

	default:
		return "unknown error"
	}
}
