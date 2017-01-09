package runstat

import (
	//standard library
	"runtime"
	"testing"
	"time"
)

func TestMeasureRuntime(t *testing.T) {

	r := RunInfo{}
	r.Name = "runstat test"
	r.CPUCount = runtime.NumCPU()
	r.StartTime = time.Now()

	time.Sleep(50 * time.Millisecond)

	r.EndTime = time.Now()
	defer r.MeasureRuntime()
}
