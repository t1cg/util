package apperror

import (
	"errors"
	"testing"
)

// TestAppError function tests the functions for the AppInfo struct.
func TestAppError(t *testing.T) {
	e := AppInfo{Msg: errors.New("error message with caller and callee")}
	e.LogError(e.Error())
}
