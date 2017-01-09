package logger

import (
	"log"
	"testing"
)

// TestGetLogLevel function tests the SetLogStream method
func TestGetLogLevel(t *testing.T) {

	t.Logf("running...")

	//get the caller/callee names
	fname := GetFuncName(1)

	if len(fname) == 0 {
		t.Errorf("caller did not return the expected string, %v", fname)
		return
	}

	level := string(L.Level)

	if level == "" {
		t.Errorf("%s did not return the expected loglevel; instead it returned, %v", fname, level)
		return
	}

	t.Logf("%s loglevel is %v", fname, level)

	t.Logf("exiting")
}

// TestGetFuncName function tests the SetLogStream method
func TestGetFuncName(t *testing.T) {

	t.Logf("running...")

	fname := GetFuncName()

	if len(fname) == 0 {
		t.Errorf("caller did not return the expected string, %v", fname)
		return
	}

	t.Logf("%s", fname)
	t.Logf("exiting")
}

// TestLogStream function tests the SetLogStream method
func TestLogStream(t *testing.T) {

	t.Logf("running...")

	fname := GetFuncName()

	if len(fname) == 0 {
		t.Errorf("caller did not return the expected string, %v", fname)
		return
	}

	l := LogInfo{}

	for _, level := range levels {
		log.Println("*****setting:", string(level))
		l.SetLogStream(level)

		l.Trace.Printf("%s trace: trace line", fname)
		l.Perf.Printf("%s perf: perf line", fname)
		l.Info.Printf("%s info: info line", fname)
		l.Warn.Printf("%s warn: warn line", fname)
		l.Error.Printf("%s error: error line", fname)
	}

	t.Logf("exiting")

	TestGetLogLevel(t)
}

// TestLogFile function tests the SetLogFile method
func TestLogFile(t *testing.T) {
	t.Logf("running...")

	fname := GetFuncName()

	if len(fname) == 0 {
		t.Errorf("caller did not return the expected string, %v", fname)
		return
	}

	l := LogInfo{}
	fpath := "./logs/"
	logname := "test"

	for _, level := range levels {
		log.Println("*****setting:", string(level))
		l.SetLogFile(fpath, logname, level)

		l.Trace.Printf("%s trace: trace line", fname)
		l.Perf.Printf("%s perf: perf line", fname)
		l.Info.Printf("%s info: info line", fname)
		l.Warn.Printf("%s warn: warn line", fname)
		l.Error.Printf("%s error: error line", fname)
	}

	t.Logf("exiting")

	TestGetLogLevel(t)
}
