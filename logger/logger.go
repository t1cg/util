package logger

import (
	//standard lib
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//global, global
var (
	L *LogInfo
)

//global, local
var (
	//content/texts to ignore
	ignores = [...]string{
		//"main.init",
		//"main.main",
		"runtime.goexit",
		"runtime.main",
		"testing.tRunner",
		"http.(*conn).serve",
		"http.(*ServeMux).ServeHTTP",
		"http.serverHandler.ServeHTTP",
		".1",
	}

	//log levels
	levels = [...]byte{TRACE, PERF, INFO, WARN, ERROR, UNKNOWN}
)

// Constant byte representing the log levels
const (
	TRACE     byte   = 'T'
	PERF      byte   = 'P'
	INFO      byte   = 'I'
	WARN      byte   = 'W'
	ERROR     byte   = 'E'
	UNKNOWN   byte   = 'U'
	FILLER    string = " "
	SEPERATOR string = "->"
	//layers of stack to look into
	defaultStackCounter int = 10
)

// init() function will run as when this package is imported.
func init() {
	FUNCNAME := "logger.init():"

	defer measureRuntime(time.Now(), FUNCNAME)

	log.Println("running...", FUNCNAME)

	//set temporary logger to trace log level
	L = &LogInfo{}
	L.SetLogStream('T')
	L.Level = TRACE

	L.Info.Printf("%s loglevel temporarily set to: %s", FUNCNAME, string(L.Level))
}

// LogInfo defines log levels specified by log4j
type LogInfo struct {
	Trace *log.Logger
	Perf  *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Level byte
}

// SetLogFile creates and returns the logger that can be written to the defined log file.
func (l *LogInfo) SetLogFile(fpath string, logname string, level byte) error {
	FUNCNAME := "SetLogFile()"

	logFile, err := os.OpenFile(fpath+string(filepath.Separator)+logname+"_app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	msg := "successful! loglevel set to:"

	//set the log level
	switch level {
	case TRACE:
		l.Trace = getLogger(false, TRACE, logFile)
		l.Perf = getLogger(false, PERF, logFile)
		l.Info = getLogger(false, INFO, logFile)
		l.Warn = getLogger(false, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = TRACE
		l.Trace.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case PERF:
		l.Trace = getLogger(true, TRACE, logFile)
		l.Perf = getLogger(false, PERF, logFile)
		l.Info = getLogger(false, INFO, logFile)
		l.Warn = getLogger(false, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = PERF
		l.Perf.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case INFO:
		l.Trace = getLogger(true, TRACE, logFile)
		l.Perf = getLogger(true, PERF, logFile)
		l.Info = getLogger(false, INFO, logFile)
		l.Warn = getLogger(false, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = INFO
		l.Info.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case WARN:
		l.Trace = getLogger(true, TRACE, logFile)
		l.Perf = getLogger(true, PERF, logFile)
		l.Info = getLogger(true, INFO, logFile)
		l.Warn = getLogger(false, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = WARN
		l.Warn.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case ERROR:
		l.Trace = getLogger(true, TRACE, logFile)
		l.Perf = getLogger(true, PERF, logFile)
		l.Info = getLogger(true, INFO, logFile)
		l.Warn = getLogger(true, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = ERROR
		l.Error.Printf("%s %s %s", FUNCNAME, msg, string(level))
	default:
		//default is info log level
		l.Trace = getLogger(true, TRACE, logFile)
		l.Perf = getLogger(true, PERF, logFile)
		l.Info = getLogger(false, INFO, logFile)
		l.Warn = getLogger(false, WARN, logFile)
		l.Error = getLogger(false, ERROR, logFile)
		l.Level = INFO
		l.Info.Printf("%s %s %s", FUNCNAME, msg, string(level))
	}

	//re-assign to the global variable
	L = l

	return nil
}

// SetLogStream creates and returns the logger that can be written to the IO stream.
func (l *LogInfo) SetLogStream(level byte) {
	FUNCNAME := "SetLogStream():"

	msg := "successful! loglevel set to:"

	//set the log level
	switch level {
	case TRACE:
		l.Trace = getLogger(false, TRACE, nil)
		l.Perf = getLogger(false, PERF, nil)
		l.Info = getLogger(false, INFO, nil)
		l.Warn = getLogger(false, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = TRACE
		l.Trace.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case PERF:
		l.Trace = getLogger(true, TRACE, nil)
		l.Perf = getLogger(false, PERF, nil)
		l.Info = getLogger(false, INFO, nil)
		l.Warn = getLogger(false, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = PERF
		l.Perf.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case INFO:
		l.Trace = getLogger(true, TRACE, nil)
		l.Perf = getLogger(true, PERF, nil)
		l.Info = getLogger(false, INFO, nil)
		l.Warn = getLogger(false, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = INFO
		l.Info.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case WARN:
		l.Trace = getLogger(true, TRACE, nil)
		l.Perf = getLogger(true, PERF, nil)
		l.Info = getLogger(true, INFO, nil)
		l.Warn = getLogger(false, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = WARN
		l.Warn.Printf("%s %s %s", FUNCNAME, msg, string(level))
	case ERROR:
		l.Trace = getLogger(true, TRACE, nil)
		l.Perf = getLogger(true, PERF, nil)
		l.Info = getLogger(true, INFO, nil)
		l.Warn = getLogger(true, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = ERROR
		l.Error.Printf("%s %s %s", FUNCNAME, msg, string(level))
	default: //default is info log level
		l.Trace = getLogger(true, TRACE, nil)
		l.Perf = getLogger(true, PERF, nil)
		l.Info = getLogger(false, INFO, nil)
		l.Warn = getLogger(false, WARN, nil)
		l.Error = getLogger(false, ERROR, nil)
		l.Level = INFO
		l.Info.Printf("%s %s %s", FUNCNAME, msg, string(level))
	}

	//re-assign to the global variable
	L = l

}

// GetFuncName function ...
func GetFuncName(depth ...int) string {

	//if the stack argument is passed, then use it; otherwise, use the defaultStackCounter
	var stackCounter int

	if depth == nil {
		stackCounter = defaultStackCounter
	} else {
		stackCounter = depth[0]
	}

	//to be used inside the loop
	fname := ""
	count := 0

	//loop through the current stack, from the oldest to the newest
	for i := stackCounter; i > 0; i-- {
		//get the function pointer, and then the function name
		c1, _, _, _ := runtime.Caller(i)
		f := runtime.FuncForPC(c1)

		//if stack doesn't exist, iterate to the next stack
		if f == nil {
			continue
		}

		skip := false

		//check for internal go function names, and skip them
		for _, ignore := range ignores {
			if strings.Contains(f.Name(), ignore) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		//only pick up the function name and not the entire path
		fn := strings.Split(strings.TrimSpace(f.Name()), "/")

		//make the format look like; first.Method->second.Method->third.Method
		if count == 0 {
			fname = fn[len(fn)-1]
			count++
		} else {
			fname = fname + "->" + fn[len(fn)-1]
			count++
		}

		//fmt.Printf("\nfunction name[%v] for index[%v]:", fname, i)
	}

	return fname + ":"
}

//******************************************************************************
//PRIVATE functions

// measureRuntime prints a string containing the runtime information
func measureRuntime(start time.Time, name string) {
	stop := time.Since(start)
	L.Perf.Printf("%v runtime[%v]", name, stop)
}

// getLogger function returns the appropriate logger.
func getLogger(discard bool, level byte, file *os.File) *log.Logger {

	//instanciate a new logger
	l := &log.Logger{}

	//if writing log to file, set the io.MultiWriter
	if file != nil {
		//if discard, set ioutil.Discard (this will prevent the actual write to the log)
		if discard {
			l = log.New(ioutil.Discard, "", 0)
		} else {
			l = log.New(io.MultiWriter(file, os.Stderr), "", 0)
		}
	} else {
		if discard {
			l = log.New(ioutil.Discard, "", 0)
		} else {
			l = log.New(os.Stdout, "", 0)
		}
	}

	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05.000000") + " " + string(level) + " | ")

	//for error log level only, print out the line number
	if level == ERROR {
		l.SetFlags(log.Lshortfile)
	}

	return l
}
