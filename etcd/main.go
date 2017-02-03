package etcd

import (
	//standard library
	"bufio"
	"errors"
	"os/exec"
	"strings"
	"time"
	//t1cg library
	"github.com/t1cg/util/apperror"
	"github.com/t1cg/util/logger"
	"github.com/t1cg/util/runstat"
)

//etcdctl terminal commands
const (
	etcdStart       string = "etcd"
	etcdctlCmd      string = "etcdctl"
	etcdctlEndpoint string = "endpoint"
	etcdctlGet      string = "get"
	etcdctlPut      string = "put"
	etcdctlHealth   string = "health"
	etcdctlPrefix   string = "--prefix"
	etcdStop        string = "pkill"
)

//global
var (
	errorEtcdServiceDidNotReturnExpectedPrefix  = errors.New("Etcd service did not return the expected prefixes")
	errorEtcdReturnedMisMatchedPair             = errors.New("Etcd service returned a mis matched key/value pair")
	errorEtcdReturnedKeyDoesNotMatchProvidedKey = errors.New("Etcd service returned key does not match the provided key")
)

// Cml struct defines the method that calls the etcd service using the etcdctl
// command line interface.
type Cml struct{}

//function to start the etcd service during testing
func startService() {
	c := exec.Command(etcdStart)

	c.Start()
}

//function to kill the etcd service after testing completes
func stopService() {
	c := exec.Command(etcdStop, etcdStart)

	c.Start()
}

//function to assign key and values for testing
func putValue(key string, value string) {
	fname := logger.GetFuncName()

	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	args := []string{etcdctlPut, key, value}

	//set the command & args
	c := exec.Command(etcdctlCmd, args...)

	c.Start()

}

// GetValue function runs the command to get the key value pair.
// 1) The "etcdctl" must be in path.
// 2) the environment variable, ETCDCTL_API=3, must be set
func (ec *Cml) GetValue(key string, valueCh chan string, aeCh chan *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	//set the argument
	args := []string{etcdctlGet, key}

	//set the command & args
	c := exec.Command(etcdctlCmd, args...)

	//connect pipe to standard error when the command starts
	stderr, err := c.StderrPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//connect pipe to standard out when the command starts
	stdout, err := c.StdoutPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	err = c.Start()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//parse the stderr if any
	errReader := bufio.NewReader(stderr)
	errStr, _ := errReader.ReadString('\n')

	if len(errStr) > 0 {
		a := &apperror.AppInfo{Msg: errors.New(strings.TrimSpace(errStr))}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//parse the stdout; read each line
	lineCount := 0
	value := ""

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		lineCount++
		value = line
	}

	if lineCount == 0 {
		a := &apperror.AppInfo{Msg: errorEtcdServiceDidNotReturnExpectedPrefix}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	valueCh <- value
}

// GetPrefix function runs the command to see if the etcd service is running and
// in good health. The following must be set correctly:
// 1) The "etcdctl" must be in path.
// 2) the environment variable, ETCDCTL_API=3, must be set
func (ec *Cml) GetPrefix(prefix string, kvCh chan map[string]string, aeCh chan *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	//set the argument
	args := []string{etcdctlGet, etcdctlPrefix, prefix}

	//set the command & args
	c := exec.Command(etcdctlCmd, args...)

	//connect pipe to standard error when the command starts
	stderr, err := c.StderrPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//connect pipe to standard out when the command starts
	stdout, err := c.StdoutPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	err = c.Start()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	kvs := make(map[string]string)

	//parse the stderr if any
	errReader := bufio.NewReader(stderr)
	errStr, _ := errReader.ReadString('\n')

	if len(errStr) > 0 {
		a := &apperror.AppInfo{Msg: errors.New(strings.TrimSpace(errStr))}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//parse the stdout; read each line
	prefixCount := 0
	key := ""

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		prefixCount++

		//value
		if prefixCount%2 == 0 {
			kvs[key] = line
		} else {
			//key
			key = line
		}
	}

	if prefixCount == 0 {
		a := &apperror.AppInfo{Msg: errorEtcdServiceDidNotReturnExpectedPrefix}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	kvCh <- kvs
}

// HasPrefix function runs the etcd command line to see if the prefix is in etcd.
// The following must be set correctly:
// 1) The "etcdctl" must be in path.
// 2) the environment variable, ETCDCTL_API=3, must be set
func (ec *Cml) HasPrefix(prefix string, hasCh chan bool, aeCh chan *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	l := logger.L

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	//set the argument
	args := []string{etcdctlGet, etcdctlPrefix, prefix}

	//set the command & args
	c := exec.Command(etcdctlCmd, args...)

	//connect pipe to standard error when the command starts
	stderr, err := c.StderrPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//connect pipe to standard out when the command starts
	stdout, err := c.StdoutPipe()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	err = c.Start()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//parse the stderr if any
	errReader := bufio.NewReader(stderr)
	errStr, _ := errReader.ReadString('\n')

	if len(errStr) > 0 {
		a := &apperror.AppInfo{Msg: errors.New(strings.TrimSpace(errStr))}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//parse the stdout; read each line
	prefixCount := 0

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		prefixCount++
	}

	//get the number of KV pairs
	numOfKV := prefixCount / 2

	l.Info.Printf("%v prefix count: %v", fname, prefixCount)
	l.Info.Printf("%v number of KVs: %v", fname, numOfKV)

	if numOfKV == 0 {
		a := &apperror.AppInfo{Msg: errorEtcdServiceDidNotReturnExpectedPrefix}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		aeCh <- a
		return
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	hasCh <- true

}
