package statgo

import (
	"log"
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

	log.Println(hi)
}

func TestCPU(t *testing.T) {
	s := NewStat()
	cpu := s.CPU()
	assert.NotNil(t, s)
	assert.NotNil(t, cpu)
	time.Sleep(100 * time.Millisecond)

	cpu = s.CPU()
	assert.False(t, math.IsNaN(cpu.User), math.IsNaN(cpu.Kernel), math.IsNaN(cpu.Idle))
	assert.False(t, math.IsNaN(cpu.LoadMin1), math.IsNaN(cpu.LoadMin5), math.IsNaN(cpu.LoadMin15))
	log.Println(cpu)
}
