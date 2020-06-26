### Result of `TestActualMemoization`:
```
called function.
got result 3 in 100.390394ms
got result 3 in 40.589µs
got result 3 in 5.603µs
got result 3 in 3.984µs
got result 3 in 3.974µs
called function.
got result  3 in 100.294405ms

```

### After running `TestActualMemoization` test what we can note is;

- during the for loop first call of function took 100ms where as consequetive calls took lesser time (significantly lesser).

- First call printed `called function` string means first time it was called and then it didnot print for other times.

- During last function call the memoization was expired and it called the function again taking 100ms.