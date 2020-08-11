#### Context

Context carries deadlines, cancellation signales, and other request-scoped values across API boundaries and between processes.

`WithCancel`, `WithDeadline` and `WithTimeout` all take context i.e extends from already present context value. They all return a cancel function which cancels the child and its children. Also stops any asoociated times.

- Failing to cancel the context will leak the child and its children until the parent is canceled or the timer fires.

- Should not store contexts inside struct type; instead pass it explicitly to each function that needs it. The context should be the first parameter, typically named context.

  ```go
  func DoSomething(ctx context.Context, arguments Arg) err{
  // do someting
  }
  ```

Here, the `gen` function generates the integer until the context passed to the function is canceled.

```go
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dst)
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

```

For loop and select pattern is commonly used in go language on handling go contexts and preventing go routine leaks. This signaling can also be done without context and with channels as well. (Context package `Done` returns a channel which will be closed when close signal is received).

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ctx, cancel = context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Cancel the generation.")
		cancel()
	}()

```

In the example shown in `main.go` we can see that; when parent `withTimeout` cancels the context after `1 second` the context is canceled immediately even though child context is not canceled. In this way, go context provides signal propagation.
