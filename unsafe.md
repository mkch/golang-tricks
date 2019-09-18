# Unsafe Tricks

## Pointer as slice

```go
var bufSize = ...
var buf *T = ...

// CAUTION: this slice can't be used after buf is deallocated.
var sliceOfCBuf = (*[math.MaxUint32]T)(unsafe.Pointer(buf))[:bufSize]
```

or

```go
// If the buffer size is a compile time constant.
const bufSize = ...
var buf *T = ...

// CAUTION: this slice can't be used after buf is deallocated.
var sliceOfCBuf = (*[bufSize]T)(unsafe.Pointer(buf))[]
```

[sample](https://github.com/mkch/golang-tricks/blob/master/examples/unsafe/PtrSlice/main.go)
