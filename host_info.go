package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

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
		NCPUs:     int(stats.ncpus),
		MaxCPUs:   int(stats.maxcpus),
		BitWidth:  int(stats.bitwidth),
		//TODO: uptime
	}

	C.sg_free_stats_buf(unsafe.Pointer(stats))

	return hi
}

func (h *HostInfo) String() string {
	return fmt.Sprintf(
		"OSName:\t%s\n"+
			"OSRelease:\t%s\n"+
			"OSVersion:\t%s\n"+
			"Platform:\t%s\n"+
			"HostName:\t%s\n"+
			"NCPUs:\t\t%d\n"+
			"MaxCPUs:\t%d\n"+
			"BitWidth:\t%d\n",
		h.OSName,
		h.OSRelease,
		h.OSVersion,
		h.Platform,
		h.HostName,
		h.NCPUs,
		h.MaxCPUs,
		h.BitWidth)
}
