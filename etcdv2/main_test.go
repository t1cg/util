package etcdv2

import (
	"testing"
	//t1cg library
	"github.com/t1cg/util/apperror"
)

func TestGetValue(t *testing.T) {
	FUNCNAME := "TestGetValue"

	t.Log(FUNCNAME + " Calling...")

	ec := ClientV3{}

	aeCh := make(chan *apperror.AppInfo)
	vCh := make(chan string)

	go ec.GetValue("mpp.loglevel", vCh, aeCh)

	select {
	case a := <-aeCh:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case value := <-vCh:
		t.Logf("value[%v]", value)
	}

	t.Log(FUNCNAME + " Complete")

}

func TestGetPrefix(t *testing.T) {
	FUNCNAME := "TestGetPrefix()"
	t.Log(FUNCNAME + " Calling...")

	ec := ClientV3{}

	prefix := "mpp"
	aeCh := make(chan *apperror.AppInfo)
	kvCh := make(chan map[string]string)

	go ec.GetPrefix(prefix, kvCh, aeCh)

	select {
	case a := <-aeCh:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case kvs := <-kvCh:
		for k, v := range kvs {
			t.Logf("key[%s], value[%s]", k, v)
		}
	}

}
