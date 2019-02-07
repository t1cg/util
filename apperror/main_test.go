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
	e2.LogError(e2.Error("this is the added msg e2"))

	e3 := AppInfo{Msg: errors.New("error message with return code"), Rc: 200}
	e3.LogError(e3.Error())

	e4 := AppInfo{Msg: errors.New("error message with added message & rc"), Rc: 404}
	e4.LogError(e4.Error("this is the added msg e4"))

	s1 := "string1"
	s2 := "string2"
	e5 := AppInfo{Msg: errors.New("info message with added message & rc"), Rc: 500}
	e5.LogInfo(e5.Errorf("this is the added msg e5[%v], e5[%v]", s1, s2))
}
