### Learning Go lang

Just my learning process for `golang`

Some of the articles I followed.
- [unit-testing-made-easy](https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318)
- [learning-to-use-go-reflection](https://medium.com/capital-one-tech/learning-to-use-go-reflection-part-2-c91657395066)

What I did:
- [x] explored go test and coverage.
- [x] Look at multiple input problem [math_test.go](/calc/math_test.go).
- [x] Separation of test using `calc -> code package` and `calc_test` as test package.
- [x] Reflection in golang `reflect` package. [reflection](/reflection/README.md)
- [x] Using reflection in golang [memoization](/reflection/memoization/README.md). 
- [x] A Look at go routines. [goroutines](/goroutines/README.md)
- [x] Async task like; sending email from go routines. using `gin framework`. `go SendEmail("Ola", "Ola Comoestas")` makes go routine and sends mail by using non-blocking way. [send-mail](/sendmail/README.md)
- [x] File upload in `go-gin` single and multiple. [fileupload](/fileupload/README.md)
- [x] About base64 encoding and decoding. [encoding-decoding](/base64/README.md)
- [x] File upload using base64 encoding. [handlers](/base64/encodingDecoding.go)
- [x] Build get request to struct parser using reflection. (exercise) [struct-parser](/structparser)
- [x] Dockerization of `go-gin` framework executables. (check note below about dockerization)
- [x] Look at Websocket Implementations in `golang`. [websockets](/websockets)

    Might need separate repository for fully implementing websockets.

- [x] Basic Hello World `grpc` app created. [gogrpc](/gogrpc)
- [ ] Build a simple console chat application using `grpc`.

- [x] Google bucket cloud storage in `gin`. [googlebucket](/googlebucket)


### Dockerization

Check [go.Dockerfile](/go.Dockerfile) which container execution steps when creating docker image. 

Command `docker-compose up -d` will create the container and `gin` app will run in `5000` port exposed. For different databases and other services we can just add another service to the `docker-compose.yml` file.
