[![wercker status](https://app.wercker.com/status/c56e26bf18be587114d26764a7a0ce7a/s/master "wercker status")](https://app.wercker.com/project/byKey/c56e26bf18be587114d26764a7a0ce7a) [![](https://godoc.org/github.com/akhenakh/statgo?status.png)](http://godoc.org/github.com/akhenakh/statgo) 

StatGo
======
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
s.HostInfos()
OSName: Darwin
OSRelease:  14.4.0
OSVersion:  Darwin Kernel Version 14.4.0: Thu May 28 11:35:04 PDT 2015; root:xnu-2782.30.5~1/RELEASE_X86_64
Platform:   x86_64
HostName:   kamoulox
NCPUs:      4
MaxCPUs:    4
BitWidth:   64

s.CPUStats()
User:       7.500000
Kernel:     2.500000
Idle:       90.000000
IOWait      0.000000
Swap:       0.000000
Nice:       0.000000
LoadMin1:   2.206055
LoadMin5:   2.031250
LoadMin15:  1.970703

s.FSInfos()[0]
DeviceName:         /dev/disk1
FSType:             hfs
MountPoint:         /
Size:               249769230336
Used:               224367140864
Free:               25402089472
Available:          25139945472
TotalInodes:        60978814
UsedInodes:         54841132
FreeInodes:         6137682
AvailableInodes:    6137682

s.InterfaceInfos()[0]
Name:   en2
Speed:  0
Factor: 1000000
Duplex: Full Duplex
State:  UP

s.MemStats()
Total:      16649420800
Free:       4323848192
Used:       12325572608
Cache:      0
SwapTotal:  3221225472
SwapUsed:   2528378880
SwapFree:   692846592

s.NetIOStats()
IntName:    en0
TX:         2310272606
RX:         3336240203
IPackets:   114473581
OPackets:   129430304
IErrors:    0
OErrors:    0
Collisions: 0

s.ProcessStats()
Total:      343
Running:    335
Sleeping:   0
Stopped:    0
Zombie:     8

s.PagesStats()
PageIn:     90173695
PageOut:    90173695
```

### Status

- [x]  Host infos
- [x]  cpu stats
- [x]  load average
- [x]  network interfaces infos
- [x]  mem stats
- [x]  swap stat 
- [x]  io stats
- [x]  net io stats
- [x]  process count
- [x]  page stats

### Contributors
* [HeinOldewage](https://github.com/HeinOldewage)
