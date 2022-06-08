// Package sentinel provides some utilities for easier use of sentinel errors.
package sentinel

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorWrapper wraps the original error.
type ErrorWrapper interface {
	OriginalError() error
}

type wrappedError struct {
	original error
	sentinel error
}

// Error implements the error interface.
func (w wrappedError) Error() string {
	return fmt.Sprintf("%v: %+v", w.sentinel, w.original)
}

// Unwrap implements the Wrapper interface.
func (w wrappedError) Unwrap() error {
	return w.sentinel
}

// OriginalError returns the original error.
func (w wrappedError) OriginalError() error {
	return w.original
}

// Wrap annotates the original err with a sentinel error.
func Wrap(original error, sentinel error) error {
	return wrappedError{
		original: original,
		sentinel: sentinel,
	}
}

// WithMessage annotates the original err with a sentinel error and a message.
func WithMessage(original error, sentinel error, context string) error {
	return wrappedError{
		original: original,
		sentinel: errors.WithMessage(sentinel, context),
	}
}

// WithMessagef annotates the original err with a sentinel error and a format specifier.
func WithMessagef(original error, sentinel error, contextFormat string, args ...interface{}) error {
	return wrappedError{
		original: original,
		sentinel: errors.WithMessagef(sentinel, fmt.Sprintf(contextFormat, args...)),
	}
}
