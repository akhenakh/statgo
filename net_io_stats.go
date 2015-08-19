package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// NetIOStat contains network interfaces stats
type NetIOStats struct {
	IntName string

	// Number of bytes transmitted
	TX int

	// Number of bytes received
	RX int

	// Number of packets received
	IPackets int

	// Number of packets transmitted
	OPackets int

	// Number of receive errors
	IErrors int

	// Number of transmit errors
	OErrors int

	// Number of collisions count
	Collisions int

	// the time period over which tx and rx were transferred
	Period time.Duration

	TimeTaken time.Time
}

// NetIOStats get interface ios related stats
// Go equivalent to sg_get_network_io_stats
func (s *Stat) NetIOStats() []*NetIOStats {
	s.Lock()
	defer s.Unlock()

	var res []*NetIOStats

	do(func() {
		var num_network_stats C.size_t
		var cArray *C.sg_network_io_stats = C.sg_get_network_io_stats_diff(&num_network_stats)
		length := int(num_network_stats)
		slice := (*[1 << 16]C.sg_network_io_stats)(unsafe.Pointer(cArray))[:length:length]

		for _, v := range slice {
			n := &NetIOStats{
				IntName:    C.GoString(v.interface_name),
				TX:         int(v.tx),
				RX:         int(v.rx),
				IPackets:   int(v.ipackets),
				OPackets:   int(v.opackets),
				IErrors:    int(v.ierrors),
				OErrors:    int(v.oerrors),
				Collisions: int(v.collisions),
				Period:     time.Duration(int(v.systime)) * time.Second,
				TimeTaken:  time.Now(),
			}

			res = append(res, n)
		}
	})
	return res
}

func (n *NetIOStats) String() string {
	return fmt.Sprintf(
		"IntName:\t%s\n"+
			"TX:\t\t%d\n"+
			"RX:\t\t%d\n"+
			"IPackets:\t%d\n"+
			"OPackets:\t%d\n"+
			"IErrors:\t%d\n"+
			"OErrors:\t%d\n"+
			"Collisions:\t%d\n"+
			"Period:\t%v\n"+
			"TimeTaken:\t%s\n",
		n.IntName,
		n.TX,
		n.RX,
		n.IPackets,
		n.OPackets,
		n.IErrors,
		n.OErrors,
		n.Collisions,
		n.Period,
		n.TimeTaken)
}
