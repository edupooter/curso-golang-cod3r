# curso-golang-cod3r
Coded the course examples:

[Go (Golang) - Explorando a Linguagem do Google](https://www.cod3r.com.br/courses/go-golang-explorando-a-linguagem-do-google)

To run all the tests:

`$ go test ./...`

To generate test coverage report:

`$ go test -coverprofile=coverage.out`

`$ go tool cover -func=coverage.out`

`$ go tool cover -html=coverage.out`
