package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
)

// CPUStats contains cpu stats
// Delivers correlated relative cpu counters (where  total is 100%)
type CPUStats struct {
	User   float64
	Kernel float64
	Idle   float64
	IOWait float64
	Swap   float64
	Nice   float64

	// System load averages
	LoadMin1  float64
	LoadMin5  float64
	LoadMin15 float64

	// The time taken in seconds since the last call of the function
	Period time.Duration

	TimeTaken time.Time
}

// CPUStats get cpu related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) CPUStats() *CPUStats {
	s.Lock()
	defer s.Unlock()

	var cpu *CPUStats
	do(func() {

		cpup := C.sg_get_cpu_percents_of(C.sg_new_diff_cpu_percent, nil)

		loadStat := C.sg_get_load_stats(nil)

		cpu = &CPUStats{
			User:      float64(cpup.user),
			Kernel:    float64(cpup.kernel),
			Idle:      float64(cpup.idle),
			IOWait:    float64(cpup.iowait),
			Swap:      float64(cpup.swap),
			Nice:      float64(cpup.nice),
			LoadMin1:  float64(loadStat.min1),
			LoadMin5:  float64(loadStat.min5),
			LoadMin15: float64(loadStat.min15),
			Period:    time.Duration(int(cpup.time_taken)) * time.Second,
			TimeTaken: time.Now(),
		}
	})
	return cpu
}

func (c *CPUStats) String() string {
	return fmt.Sprintf(
		"User:\t%f\n"+
			"Kernel:\t%f\n"+
			"Idle:\t%f\n"+
			"IOWait:\t%f\n"+
			"Swap:\t%f\n"+
			"Nice:\t%f\n"+
			"LoadMin1:\t%f\n"+
			"LoadMin5:\t%f\n"+
			"LoadMin15:\t%f\n"+
			"Period:\t%v\n"+
			"TimeTaken:\t%s\n",
		c.User,
		c.Kernel,
		c.Idle,
		c.IOWait,
		c.Swap,
		c.Nice,
		c.LoadMin1,
		c.LoadMin5,
		c.LoadMin15,
		c.Period,
		c.TimeTaken)
}
