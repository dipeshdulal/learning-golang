### Testing in Go lang

Following articles
- [unit-testing-made-easy](https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318)
- [learning-to-use-go-reflection](https://medium.com/capital-one-tech/learning-to-use-go-reflection-part-2-c91657395066)

What I did:
- [x] explored go test and coverage.
- [x] Look at multiple input problem [math_test.go](/calc/math_test.go).
- [x] Separation of test using `calc -> code package` and `calc_test` as test package.
- [x] Reflection in golang `reflect` package. [reflection](/reflection/README.md)
- [x] Using reflection in golang [memoization](/reflection/memoization/README.md). 
- [x] A Look at go routines.
- [x] Async task like; sending email from go routines. using `gin framework`. `go SendEmail("Ola", "Ola Comoestas")` makes go routine and sends mail by using non-blocking way. [send-mail](/sendmail/README.md)