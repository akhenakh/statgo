package statgo

// #cgo LDFLAGS: -lstatgrab
// #include <statgrab.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// InterfaceInfos
type InterfaceInfos struct {
	Name string
	// In factor/sec
	Speed  int
	Factor int
	Duplex InterfaceDuplexType
	State  InterfaceState
}

// InterfaceDuplexType network interface duplex type
type InterfaceDuplexType int

const (
	InterfaceFullDuplex InterfaceDuplexType = iota
	InterfaceHalfDuplex
	InterfaceUnknownDuplex
)

var (
	interfaceDuplexTypes = []string{
		"Full Duplex", "Half Duplex", "Unknown Duplex",
	}
)

func (s InterfaceDuplexType) String() string {
	return interfaceDuplexTypes[s]
}

func (s InterfaceDuplexType) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *InterfaceDuplexType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := range interfaceDuplexTypes {
		if enum == interfaceDuplexTypes[i] {
			*s = InterfaceDuplexType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown status")
}

// InterfaceState up or down
type InterfaceState int

const (
	InterfaceUpState InterfaceState = iota
	InterfaceDownState
)

var (
	interfaceStates = []string{
		"UP", "DOWN",
	}
)

func (s InterfaceState) String() string {
	return interfaceStates[s]
}

func (s InterfaceState) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *InterfaceState) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := range interfaceStates {
		if enum == interfaceStates[i] {
			*s = InterfaceState(i)
			return nil
		}
	}
	return fmt.Errorf("unknown status")
}

func (s *Stat) InteraceInfos() []*InterfaceInfos {
	s.Lock()
	defer s.Unlock()
	var iface_count C.size_t
	var cArray *C.sg_network_iface_stats = C.sg_get_network_iface_stats(&iface_count)
	length := int(iface_count)
	slice := (*[1 << 16]C.sg_network_iface_stats)(unsafe.Pointer(cArray))[:length:length]

	var res []*InterfaceInfos

	for _, v := range slice {
		i := &InterfaceInfos{
			Name:   C.GoString(v.interface_name),
			Speed:  int(v.speed),
			Factor: int(v.factor),
		}
		switch v.duplex {
		case C.SG_IFACE_DUPLEX_FULL:
			i.Duplex = InterfaceFullDuplex
		case C.SG_IFACE_DUPLEX_HALF:
			i.Duplex = InterfaceHalfDuplex
		case C.SG_IFACE_DUPLEX_UNKNOWN:
			i.Duplex = InterfaceUnknownDuplex
		}

		switch v.up {
		case C.SG_IFACE_DOWN:
			i.State = InterfaceDownState
		case C.SG_IFACE_UP:
			i.State = InterfaceUpState
		}

		res = append(res, i)
	}
	return res
}

func (i *InterfaceInfos) String() string {
	return fmt.Sprintf(
		"Name:\t%s\n"+
			"Speed:\t\t%d\n"+
			"Factor:\t\t%d\n"+
			"Duplex:\t\t%s\n"+
			"State:\t\t%s\n",
		i.Name,
		i.Speed,
		i.Factor,
		i.Duplex,
		i.State)
}
