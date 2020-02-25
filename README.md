# TMESS 
## Terminal Messenger
(In development)

Simple client and server to integrate messaging in a terminal environment

- Build with goa V3, a design-first approach to building microservices
- Includes Server Executable, CLI, OpenAPI
- GoLang 

If the design of the application is updated (`design.go`), you must rerun code generation.

This will recreate the `/gen` dir with the appropriate updates. 
```
goa gen tmess/design
```

Goa can also generate example implementations for both the service and client
```
goa example calcsvc/design

calc.go
cmd/calc-cli/http.go
cmd/calc-cli/main.go
cmd/calc/http.go
cmd/calc/main.go
```
#### Run and compile
```
cd cmd/tmess-ci
go build
```
