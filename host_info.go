package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
)

// HostInfo contains informations related to the system
type HostInfos struct {
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
func (s *Stat) HostInfos() *HostInfos {
	s.Lock()
	stats := C.sg_get_host_info(nil)
	s.Unlock()

	hi := &HostInfos{
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

	return hi
}

func (h *HostInfos) String() string {
	return fmt.Sprintf(
		"OSName:\t\t%s\n"+
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
