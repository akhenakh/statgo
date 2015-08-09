package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import "fmt"

// MemStats contains memory & swap stats
// expressed in bytes
type MemStats struct {
	// The total amount of real memory in bytes
	Total int

	// Theount of real memory in bytes.
	Free int

	// The used amount of real memory in bytes
	Used int

	// The amount of real memory in bytes used for caching
	Cache int

	// The total swap space in bytes.
	SwapTotal int

	// The used swap in bytes
	SwapUsed int

	// The free swap in bytes
	SwapFree int
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
