### Struct Parser

This is a exercise of using reflections in golang. What this completes is that given the get request 

```
localhost:5000/handle-request-parsing?id=12&name=Dipesh Dulal&json={ "name": "dipesh", "id": 64 } 
```

The handler parses the get request into struct after getting the key from tags present in struct in golang. `here -> getParser:"name"`. The output from the above code look like;

```json
{
    "data": {
        "name": "Dipesh Dulal",
        "id": 12,
        "json": {
            "name": "dipesh",
            "id": 64
        }
    },
    "message": "succesfully bind to getRequestInput"
}

```

`key := field.Tag.Get("getParser")` is used to get the tags value provided in the struct. 

We can look at the code snippet below to know how we are checking the types defined in the golang struct and then using different logic to ser the value accordingly.

```go
switch kind {

case reflect.Int:
    // set value in case of int
    break

case reflect.String:
    // set value incase of string
    break

case reflect.Struct:
    // set value in case of struct
    break
}
```