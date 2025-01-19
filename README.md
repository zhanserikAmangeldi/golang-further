# Let's Go Further - Greenlight

There you can find my way of learning golang using the book "Let's Go Further" by Alex Edward. In the end we should get application called Greenlight for retrieving and managing information about movies.

## Getting Started

In this section of the book, we're going to set up a project directory and lay the groundwork for building our Greenlight API. We will:

- Create a skeleton directory structure.
- Establish a HTTP server.
- Introduce a sensible patter for managing configuration settings(via commandline flags) and using dependency injection.
- Use the `httprouter` package.

### Project Setup

For creating golang module, you should use command:

```bash
$ go mod init [name_the_module]
go: creating new go.mod: module [name_the_module]
```

After that we can create the skeleton directory structure of our project:

```bash
$ mkdir -p bin cmd/api internal migrations remote
$ touch Makefile
$ touch cmd/api/main.go
```

You should get structure like that:

```bash
.
├── Makefile
├── README.md
├── bin
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── internal
├── migrations
└── remote
```

You can create project with any structure, but it is structure use all good practices in industry. All directory and files have their purpose.

- The bin directory will contain our complied binaries.
- The cmd/api directory will contain the application the application-specific code for our application.
- The internal directory will contain various ancillary packages used by our API.
- The migrations directory will contain the SQL migration files for our database.
- The remote directory will contain the configuration fields and setup scripts for our production server.
- The go.mod file will declare our project dependencies.
- The Makefile will contain recipes for automating common administrative tasks.

For checking that everything is working correct, try run `Hell world!` app. 

```go
# File: cmd/api/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

and run the command to get result:

```bash
$ go run ./cmd/api
Hello world!
```

### A Basic HTTP Server

To start with, we'l configure our server to have just one endpoint: `/v1/healthcheck`. This endpoint will return some basic information about our API.

|URL Pattern|Handler|Action|
|-----------|-------|------|
|/v1/healthcheck|healthcheckHandler|Show application information|

I will not describe whole new code, but you can understand by reading comments. In this step we add code to start up our server and the first handler to check health of our application. 
New:
    - ./cmd/api/healthcheck.go
Updated:
    - ./cmd/api/main.go

After add new lines, you can run and check the new endpoint `/v1/healthcheck`.

```bash
$ go run ./cmd/api
2025/01/14 12:10:53 starting development server on :4000
```

And check the endpoint using curl(or ui solutions):

```bash
$ curl localhost:4000/v1/healthcheck
status: available
environment: development
version: 1.0.0
```

NOTE: you can use the flags that we add like that `go run ./cmd/api -port=8080 -env=production`

### API Endpoints and RESgithub.com/julienschmidt/httprouter v1.3.0Tful Routing

Before we just add simple handler to check the health, but now we add some basis endpoints for our application, like:

Method | URL Pattern | Handler | Action
-------|-------------|---------|-------
GET | /v1/healthcheck | healthcheckHandler | Show application information
GET | /v1/movies | listMoviesHandler | Show the details of all movies
POST | /v1/movies | createMovieHandler | Create a new movie
GET | /v1/movies/:id | showMovieHandler | Show the details of a specific movie
PUT | /v1/movies/:id | editMovieHandler | Update the details of a specific movie
DELETE | /v1/movies/:id | deleteMovieHandler | Delete a specific movie

And for routing we will be using third-party router - `httprouter`. To download that router you should write:

```bash
$ go get github.com/julienschmidt/httprouter
go: downloading github.com/julienschmidt/httprouter v1.3.0
go: added github.com/julienschmidt/httprouter v1.3.0
```

## Setup Database  

In this project we will use postgresql. 

### Create the database

First we must create new database to our project and connect to the new database:

```
postgres=# create database greenlight;
CREATE DATABASE
postgres=# \c greenlight
You are now connected to database "greenlight" as user "zhans".
greenlight=# 
```

### Create new role

Second we must create the new role to manage greenlight database:

```
greenlight=# CREATE ROLE greenlight WITH LOGIN PASSWORD 'password';
CREATE ROLE
greenlight=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

### Connect as the new user

Connect to the database as the new user:

```
$ psql --host=localhost --dbname=greenlight --username=greenlight
psql (14.15 (Homebrew))
Type "help" for help.

greenlight=> SELECT current_user;
 current_user 
--------------
 greenlight
(1 row)

greenlight=> 
```
