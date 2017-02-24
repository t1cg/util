package etcdv2

import (
	"context"
	"errors"
	"os"
	"time"
	//t1cg library
	"github.com/t1cg/util/apperror"
	"github.com/t1cg/util/logger"
	"github.com/t1cg/util/runstat"
	//coreos etcd library
	"github.com/coreos/etcd/clientv3"
)

// EtcdConstant struct defines the constants to be used.
var EtcdConstant = struct {
	EnvEndPoint       string
	ErrEnvVarNotFound error
	ErrExpectedOneKV  error
	TO                int
}{
	EnvEndPoint:       "ETCD_ENDPOINT",
	ErrEnvVarNotFound: errors.New("Expected environment variable not found"),
	ErrExpectedOneKV:  errors.New("Expected etcd to return one key/value pair"),
	TO:                5,
}

// global
var (
	endpoint string
)

func init() {

	//get environment variable
	endpoint = os.Getenv(EtcdConstant.EnvEndPoint)
	if len(endpoint) == 0 {
		a := &apperror.AppInfo{Msg: EtcdConstant.ErrEnvVarNotFound}
		a.LogError(a.Error())
		os.Exit(1)
	}

}

// ClientV3 struct ...
type ClientV3 struct{}

// GetValue function returns a single key/value pair if found. If not found, or more
// than one key/value pair is found, an error is returned.
func (c ClientV3) GetValue(key string, valueCh chan string, aeCh chan *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(EtcdConstant.TO) * time.Millisecond,
	})
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(EtcdConstant.TO)*time.Millisecond)

	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}

	var value string

	if len(resp.Kvs) != 1 {
		a := &apperror.AppInfo{Msg: EtcdConstant.ErrExpectedOneKV}
		a.LogError(a.Error())
		aeCh <- a
		return
	}

	for _, ev := range resp.Kvs {
		value = string(ev.Value)
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	valueCh <- value
}

// GetPrefix function returns all the key/value pairs that matches the provided key.
func (c ClientV3) GetPrefix(prefix string, valueCh chan map[string]string, aeCh chan *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(EtcdConstant.TO) * time.Millisecond,
	})
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}

	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(EtcdConstant.TO)*time.Millisecond)

	resp, err := cli.Get(ctx, prefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}

	var finalKv = make(map[string]string)

	for _, ev := range resp.Kvs {
		finalKv[string(ev.Key)] = string(ev.Value)
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	valueCh <- finalKv

}
