package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

var lock sync.Mutex

// Stat handle to access libstatgrab
type Stat struct {
	cpu_percent *C.sg_cpu_percents
}

// HostInfo contains informations related to the system
type HostInfo struct {
	OSName    string
	OSRelease string
	OSVersion string
	Platform  string
	HostName  string
	NCPUs     int
	MaxCPUs   int
	BitWidth  int //32, 64 bits
	uptime    time.Time
	systime   time.Time
}

// CPUMetrics
type CPUMetrics struct {
	User      float64
	Kernel    float64
	Idle      float64
	IOWait    float64
	Swap      float64
	Nice      float64
	timeTaken time.Time
	LoadMin1  float64
	LoadMin5  float64
	LoadMin15 float64
}

// NewStat return a new Stat handle
func NewStat() *Stat {
	s := &Stat{}
	runtime.SetFinalizer(s, free)

	lock.Lock()
	C.sg_init(1)
	lock.Unlock()
	return s
}

// HostInfo get the host informations
// Go equivalent to sg_host_info
func (s *Stat) HostInfo() *HostInfo {
	lock.Lock()
	stats := C.sg_get_host_info(nil)
	lock.Unlock()

	hi := &HostInfo{
		OSName:    C.GoString(stats.os_name),
		OSRelease: C.GoString(stats.os_release),
		OSVersion: C.GoString(stats.os_version),
		Platform:  C.GoString(stats.platform),
		HostName:  C.GoString(stats.hostname),
		NCPUs:     int(C.uint(stats.ncpus)),
		MaxCPUs:   int(C.uint(stats.maxcpus)),
		BitWidth:  int(C.uint(stats.bitwidth)),
		//TODO: uptime
	}

	C.sg_free_stats_buf(unsafe.Pointer(stats))

	return hi
}

// CPU returns a CPUMetrics object
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) CPU() *CPUMetrics {
	lock.Lock()
	defer lock.Unlock()
	// Throw away the first reading as thats averaged over the machines uptime
	C.sg_snapshot()
	C.sg_get_cpu_stats_diff(nil)

	s.cpu_percent = C.sg_get_cpu_percents_of(C.sg_last_diff_cpu_percent, nil)

	C.sg_snapshot()

	load_stat := C.sg_get_load_stats(nil)

	cpu := &CPUMetrics{
		User:      float64(C.double(s.cpu_percent.user)),
		Kernel:    float64(C.double(s.cpu_percent.kernel)),
		Idle:      float64(C.double(s.cpu_percent.idle)),
		IOWait:    float64(C.double(s.cpu_percent.iowait)),
		Swap:      float64(C.double(s.cpu_percent.swap)),
		Nice:      float64(C.double(s.cpu_percent.nice)),
		LoadMin1:  float64(C.double(load_stat.min1)),
		LoadMin5:  float64(C.double(load_stat.min5)),
		LoadMin15: float64(C.double(load_stat.min15)),
		//TODO: timetaken
	}

	return cpu
}

func free(s *Stat) {
	lock.Lock()
	C.sg_shutdown()
	lock.Unlock()
}

func (c *CPUMetrics) String() string {
	return fmt.Sprintf(
		"User:\t\t%f\n"+
			"Kernel:\t\t%f\n"+
			"Idle:\t\t%f\n"+
			"IOWait\t\t%f\n"+
			"Swap:\t\t%f\n"+
			"Nice:\t\t%f\n"+
			"LoadMin1:\t%f\n"+
			"LoadMin5:\t%f\n"+
			"LoadMin15:\t%f\n",
		c.User,
		c.Kernel,
		c.Idle,
		c.IOWait,
		c.Swap,
		c.Nice,
		c.LoadMin1,
		c.LoadMin5,
		c.LoadMin15)
}
