package runstat

import (
	"runtime"
	"strconv"
	"strings"
	"time"
	//project library
	"github.com/t1cg/util/logger"
)

//global
var (
	CPUCount int
)

func init() {
	CPUCount = runtime.NumCPU()
}

// MeasureRuntimer interface ...
type MeasureRuntimer interface {
	MeasureRuntime()
	SetEndTime(time.Time)
}

// LoadAsyncInfo defines a model for a CPU's runtime
type LoadAsyncInfo struct {
	Name                    string
	CPUCount                int
	LineCounter             int64
	ReadCounter             int64
	TransformCounter        int64
	WriteCounter            int64
	TransformChannelRef     int64
	WriteChannelRef         int64
	BulkQueueRef            int64
	BulkQueuesInsertCounter int64
	BulkQueueCounter        int64
	BulkRunCounter          int64
	StartTime               time.Time
	EndTime                 time.Time
	ReadStartTime           time.Time
	ReadEndTime             time.Time
	TransformStartTime      time.Time
	TransformEndTime        time.Time
	WriteStartTime          time.Time
	WriteEndTime            time.Time
}

//test to ensure proper interface is implemented; otherwise, will not compile
var _ MeasureRuntimer = (*LoadAsyncInfo)(nil)

// MeasureRuntime prints a string containing the runtime information
func (r *LoadAsyncInfo) MeasureRuntime() {
	logger.L.Perf.Printf("%v runtime[%v] on cpus[%d]", r.Name, r.EndTime.Sub(r.StartTime), r.CPUCount)

	out := "\n*****************************************************************\n"
	out = out + "name:.....................: " + r.Name + "\n"
	out = out + "# of cpus:................: " + strconv.Itoa(r.CPUCount) + "\n"
	out = out + "file line count:..........: " + strconv.FormatInt(r.LineCounter, 10) + "\n"
	out = out + "read count:...............: " + strconv.FormatInt(r.ReadCounter, 10) + "\n"
	out = out + "transform count:..........: " + strconv.FormatInt(r.TransformCounter, 10) + "\n"
	out = out + "write count:..............: " + strconv.FormatInt(r.WriteCounter, 10) + "\n"
	out = out + "*****************************************************************\n"
	out = out + "transform channel queue:..: " + strconv.FormatInt(r.TransformChannelRef, 10) + "\n"
	out = out + "write channel queue:......: " + strconv.FormatInt(r.WriteChannelRef, 10) + "\n"
	out = out + "bulk queue:...............: " + strconv.FormatInt(r.BulkQueueRef, 10) + "\n"
	out = out + "*****************************************************************\n"
	out = out + "run start time:...........: " + r.StartTime.String() + "\n"
	out = out + "run end time:.............: " + r.EndTime.String() + "\n"
	out = out + "total run time:...........: " + r.EndTime.Sub(r.StartTime).String() + "\n"
	out = out + "read start time:..........: " + r.ReadStartTime.String() + "\n"
	out = out + "read end time:............: " + r.ReadEndTime.String() + "\n"
	out = out + "total read time:..........: " + r.ReadEndTime.Sub(r.ReadStartTime).String() + "\n"
	out = out + "transform start time:.....: " + r.TransformStartTime.String() + "\n"
	out = out + "transform end time:.......: " + r.TransformEndTime.String() + "\n"
	out = out + "total transform time:.....: " + r.TransformEndTime.Sub(r.TransformStartTime).String() + "\n"
	out = out + "write start time:.........: " + r.WriteStartTime.String() + "\n"
	out = out + "write end time:...........: " + r.WriteEndTime.String() + "\n"
	out = out + "total write time:.........: " + r.WriteEndTime.Sub(r.WriteStartTime).String() + "\n"
	out = out + "*****************************************************************\n"

	logger.L.Info.Println(out)

}

// SetEndTime function sets the end time of the run.
func (r *LoadAsyncInfo) SetEndTime(t time.Time) {
	r.EndTime = t
}

// RunInfo defines a model for a CPU's runtime
type RunInfo struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

//test to ensure proper interface is implemented; otherwise, will not compile
var _ MeasureRuntimer = (*RunInfo)(nil)

// MeasureRuntime prints a string containing the runtime information
func (r *RunInfo) MeasureRuntime() {
	logger.L.Perf.Printf("%v runtime[%v] on cpus[%d] ", strings.TrimSpace(r.Name), r.EndTime.Sub(r.StartTime), CPUCount)
}

// SetEndTime function sets the end time of the run.
func (r *RunInfo) SetEndTime(t time.Time) {
	r.EndTime = t
}
