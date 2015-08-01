package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
)

// CPUStats contains cpu stats
type CPUStats struct {
	User      float64
	Kernel    float64
	Idle      float64
	IOWait    float64
	Swap      float64
	Nice      float64
	LoadMin1  float64
	LoadMin5  float64
	LoadMin15 float64
	timeTaken time.Time
}

// CPUStats get cpu related stats
// note that 1st call to 100ms may return NaN as values
// Go equivalent to sg_cpu_percents
func (s *Stat) CPUStats() *CPUStats {
	lock.Lock()
	defer lock.Unlock()

	// Throw away the first reading as thats averaged over the machines uptime
	C.sg_snapshot()
	C.sg_get_cpu_stats_diff(nil)

	cpu_percent := C.sg_get_cpu_percents_of(C.sg_last_diff_cpu_percent, nil)

	C.sg_snapshot()

	load_stat := C.sg_get_load_stats(nil)

	cpu := &CPUStats{
		User:      float64(cpu_percent.user),
		Kernel:    float64(cpu_percent.kernel),
		Idle:      float64(cpu_percent.idle),
		IOWait:    float64(cpu_percent.iowait),
		Swap:      float64(cpu_percent.swap),
		Nice:      float64(cpu_percent.nice),
		LoadMin1:  float64(load_stat.min1),
		LoadMin5:  float64(load_stat.min5),
		LoadMin15: float64(load_stat.min15),
		//TODO: timetaken
	}
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
			"LoadMin15:\t%f\n",
		c.User,
		c.Kernel,
		c.Idle,
		c.IOWait,
		c.Swap,
		c.Nice,
		c.LoadMin1,
		c.LoadMin5,
		c.LoadMin15)
}
