package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"runtime"
	"sync"
)

var lock sync.Mutex

// Stat handle to access libstatgrab
type Stat struct {
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

func free(s *Stat) {
	lock.Lock()
	C.sg_shutdown()
	lock.Unlock()
}
