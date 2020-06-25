### Reflection

Following [learning to use go reflection](https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7)

When we want to work with variables at runting using information that 
didn't exist when the program was written. Or maybe we need to dynamically
map data into variables for test cases. In these situations we need to to use reflection.

It Allows to examine, modify and create variables, functions and structs at runtime.

- `t := reflect.TypeOf(var)` to find the type of variable in runtime. Returns `Reflect.Type` type
- In addition to examining types of variables, we can use it to read, set or create values. Providing value in `ValueOf` means we are only reading the values. And providing pointer in `ValueOf` means we are changing the underlying variable.
