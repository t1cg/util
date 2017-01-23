package etcd

import (
	"testing"
	//custom library
	"github.com/t1cg/util/apperror"
)

const (
	testPrefix   = "mpp"
	testKey      = "mpp.api.endpoint.host"
	testKeyBlank = ""
)

func TestGetPrefix(t *testing.T) {
	FUNCNAME := "TestGetPrefix()"

	t.Log(FUNCNAME + " calling...")

	kv := make(chan map[string]string)
	ae := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetPrefix(testPrefix, kv, ae)
	select {
	case a := <-ae:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case kvs := <-kv:
		for k, v := range kvs {
			t.Logf("key[%s], value[%s]", k, v)
		}
	}

	t.Log(FUNCNAME + " complete")
}

func TestGetValue(t *testing.T) {
	FUNCNAME := "TestGetValue()"

	t.Log(FUNCNAME + " calling...")

	valueCh := make(chan string)
	aeCh := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetValue(testKeyBlank, valueCh, aeCh)
	select {
	case a := <-aeCh:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case value := <-valueCh:
		t.Logf("key[%s], value[%s]", testKey, value)
	}

	t.Log(FUNCNAME + " complete")
}

func TestHasPrefix(t *testing.T) {
	FUNCNAME := "TestHasPrefix()"

	t.Log(FUNCNAME + " calling...")

	hasCh := make(chan bool)
	aeCh := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.HasPrefix(testPrefix, hasCh, aeCh)
	select {
	case a := <-aeCh:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case <-hasCh:
		t.Log(FUNCNAME + " has prefix!")
	}

	t.Log(FUNCNAME + " complete")
}
