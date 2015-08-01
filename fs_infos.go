package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// FSInfo contains filesystem & mountpoints informations
type FSInfo struct {
	DeviceName      string
	FSType          string
	MountPoint      string
	Size            int
	Used            int
	Free            int
	Available       int
	TotalInodes     int
	UsedInodes      int
	FreeInodes      int
	AvailableInodes int
}

// FSInfos return an FSInfo struct per mounted filesystem
// Go equivalent to sg_get_fs_stats
func (s *Stat) FSInfos() []*FSInfo {
	lock.Lock()
	defer lock.Unlock()
	var fs_size C.size_t
	var cArray *C.sg_fs_stats = C.sg_get_fs_stats(&fs_size)
	length := int(fs_size)
	slice := (*[1 << 30]C.sg_fs_stats)(unsafe.Pointer(cArray))[:length:length]

	var res []*FSInfo

	for _, v := range slice {
		f := &FSInfo{
			DeviceName:      C.GoString(v.device_name),
			FSType:          C.GoString(v.fs_type),
			MountPoint:      C.GoString(v.mnt_point),
			Size:            int(v.size),
			Used:            int(v.used),
			Free:            int(v.free),
			Available:       int(v.avail),
			TotalInodes:     int(v.total_inodes),
			UsedInodes:      int(v.used_inodes),
			FreeInodes:      int(v.free_inodes),
			AvailableInodes: int(v.avail_inodes),
		}
		res = append(res, f)
	}
	return res
}

func (fs *FSInfo) String() string {
	return fmt.Sprintf(
		"DeviceName:\t\t%s\n"+
			"FSType:\t\t%s\n"+
			"MountPoint:\t\t%s\n"+
			"Size:\t\t\t%d\n"+
			"Used:\t\t\t%d\n"+
			"Free:\t\t\t%d\n"+
			"Available:\t\t%d\n"+
			"TotalInodes:\t\t%d\n"+
			"UsedInodes:\t\t%d\n"+
			"FreeInodes:\t\t%d\n"+
			"AvailableInodes:\t%d\n",
		fs.DeviceName,
		fs.FSType,
		fs.MountPoint,
		fs.Size,
		fs.Used,
		fs.Free,
		fs.Available,
		fs.TotalInodes,
		fs.UsedInodes,
		fs.FreeInodes,
		fs.AvailableInodes)
}
