package prompt

// Error is the error type.
type Error uint8

const (
	// ErrNoTemplateWithName is the error returned when a template with the name selected cannot be found.
	ErrNoTemplateWithName Error = iota + 1
	// ErrSelectingTemplate is the error returned when something goes wrong selecting a template.
	ErrSelectingTemplate
	// ErrGettingFilesInDirectory is the error returned when something goes wrong getting the
	// files inside the specified directory.
	ErrGettingFilesInDirectory
)

// Error returns the message string for the given error.
func (e Error) Error() string {
	switch e {
	case ErrNoTemplateWithName:
		return "no template in config file with name"

	case ErrSelectingTemplate:
		return "error selecting template"

	case ErrGettingFilesInDirectory:
		return "error getting files in directory"

	default:
		return "unknown error"
	}
}
