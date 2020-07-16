### Future Pattern

A future indicates that any data that is needed in future but its computation can be started in parallel so that it can be fetched from background when needed. In go lang, channel concurrency primitives can be used to build up the functionality. Mostly futures can be used to send asynchronous http requests. 

For this we first create a function that returns a channel. The function will execute another go routine in parallel and the routine will push data to the channel. The pushed data is being listened by `main` func. 

```go
func getData(url) chan data {
    c := make(chan data)

    // go routine
    go func(){
        // expensive process 
        c <- d // push d (result of expensive process) into channel.
    }()

    return c
}

```

While in `main` function. We are taking out data from returned channel and then printing it.

```go
func main(){
    future := getData(url)

    // do other things that doesn't require future data 

    // takeout from future data.
    body := <-future 

    // process the future data.
}
```
