### Reflection

Following [learning to use go reflection](https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7)

When we want to work with variables at runting using information that 
didn't exist when the program was written. Or maybe we need to dynamically
map data into variables for test cases. In these situations we need to to use reflection.

It Allows to examine, modify and create variables, functions and structs at runtime.

- `t := reflect.TypeOf(var)` to find the type of variable in runtime. Returns `Reflect.Type` type
- In addition to examining types of variables, we can use it to read, set or create values. Providing value in `ValueOf` means we are only reading the values. And providing pointer in `ValueOf` means we are changing the underlying variable.

- Most common use of golang is marshaling and unmarshaling data from a file or network. We could also build memoization and short-term caching using reflection.


#### Notes
- Memoization is a process of creating function that wraps the functions and remembers the input and output of the program so that work is only done once per set of input values. For same input values output values should not change (functional programming principle) so for this return values are pulled from cache rather than recomputed. For functions that do complex or slow things, the performance saving can be tremendous. Memoization are also used in microservices architecture to avoid extra network calls.

- There is performance penalty for generated functions or middle layer which is caused by reflection. Although the performance penaltity doesnot seem that much we still need to decide if the added functionality is worth the slower performance and more complicated code.

- When there is a problem that contains dynamic data types and cannot know the data types before hand then the reflection can be used.

