### MemoryLeaks in go routines

Example taken out from [here](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html)

There are two implementation of go routines provided in the example `main.go`
Running the example produces:
```
3 random ints.
1: 5577006791947779410
2: 8674665223082153551
3: 6129484611666145821
3 random ints.
0: 3916589616287113937
1: 6334824724549167320
2: 605394647632969758
newRandStreamWithoutMemoryLeak closure exited.
```

From this example we can see that the go routine closure inside `newRandStreamWithMemoryLeak` never exited and is waiting for main go routine to take out variable from channel. This indicates that there is memory leak in out go routines. 

Similarly, the proof that the closure inside `newRandStreamWithoutMemoryLeak` exited is the printed statement closure exited. In the function without memory leak. We are adding one responsibility to the go routine that creates a go routine i.e it being responsible for closing that go routines that it creates. 

> If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.


This pattern can be used to ensure that go routines don't leak and is being properly cleaned up. 

The convention can help ensure us that programs will not create memory leaks. `pipelies` or `context package` also works in similar way i.e passing in the `done` channel. 