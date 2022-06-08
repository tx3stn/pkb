package config

const (
	// ErrConfigNotFound is the error thrown when a config file cannot be found.
	ErrConfigNotFound Error = iota
	// ErrReadingConfigFile is the error thrown when the config file cannot be parsed.
	ErrReadingConfigFile
	// ErrUnmashallingJSON is the error thrown when the provided config file can't be unmarshalled.
	ErrUnmashallingJSON
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

	default:
		return "unknown error"
	}
}
