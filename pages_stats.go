package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
)

// PageStats contains pages stats
type PageStats struct {
	// The number of pages swapped into memory.
	PageIn int

	// The number of pages swapped out of memory (to swap).
	PageOut int

	// The time period over which pages_pagein and pages_pageout weretransferred.
	Period time.Duration

	TimeTaken time.Time
}

// PageStats get pages related stats
// Go equivalent to sg_get_page_stats_diff
func (s *Stat) PageStats() *PageStats {
	s.Lock()
	defer s.Unlock()

	var p *PageStats

	do(func() {
		page_stats := C.sg_get_page_stats_diff(nil)

		p = &PageStats{
			PageIn:    int(page_stats.pages_pagein),
			PageOut:   int(page_stats.pages_pageout),
			Period:    time.Duration(int(page_stats.systime)),
			TimeTaken: time.Now(),
		}
	})
	return p
}

func (p *PageStats) String() string {
	return fmt.Sprintf(
		"PageIn:\t%d\n"+
			"PageOut:\t%d\n"+
			"Period:\t%v\n"+
			"TimeTaken:\t%s\n",
		p.PageIn,
		p.PageOut,
		p.Period,
		p.TimeTaken)
}
