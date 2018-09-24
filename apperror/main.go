package apperror

import (
	"fmt"
	//t1cg library
	"github.com/t1cg/util/logger"
)

// AppInfo struct type is a custom error handler that inherits the standard Error() method.
type AppInfo struct {
	Msg error
	Rc  int
}

// Error function returns a custom, parsed string; extends the standad Error() function.
func (a AppInfo) Error(msg ...interface{}) string {

	//get the caller and callee function names
	fname := logger.GetFuncName()

	if msg != nil {
		if a.Rc > 0 {
			return fmt.Sprintf("%v message[%v], rc[%v], added message%v", fname, a.Msg.Error(), a.Rc, msg)
		}

		return fmt.Sprintf("%v message[%v], added message%v", fname, a.Msg.Error(), msg)
	}

	if a.Rc > 0 {
		return fmt.Sprintf("%v message[%v], rc[%v]", fname, a.Msg.Error(), a.Rc)
	}

	return fmt.Sprintf("%v message[%v]", fname, a.Msg.Error())
}

// LogError function writes the error string to the stdout.
func (a AppInfo) LogError(e string) {
	logger.L.Error.Println(e)
}

// LogInfo function writes the info string to the stdout.
func (a AppInfo) LogInfo(e string) {
	logger.L.Info.Println(e)
}
