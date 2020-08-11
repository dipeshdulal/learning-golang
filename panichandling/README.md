#### Paninc Handling

From [here](https://medium.com/swlh/simple-guide-to-panic-handling-and-recovery-in-golang-72d6181ae3e8)

When a function panics its execution is stopped, any deferred function are executed and then the function returns to its caller.

When we defer a function, we are guaranteeing its exection at the end of the surrounding function.

When a function panics, it's execution is stopped and any deferred functions are executed.

- `RECOVER` is a function provided by Go which takes control of panicking go-routine.
- `recover()` can only be used inside deferred function.
- If you call recover during a normal flow, it will simply return null but, if it is panicking. `recover` will capture the panic value.

```go
func printAllOperations(x, y int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from error: %v \n", r)
		}
	}()

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf(
		"sum = %v, sub = %v, mul = %v, div = %v \n",
		sum,
		subtract,
		multiply,
		divide,
	)

}
```

- Shows how we handle error (divide by zero error) with `recover()` statement inside defer function.

- This method is useful in logging panic data and providing graceful exit but we can also use panic recovery to provide alternative logic flows when one flow panics.

- In above scenario only division can cause panic so, let us make such that division is not called when panic is occurred.

```go
func printSkipDivide(x, y int) {
	sum, subtract, multiply := x+y, x-y, x*y
	fmt.Printf(
		"sum = %v, sub = %v, mul = %v \n",
		sum,
		subtract,
		multiply,
	)
}
```

- Call this function in case of defer statement.