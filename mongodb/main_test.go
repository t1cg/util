package mongodb

import (
	//standard library
	"testing"

	//t1cg library
	"github.com/t1cg/util/logger"
)

// TestMongoDB function tests the functions for the mongodb package.
func TestMongoDB(t *testing.T) {
	t.Logf("running...")

	//get the function name
	fname := logger.GetFuncName()

	if len(fname) == 0 {
		t.Errorf("caller did not return the expected string, %v", fname)
		return
	}

	t.Logf("%v: call expected to fail", fname)
	ci := ConnectionInfo{Host: "", DBName: "mpp", User: "mppReader", Pw: "iamamppreader", To: 15}
	_, ae := GetSession(ci)
	if ae != nil {
		t.Logf("%v: call failed; contineu", fname)
	}

	t.Logf("%v: call expected to fail due to incorrect connection string", fname)
	ci = ConnectionInfo{Host: "localhost:27017", DBName: "mpp", User: "mppReader", Pw: "xxxxx", To: 0}
	_, ae = GetSession(ci)
	if ae != nil {
		t.Logf("%v: call failed, %v", fname, ae.Error())
	}

	t.Logf("%v: call expected to pass with correct connection string", fname)
	ci = ConnectionInfo{Host: "localhost:27017", DBName: "mpp", User: "mppReader", Pw: "iamamppreader", To: 15}
	cn, ae := GetSession(ci)
	if ae != nil {
		t.Errorf("%v: call failed, %v", fname, ae.Error())
		return
	}

	t.Logf("%v: return object should not be nil", fname)
	if cn == nil {
		t.Errorf("%v: return object nil, %v", fname, ae.Error())
		return
	}

	t.Logf("exiting")
}
