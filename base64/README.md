### Base64

Base64 is binary to ASCII string conversion format. It provides flexibility when transfering data through wire as plain text. In golang;

- `b64.StdEncoding.EncodeToString([]byte(str))` is used to encode the incoming string into base64 string.

- `b, err := b64.StdEncoding.DecodeString(str)` is used to decode base64 string into original string.

- Similarly, in `gin` we can use `err = ioutil.WriteFile(input.Name, byteArray, 0777)` to write the incomming base64 string into file as shown by `HandleBase64FileUpload` handler function.

- Opposite i.e octet streaming also is supported by `gin` for this instead of `c.JSON` we use `c.Stream` which takes in `func(w io.Writer) bool` as its parameter and by using `w.Write([]byte)` we can serve the binary data to client. 

- Also, `b, err := ioutil.ReadFile("sample")` is used to read the saved file as byte array.