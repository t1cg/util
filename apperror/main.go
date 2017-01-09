package apperror

import (
	"fmt"
	//t1cg library
	"github.com/t1cg/util/logger"
)

// AppInfo struct type is a custom error handler that inherits the standard Error() method.
type AppInfo struct {
	Msg error
}

// Error function returns a custom, parsed string; extends the standad Error() function.
func (a AppInfo) Error() string {

	//get the caller and callee function names
	fname := logger.GetFuncName()

	return fmt.Sprintf("%v [message: %v]", fname, a.Msg.Error())
}

// LogError function writes the error string to the stdout.
func (a AppInfo) LogError(e string) {
	logger.L.Error.Println(e)
}
