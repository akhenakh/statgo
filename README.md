[![wercker status](https://app.wercker.com/status/c56e26bf18be587114d26764a7a0ce7a/m "wercker status")](https://app.wercker.com/project/bykey/c56e26bf18be587114d26764a7a0ce7a)

StatGo
======

WORK IN PROGRESS


StatGo give you access to OS metrics like network interface bandwith, cpus usage ...  
It supports FreeBSD, Linux, OSX & more, it's in fact a [libstatgrab](http://www.i-scream.org/libstatgrab/) bindings for Golang.


### Compilation 
You need at least libstatgrab 0.91, Debian & Ubuntu only have 0.90 ...

On Linux, OSX & FreeBSD, you can simply install libstatgrab with the usual commands:
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
fmt.Println(cpu)
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
DeviceName: /dev/disk1
FSType: hfs
MountPoint: /
Size:   249769230336
Used:   224467705856
Free:   25301524480
Available:  25039380480
TotalInodes:    60978814
UsedInodes: 54865684
FreeInodes: 6113130
AvailableInodes:    6113130
```