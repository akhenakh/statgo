package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"runtime"
	"sync"
	"time"
	"unsafe"
)

var lock sync.Mutex

// Stat handle to access libstatgrab
type Stat struct {
}

// HostInfo contains the Go equivalent to sg_host_info
type HostInfo struct {
	OSName    string
	OSRelease string
	OSVersion string
	Platform  string
	HostName  string
	NCPUs     int
	MaxCPUs   int
	BitWidth  int //32, 64 bits
	uptime    *time.Time
	systime   *time.Time
}

//
func NewStat() *Stat {
	s := &Stat{}
	runtime.SetFinalizer(s, free)

	lock.Lock()
	C.sg_init(1)
	lock.Unlock()
	return s
}

func (s *Stat) HostInfo() *HostInfo {
	lock.Lock()
	stats := C.sg_get_host_info(nil)

	hi := &HostInfo{
		OSName:    C.GoString(stats.os_name),
		OSRelease: C.GoString(stats.os_release),
		OSVersion: C.GoString(stats.os_version),
		Platform:  C.GoString(stats.platform),
		HostName:  C.GoString(stats.hostname),
		NCPUs:     int(C.uint(stats.ncpus)),
		MaxCPUs:   int(C.uint(stats.maxcpus)),
		BitWidth:  int(C.uint(stats.bitwidth)),
	}

	C.sg_free_stats_buf(unsafe.Pointer(stats))
	lock.Unlock()

	return hi
}

func free(s *Stat) {
	lock.Lock()
	defer lock.Unlock()
	C.sg_shutdown()
}
