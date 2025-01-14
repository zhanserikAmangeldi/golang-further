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
