# ChanelTricks

Channel is a great feature in go.

## Send with timeout

```go
var ch = make(chan T)
const timeout time.Duration = ...

timer := time.NewTimer(timeout)
select {
case ch <- value:
    // Sent successfully.
case <-timer.C:
    // Timed out.
}
timer.Stop()
```

[sample](https://github.com/mkch/golang-tricks/blob/master/examples/channel/SendWithTimeout/main.go) | [run](https://play.golang.org/p/f9tq5AtBXxs)

## Drop on backpressure

The sender drops items unless there is a pending receiving.

```go
var ch = make(chan T)

select {
case ch <- item:
    // Sent successfully.
default:
    // Drop.
}
```

[sample](https://github.com/mkch/golang-tricks/blob/master/examples/channel/OnBackpressureDrop/main.go) | [run](https://play.golang.org/p/2Viq2a2hRlT)

## Keep latest on backpressure

The sender holds on to the most-recently sent item and makes it available to its receiver upon request. Drops any other items that it sends between requests from its receiver.

```go
var ch = make(chan T, 1) // Buffer 1

select {
case <-ch:
    // Discard old value if any.
default:
    // Nop if no old value.
}
// Send the latest.
ch <- value
```

[sample](https://github.com/mkch/golang-tricks/blob/master/examples/channel/OnBackpressureLatest/main.go) | [run](https://play.golang.org/p/ecRX-x-zC_9)
