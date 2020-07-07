### Select operation in channels

The select statement selects between multiple go routine operations. And runs the one that is available. 

For example; 
- worker1 is available every 200ms where as 
- worker2 is available every 400ms 

without select statement it would be;
- if we wait for worker1 then worker2 will never run because program will wait for worker 1 always. if we take out from w1 first.

So using select statement; we select between multiple channel output. If one channel output is resolved then the code under that case statement runs if another, another one runs.

If multiple channels are ready, then select chooses one at random. 
