package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// DiskIOStats contains disk io stats
type DiskIOStat struct {
	DiskName   string
	ReadBytes  int
	WriteBytes int
}

// CPUStats get cpu related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) DiskIOStats() []*DiskIOStat {
	lock.Lock()
	defer lock.Unlock()

	var num_diskio_stats C.size_t
	var cArray *C.sg_disk_io_stats = C.sg_get_disk_io_stats_diff(&num_diskio_stats)
	length := int(num_diskio_stats)
	slice := (*[1 << 16]C.sg_disk_io_stats)(unsafe.Pointer(cArray))[:length:length]

	var res []*DiskIOStat

	for _, v := range slice {
		f := &DiskIOStat{
			DiskName:   C.GoString(v.disk_name),
			ReadBytes:  int(v.read_bytes),
			WriteBytes: int(v.write_bytes),
		}
		res = append(res, f)
	}
	return res

}

func (d *DiskIOStat) String() string {
	return fmt.Sprintf(
		"DiskName:\t%s\n"+
			"ReadBytes:\t%d\n"+
			"WriteBytes:\t%d\n",
		d.DiskName,
		d.ReadBytes,
		d.WriteBytes)
}
