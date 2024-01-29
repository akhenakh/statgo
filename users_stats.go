package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// UserStats contains user stats
// expressed in bytes
type UserStats struct {
	// The username which was used to log in
	LoginName string

	// Record identifier of host database containing login information (not necessarily 0-terminated)
	RecordId string

	// Size of the record identifier
	recordIdSize int

	// Device name (eg. "pts/0") of the tty assigned to the login session
	Device string

	// (remote) Hostname from where the user is logged on, eg. "infoterm7.some.kind.of.domain.local", "localhost", "10.42.17.4" or ":0.0" (in case it's a local logon via new xterm)
	Hostname string

	// Process identifier of the process which made the entry to the logged on users database
	pid int

	// Timestamp (time in seconds since epoch) when the user logged on
	LoginTime time.Duration

	// The timestamp when the above stats where collected in seconds since epoch
	SysTime time.Time
}

// UserStats return an UserStats list
func (s *Stat) UserStats() []*UserStats {
	s.Lock()
	defer s.Unlock()
	var userSize C.size_t
	var cArray *C.sg_user_stats = C.sg_get_user_stats(&userSize)
	length := int(userSize)
	slice := (*[1 << 16]C.sg_user_stats)(unsafe.Pointer(cArray))[:length:length]

	var res []*UserStats

	for _, v := range slice {
		f := &UserStats{
			LoginName: C.GoString(v.login_name),
			RecordId:  C.GoString(v.record_id),
			Device:    C.GoString(v.device),
			Hostname:  C.GoString(v.hostname),
			LoginTime: time.Duration(int(v.login_time)) * time.Second,
		}
		res = append(res, f)
	}
	return res
}

func (m *UserStats) String() string {
	return fmt.Sprintf(
		"LoginName:\t%s\n"+
			"RecordId:\t\t%d\n"+
			"Device:\t\t%d\n"+
			"Hostname:\t\t%d\n"+
			"LoginTime:\t%d\n",
		m.LoginName,
		m.RecordId,
		m.Device,
		m.Hostname,
		m.LoginTime)
}
