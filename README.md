[![wercker status](https://app.wercker.com/status/c56e26bf18be587114d26764a7a0ce7a/m "wercker status")](https://app.wercker.com/project/bykey/c56e26bf18be587114d26764a7a0ce7a)

StatGo
======

WORK IN PROGRESS
- [x]  Host infos
- [x]  cpu stats
- [x]  load average
- [x]  network interfaces infos
- [x]  mem stats
- [x]  swap stat 
- [x]  io stats
- [ ]  net io stats
- [ ]  process count
- [ ]  page stats

StatGo give you access to OS metrics like network interface bandwith, cpus usage ...  
It supports FreeBSD, Linux, OSX & more, it's in fact a [libstatgrab](http://www.i-scream.org/libstatgrab/) binding for Golang.  
Tested on FreeBSD, OSX, Linux amd64, Linux arm.


### Compilation 
You need at least libstatgrab 0.91, Debian & Ubuntu only have 0.90 ...

On Debian/Ubunt & OSX, you can simply install libstatgrab with the usual commands:
```
./configure --prefix=/usr/local
make
sudo make install
```

You may have to set CGO_LDFLAGS and CGO_CFLAGS environment according to your path:
```
export CGO_CFLAGS=-I/usr/local/include
export CGO_LDFLAGS=-L/usr/local/lib
```

Note: On OSX you need to install gcc to access cgo.

    go get github.com/akhenakh/statgo

### Usage
```
s := NewStat()
hi := s.HostInfo()
fmt.Println(hi)
OSName: Darwin
OSRelease:  14.4.0
OSVersion:  Darwin Kernel Version 14.4.0: Thu May 28 11:35:04 PDT 2015; root:xnu-2782.30.5~1/RELEASE_X86_64
Platform:   x86_64
HostName:   kamoulox
NCPUs:      4
MaxCPUs:    4
BitWidth:   64

cpu := s.CPUStats()
User:       7.500000
Kernel:     2.500000
Idle:       90.000000
IOWait      0.000000
Swap:       0.000000
Nice:       0.000000
LoadMin1:   2.206055
LoadMin5:   2.031250
LoadMin15:  1.970703

f := s.FSInfos()
fmt.Println(f[0])
DeviceName:     /dev/disk1
 FSType:         hfs
MountPoint:     /
Size:           249769230336
Used:           224248410112
Free:           25520820224
Available:      25258676224
TotalInodes:        60978814
UsedInodes:     54812145
FreeInodes:     6166669
AvailableInodes:    6166669

interfaces := s.InteraceInfos()
fmt.Println(interfaces[0])
Name:   en2
Speed:  0
Factor: 1000000
Duplex: Full Duplex
State:  UP

m := s.MemStats()
Total:      16649420800
Free:       4323848192
Used:       12325572608
Cache:      0
SwapTotal:  3221225472
SwapUsed:   2528378880
SwapFree:   692846592
```