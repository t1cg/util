package apperror

import (
	"errors"
	"testing"
)

// TestAppError function tests the functions for the AppInfo struct.
func TestAppError(t *testing.T) {
	e1 := AppInfo{Msg: errors.New("error message with caller and callee")}
	e1.LogError(e1.Error())

	e2 := AppInfo{Msg: errors.New("error message with added message")}
	e2.LogError(e2.Error("this is the added msg"))
}
