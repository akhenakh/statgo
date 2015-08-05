package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"runtime"
	"sync"
)

// Stat handle to access libstatgrab
type Stat struct {
	sync.Mutex
}

// NewStat return a new Stat handle
func NewStat() *Stat {
	s := &Stat{}
	runtime.SetFinalizer(s, (*Stat).free)
	C.sg_init(1)
	return s
}

func (s *Stat) free() {
	s.Lock()
	C.sg_shutdown()
	s.Unlock()
}
