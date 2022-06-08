// Package test provides utilities to help with unit tests.
package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// IsSentinelError returns a require.ErrorAssertionFunc compatible func
// that checks if the error being tested is equal to the given sentinel error.
func IsSentinelError(target error) require.ErrorAssertionFunc {
	return func(t require.TestingT, err error, _ ...interface{}) {
		require.Error(t, err)

		if !assert.True(t, errors.Is(err, target)) {
			if tt, ok := t.(*testing.T); ok {
				tt.Logf("incorrect error type for error '%s' expected %T, got %T", err, target, err)
			}
		}
	}
}
