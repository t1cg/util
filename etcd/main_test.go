package etcd

import (
	"testing"
	//custom library
	"github.com/t1cg/util/apperror"
)

const (
	PREFIX = "mpp"
)

func TestGetPrefix(t *testing.T) {
	FUNCNAME := "TestGetPrefix()"

	t.Log(FUNCNAME + " calling...")

	kv := make(chan map[string]string)
	ae := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetPrefix(PREFIX, kv, ae)
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

func TestHasPrefix(t *testing.T) {
	FUNCNAME := "TestHasPrefix()"

	t.Log(FUNCNAME + " calling...")

	has := make(chan bool)
	ae := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.HasPrefix(PREFIX, has, ae)
	select {
	case a := <-ae:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case <-has:
		t.Log(FUNCNAME + " has prefix!")
	}

	t.Log(FUNCNAME + " complete")
}
