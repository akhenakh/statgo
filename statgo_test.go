package statgo

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHostInfo(t *testing.T) {
	s := NewStat()
	hi := s.HostInfo()
	assert.NotNil(t, s)
	assert.NotEmpty(t, hi.HostName, hi.OSName, hi.OSRelease, hi.OSVersion, hi.Platform)
	assert.True(t, hi.NCPUs > 0, hi.MaxCPUs > 0)

	t.Log(hi)
}

func TestCPU(t *testing.T) {
	s := NewStat()
	cpu := s.CPUStats()
	assert.NotNil(t, s)
	assert.NotNil(t, cpu)
	time.Sleep(100 * time.Millisecond)

	cpu = s.CPUStats()
	assert.False(t, math.IsNaN(cpu.User), math.IsNaN(cpu.Kernel), math.IsNaN(cpu.Idle))
	assert.False(t, math.IsNaN(cpu.LoadMin1), math.IsNaN(cpu.LoadMin5), math.IsNaN(cpu.LoadMin15))
	t.Log(cpu)
}

func TestFSInfos(t *testing.T) {
	s := NewStat()
	f := s.FSInfos()
	assert.True(t, len(f) > 0)

	for _, fs := range f {
		t.Log(fs)
	}
}

func TestInterfaceInfos(t *testing.T) {
	s := NewStat()
	interfaces := s.InteraceInfos()
	assert.True(t, len(interfaces) > 0)

	for _, i := range interfaces {
		t.Log(i)
	}
}

func TestVM(t *testing.T) {
	s := NewStat()
	m := s.MemStats()
	assert.NotNil(t, s)
	assert.NotNil(t, m)

	t.Log(m)
}
