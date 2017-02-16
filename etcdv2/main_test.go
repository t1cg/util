package etcdv2

import (
	"testing"
	//t1cg library
	"github.com/t1cg/util/apperror"
)

func TestServiceRunning(t *testing.T) {
	FUNCNAME := "TestServiceRuning"

	t.Log(FUNCNAME + " Calling...")

	go startService()

	t.Log(FUNCNAME + " Complete")

}

func TestPutTheValue(t *testing.T) {
	FUNCNAME := "TestPutTheValue"

	t.Log(FUNCNAME + " Calling...")

	keys := []string{"mpp", "mpp2", "mpp3", "test", "key"}
	values := []string{"mppvalue", "mppvalue2", "mppvalue3", "testvalue", "keyvalue"}
	ae := make(chan *apperror.AppInfo)

	go PutTheValue(keys, values, ae)
}

func TestGetTheValue(t *testing.T) {
	FUNCNAME := "TestGetTheValue"

	t.Log(FUNCNAME + " Calling...")

	ae := make(chan *apperror.AppInfo)
	kv := make(chan map[string]string)

	go GetTheValue("test", kv, ae)

	select {
	case a := <-ae:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case kvs := <-kv:
		for k, v := range kvs {
			t.Logf("key:[%s] value:[%s]", k, v)
		}
	}

	t.Log(FUNCNAME + " Complete")

}

//  TestGetTheValueFail is currently not working, gets an index out of range and panics, need to look into
// func TestGetTheValueFail(t *testing.T) {
// 	FUNCNAME := "TestGetTheValueFail"

// 	t.Log(FUNCNAME + " Calling...")

// 	ae := make(chan *apperror.AppInfo)
// 	kv := make(chan map[string]string)

// 	go GetTheValue("failTest", "failure", kv, ae)

// 	select {
// 	case a := <-ae:
// 		t.Log(FUNCNAME+" EXPECTED ERROR", a.Msg)
// 		return
// 	case kvs := <-kv:
// 		for k, v := range kvs {
// 			t.Errorf(FUNCNAME+" Unexpected key value FAIL"+"key[%s], value[%s]", k, v)
// 		}
// 	}
// 	t.Log(FUNCNAME + " Complete")
// }

func TestGetThePrefix(t *testing.T) {
	FUNCNAME := "TestGetThePrefix()"
	t.Log(FUNCNAME + " Calling...")

	prefix := "mpp"
	ae := make(chan *apperror.AppInfo)
	kv := make(chan map[string]string, 10)

	go GetThePrefix(prefix, kv, ae)

	select {
	case a := <-ae:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case kvs := <-kv:
		for k, v := range kvs {
			t.Logf("key:[%s] value:[%s]\n", k, v)
		}
	}
	t.Log(FUNCNAME + " complete")
}

func TestStopService(t *testing.T) {
	FUNCNAME := "TestStopService()"
	t.Log(FUNCNAME + " Calling...")

	stopService()

	t.Log(FUNCNAME + " Complete")
}
