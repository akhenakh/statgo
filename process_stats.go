package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import "fmt"

// ProcessStats contains processes count stats
type ProcessStats struct {
	// The total number of processes
	Total int

	// The number of running processes
	Running int

	// The number of sleeping processes
	Sleeping int

	// The number of stopped processes
	Stopped int

	// The number of zombie processes
	Zombie int
}

// ProcessStats get prceosses related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) ProcessStats() *ProcessStats {
	s.Lock()
	defer s.Unlock()

	// Throw away the first reading as thats averaged over the machines uptime
	pstat := C.sg_get_process_count_of(C.sg_entire_process_count)

	p := &ProcessStats{
		Total:    int(pstat.total),
		Running:  int(pstat.running),
		Sleeping: int(pstat.sleeping),
		Stopped:  int(pstat.stopped),
		Zombie:   int(pstat.zombie),
	}
	return p
}

func (p *ProcessStats) String() string {
	return fmt.Sprintf(
		"Total:\t\t%d\n"+
			"Running:\t%d\n"+
			"Sleeping:\t%d\n"+
			"Stopped:\t%d\n"+
			"Zombie:\t\t%d\n",
		p.Total,
		p.Running,
		p.Sleeping,
		p.Stopped,
		p.Zombie)
}
