package etcd

import (
	"testing"
	//custom library
	"github.com/t1cg/util/apperror"
)

const (
	testPrefix     = "mpp"
	pre2           = "mpp_2"
	pre2Val        = "mppval_2"
	pre3           = "mpp_3"
	pre3Val        = "mppval_3"
	pre4           = "mpp_4"
	pre4Val        = "mppval_4"
	testKey        = "mpp.api.endpoint.host"
	testValue      = "testKeyValue"
	testPrefixFail = "fail"
	testKeyFail    = "different"
	testKeyBlank   = "blank"
)

func TestNotStarted(t *testing.T) {
	FUNCNAME := "TestNotStarted"

	t.Log(FUNCNAME + " calling...")

	kv := make(chan map[string]string)
	ae := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetPrefix(testPrefix, kv, ae)
	select {
	case a := <-ae:
		t.Log(FUNCNAME+" EXPECTED ERROR  ETCD NOT RUNNING: ", a.Msg)
		go startService()
		go putValue(testKey, testValue)
		go putValue(pre2, pre2Val)
		go putValue(testKeyBlank, testKeyFail)
		go putValue(pre3, pre3Val)
		go putValue(pre4, pre4Val)
		return
	case kvs := <-kv:
		for k, v := range kvs {
			t.Errorf("ERROR: etcd already running: key[%s], value[%s]", k, v)

		}
	}

	t.Log(FUNCNAME + " complete")
}

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

func TestGetPrefixFail(t *testing.T) {
	FUNCNAME := "TestGetPrefixFail()"

	t.Log(FUNCNAME + " calling...")

	kv := make(chan map[string]string)
	ae := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetPrefix(testPrefixFail, kv, ae)
	select {
	case a := <-ae:
		t.Log(FUNCNAME+" EXPECTED ERROR:", a.Msg)
		return
	case kvs := <-kv:
		for k, v := range kvs {
			t.Errorf("No keys expected, Returned: key[%s], value[%s]", k, v)
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

	go ec.GetValue(testKey, valueCh, aeCh)
	select {
	case a := <-aeCh:
		t.Error(FUNCNAME+" ERROR:", a.Msg)
		return
	case value := <-valueCh:
		t.Logf("key[%s], value[%s]", testKey, value)
	}

	t.Log(FUNCNAME + " complete")
}

func TestGetValueFail(t *testing.T) {
	FUNCNAME := "TestGetValueFail()"

	t.Log(FUNCNAME + " calling...")

	valueCh := make(chan string)
	aeCh := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetValue(testPrefixFail, valueCh, aeCh)
	select {
	case a := <-aeCh:
		t.Log(FUNCNAME+"  EXPECTED ERROR: ", a.Msg)
		return
	case value := <-valueCh:
		t.Errorf("No key expected Returned: key[%s], value[%s]", testKey, value)
	}

	t.Log(FUNCNAME + " complete")
}

func TestGetValueFailPair(t *testing.T) {
	FUNCNAME := "TestGetValueFailPair()"

	t.Log(FUNCNAME + " calling...")

	valueCh := make(chan string)
	aeCh := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.GetValue(testPrefixFail, valueCh, aeCh)
	select {
	case a := <-aeCh:
		t.Log(FUNCNAME+"  EXPECTED ERROR: ", a.Msg)
		return
	case value := <-valueCh:
		t.Errorf("No key expected Returned: key[%s], value[%s]", testKey, value)
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

func TestHasPrefixFail(t *testing.T) {
	FUNCNAME := "TestHasPrefix()"

	t.Log(FUNCNAME + " calling...")

	hasCh := make(chan bool)
	aeCh := make(chan *apperror.AppInfo)

	ec := Cml{}

	go ec.HasPrefix(testPrefixFail, hasCh, aeCh)
	select {
	case a := <-aeCh:
		t.Log(FUNCNAME+" EXPECTED ERROR:", a.Msg)
		return
	case <-hasCh:
		t.Error(FUNCNAME + " has UNEXPECTED prefix!")
	}

	t.Log(FUNCNAME + " complete")
}

func TestStop(t *testing.T) {
	FUNCNAME := "TestStop()"
	t.Log(FUNCNAME + " calling...")

	stopService()

	t.Log(FUNCNAME + " complete")
}
