# Unsafe Tricks

## Access C pointer as slice

```go
// Get from C code somewhere.
var bufSize = ...
var buf *C.T = ...

// Wraps the C buffer to a go slice.
// C.T and T must be compatible(i.e. same size, same memory layout ……)
// CAUTION: this slice can't be used after buf is freed(C.free)
var sliceOfCBuf = (*[math.MaxUint32]T)(unsafe.Pointer(buf))[:bufSize]
```

or

```go
// If the buffer size is a compile time constant.
const bufSize = ...
var buf *C.T = ...

var sliceOfCBuf = (*[bufSize]T)(unsafe.Pointer(buf))[]
```

[sample](https://github.com/mkch/golang-tricks/blob/master/examples/unsafe/PtrSlice/main.go)
