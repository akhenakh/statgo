package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// DiskIOStats contains disk io stats
// Expressed in bytes
type DiskIOStats struct {
	DiskName   string
	ReadBytes  int
	WriteBytes int

	// The time period over which read_bytes and write_bytes were transferred.
	Period time.Duration

	TimeTaken time.Time
}

// CPUStats get cpu related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_disk_io_stats
func (s *Stat) DiskIOStats() []*DiskIOStats {
	s.Lock()
	defer s.Unlock()

	var res []*DiskIOStats
	do(func() {
		var num_diskio_stats C.size_t
		var cArray *C.sg_disk_io_stats = C.sg_get_disk_io_stats_diff(&num_diskio_stats)
		length := int(num_diskio_stats)
		slice := (*[1 << 16]C.sg_disk_io_stats)(unsafe.Pointer(cArray))[:length:length]

		for _, v := range slice {
			f := &DiskIOStats{
				DiskName:   C.GoString(v.disk_name),
				ReadBytes:  int(v.read_bytes),
				WriteBytes: int(v.write_bytes),
				Period:     time.Duration(int(v.systime)) * time.Second,
				TimeTaken:  time.Now(),
			}
			res = append(res, f)
		}
	})

	return res

}

func (d *DiskIOStats) String() string {
	return fmt.Sprintf(
		"DiskName:\t%s\n"+
			"ReadBytes:\t%d\n"+
			"WriteBytes:\t%d\n"+
			"Period:\t%v\n"+
			"TimeTaken:\t%s\n",
		d.DiskName,
		d.ReadBytes,
		d.WriteBytes,
		d.Period,
		d.TimeTaken)
}
