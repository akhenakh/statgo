package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import "fmt"

// MemStats contains memory & swap stats
type MemStats struct {
	Total     int
	Free      int
	Used      int
	Cache     int
	SwapTotal int
	SwapUsed  int
	SwapFree  int
}

// MemStats get memory & swap related stats
// Go equivalent to sg_get_mem_stats & sg_get_swap_stats
func (s *Stat) MemStats() *MemStats {
	s.Lock()
	defer s.Unlock()

	mem_stats := C.sg_get_mem_stats(nil)
	swap_stats := C.sg_get_swap_stats(nil)

	m := &MemStats{
		Total:     int(mem_stats.total),
		Free:      int(mem_stats.free),
		Used:      int(mem_stats.used),
		Cache:     int(mem_stats.cache),
		SwapTotal: int(swap_stats.total),
		SwapUsed:  int(swap_stats.used),
		SwapFree:  int(swap_stats.free),
	}
	return m
}

func (m *MemStats) String() string {
	return fmt.Sprintf(
		"Total:\t%d\n"+
			"Free:\t\t%d\n"+
			"Used:\t\t%d\n"+
			"Cache:\t\t%d\n"+
			"SwapTotal:\t%d\n"+
			"SwapUsed:\t%d\n"+
			"SwapFree:\t%d\n",
		m.Total,
		m.Free,
		m.Used,
		m.Cache,
		m.SwapTotal,
		m.SwapUsed,
		m.SwapFree)
}
