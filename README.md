# Todo App in Golang
A todo application in Golang, developed using MVC Architecture

## Dependencies
```
> Go: v1.18+
> Postgres v14.8+
```
## How to run

### Clone the Repository
```
clone the repository using `git clone https://github.com/anxmukul/todo-golang.git`

open the repo using `cd todo-golang`
```
### Set env variable
You have to provide DB credentials to connect to postgres database, this application takes this credentials as environment variable.

Here I am going to show you how to set env variable in Ubuntu

```
You have to Provide 5 env variable for DB in key-value pair

$export host="localhost"
$export user="User"
$export password=""
$export database=""
$export port=""
```

### How to run
```
> run `go mod tidy` to get all the packages and modules

> run `go run main.go` to start the application
```

