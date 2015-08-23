package statgo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHostInfo(t *testing.T) {
	s := NewStat()
	hi := s.HostInfos()
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
	t.Log(cpu)
}

func TestCPULoad(t *testing.T) {
	s := NewStat()
	initialCPU := s.CPUStats()
	cpuCount := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuCount)
	var wg sync.WaitGroup
	wg.Add(cpuCount)
	doneChan := make(chan bool, cpuCount)
	for k := 0; k < cpuCount; k++ {
		go func() {
			defer wg.Done()
			var i uint64 = 2
			for {
				select {
				case <-doneChan:
					{
						return
					}
				default:
					{
					}
				}
				i = i * i
			}
		}()
	}
	s.CPUStats()
	testDuration := 5 * time.Second
	time.Sleep(testDuration)
	cpu := s.CPUStats()
	for k := 0; k < cpuCount; k++ {
		doneChan <- true
	}
	wg.Wait()

	//Assure that the Period of the stats is about the same as the Duration of the stats.
	assert.True(t, cpu.Period-testDuration < time.Second*2 || cpu.Period-testDuration > 2*time.Second)
	t.Log("CPU Idle %:", cpu.Idle)
	//The CPU should not be idle if we run cpuCount goroutines
	assert.True(t, cpu.Idle < 50.0)
	//The stats should have changed from the start till the finish of this test
	assert.True(t, cpu.Idle != initialCPU.Idle)
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

func TestDisksIO(t *testing.T) {
	s := NewStat()
	d := s.DiskIOStats()
	assert.NotNil(t, s)
	assert.NotNil(t, d)

	t.Log(d)
}

func TestNetIO(t *testing.T) {
	s := NewStat()
	n := s.NetIOStats()
	assert.NotNil(t, s)
	assert.NotNil(t, n)

	t.Log(n)
}

func getInterface(arr []*NetIOStats, names ...string) (*NetIOStats, error) {
	for _, ns := range arr {
		for _, name := range names {
			if ns.IntName == name {
				return ns, nil
			}
		}
	}
	return nil, fmt.Errorf("No interface not found matching any of %v", names)
}

func TestNetIOTXRX(t *testing.T) {
	s := NewStat()
	beforeNetIOArr := s.NetIOStats()
	beforeNetIO, err := getInterface(beforeNetIOArr, "lo", "lo0")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Log("Could not listen on port 8080", err)
		t.SkipNow()
		return
	}
	defer ln.Close()
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		io.Copy(ioutil.Discard, conn)
	}()
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Log("Could not connect to server:", err)
		t.SkipNow()
		return
	}
	defer conn.Close()
	conn.Write(make([]byte, 1024*1024))
	afterNetIOArr := s.NetIOStats()
	afterNetIO, err := getInterface(afterNetIOArr, "lo", "lo0")
	if err != nil {
		t.Log(err)
		t.SkipNow()
		return
	}
	t.Log("BeforeData:", beforeNetIO)
	t.Log("After 1MB data:", afterNetIO)
	assert.True(t, beforeNetIO.TX < afterNetIO.TX)
	assert.True(t, beforeNetIO.RX < afterNetIO.TX)
}

func TestProcess(t *testing.T) {
	s := NewStat()
	p := s.ProcessStats()
	assert.NotNil(t, s)
	assert.NotNil(t, p)

	t.Log(p)
}

func TestPages(t *testing.T) {
	s := NewStat()
	p := s.PageStats()
	assert.NotNil(t, s)
	assert.NotNil(t, p)

	t.Log(p)
}

func TestGoRoutineCleanup(t *testing.T) {
	var wg sync.WaitGroup

	s := NewStat()
	s.PageStats()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.free()
	}()

	wg.Wait()
}

func TestGoRoutines(t *testing.T) {
	// test for ticket #2
	// ping -s 20000 localhost, check for growing lo0 stats ([0] at least on OSX)

	t.Skip()
	var wg sync.WaitGroup
	s := NewStat()

	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("1", s.NetIOStats()[0])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2500 * time.Millisecond)
		t.Log("2", s.NetIOStats()[0])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		t.Log("4", s.NetIOStats()[0])
	}()
	wg.Wait()
}

func TestHostsInfos(t *testing.T) {

	s := NewStat()
	s.HostInfos()
	s.HostInfos()

}
