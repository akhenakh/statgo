package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import "fmt"

// PageStats contains pages stats
type PageStats struct {
	// The number of pages swapped into memory
	PageIn int

	// The number of pages swapped out of memory (to swap)
	PageOut int
}

// PageStats get pages related stats
// Go equivalent to sg_get_page_stats_diff
func (s *Stat) PageStats() *PageStats {
	s.Lock()
	defer s.Unlock()

	page_stats := C.sg_get_page_stats_diff(nil)

	p := &PageStats{
		PageIn:  int(page_stats.pages_pagein),
		PageOut: int(page_stats.pages_pageout),
	}
	return p
}

func (p *PageStats) String() string {
	return fmt.Sprintf(
		"PageIn:\t%d\n"+
			"PageOut:\t%d\n",
		p.PageIn,
		p.PageOut)
}
