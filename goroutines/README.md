### Go Routines

Similar to blocking and non-blocking in node js go lang comes with go routines to control asynchronous exection of the code. 

From nodejs documentation about [blocking-vs-non-blocking](https://nodejs.org/en/docs/guides/blocking-vs-non-blocking/)
> Blocking is when the execution of additional JavaScript in the Node.js process must wait until a non-JavaScript operation completes. This happens because the event loop is unable to continue running JavaScript while a blocking operation is occurring.

Go routines are functions or methods that run concurrently with other functions or methods. They can be thought as light weight threads. The cost of creating goroutines is tiny when compared with created thread.

starting go routines is easier just by using `go` keyword ahead of function calls. Using this keyword control doesnot wait for the go routines to finish executing. Control returns immediately to the next line of code after the call and return of values are ignored.

### Notes

- The output of `TestAdd2ValuesAfter2Seconds` is as follows;
    ```
    === RUN   TestAdd2ValuesAfter2Seconds
    Value after addition is:  6
        TestAdd2ValuesAfter2Seconds: firstgoroutine_test.go:14: 1. Value after adding 2 values is 6 took 2.000375917s
        TestAdd2ValuesAfter2Seconds: firstgoroutine_test.go:19: 1. Value after adding 2 values is 6 took 44.152µs
    --- PASS: TestAdd2ValuesAfter2Seconds (2.00s)
    ```

    Here, we know that first exection waited for the function to be executed but second execution didn,t wait for program to close. 

    Second function call was not executed because for a go routine to be executed `main go routine` also need to be running. i.e. the program stopped running and the sub go routines was just stopped without execution. (This point is important when comparing threads with go routines, threads actually create another process in the os where as go routines are run in the same process as main program so when the main program terminates sub program `goroutines` also get terminated) 