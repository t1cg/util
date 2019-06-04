package runstat

import (
	//standard library

	"testing"
	"time"
)

func TestRMeasureRuntime(t *testing.T) {

	r := RunInfo{}
	r.Name = "runstat test"
	r.StartTime = time.Now()

	defer r.MeasureRuntime()

	time.Sleep(50 * time.Millisecond)

	r.SetEndTime(time.Now())

}

func TestMeasureRuntime(t *testing.T) {

	defer MeasureRuntime(time.Now())
	time.Sleep(50 * time.Millisecond)

}
