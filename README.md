StatGo
======

WORK IN PROGRESS

[libstatgrab](http://www.i-scream.org/libstatgrab/) bindings for Golang.

### Compilation 
You need at least libstatgrab 0.91, Debian & Ubuntu only have 0.90 ...

On Linux, OSX & FreeBSD, you can simply install libstatgrab it with the usual
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
fmt.Println(hi.OSName)
FreeBSD
```