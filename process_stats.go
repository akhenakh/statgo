package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import "fmt"

// ProcessStat contains processes count stats
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

// CPUStats get cpu related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) ProcessStats() *ProcessStats {
	s.Lock()
	defer s.Unlock()

	// Throw away the first reading as thats averaged over the machines uptime
	p_stat := C.sg_get_process_count_of(C.sg_entire_process_count)

	p := &ProcessStats{
		Total:    int(p_stat.total),
		Running:  int(p_stat.running),
		Sleeping: int(p_stat.sleeping),
		Stopped:  int(p_stat.stopped),
		Zombie:   int(p_stat.zombie),
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
