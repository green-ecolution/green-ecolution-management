# Smart City Flensburg - Grünflächenmanagement

- [Roadmap](https://scf-green-space.notion.site/ef9f26cd4a1e48ba9bac5156f72d604c?v=d326e873ae1c4cf8af00cfe2dfb4cc0d&pvs=4)

## Project structure

```
.
├── config       <- configuration files
│   ├── app.go
│   └── ...
├── internal
│   ├── entities <- domain entities (models)
│   ├── server   <- server setup (http, grpc, etc.)
│   ├── service  <- business logic (services)
│   └── storage  <- storage logic repository (database, cache, etc.)
└ main.go
```

## Technologies

- [Golang](https://go.dev/) as the main programming language
- [env](https://github.com/caarlos0/env) for environment variables
- [godotenv](https://github.com/joho/godotenv) for loading environment variables from a `.env` file
- [fiber](https://docs.gofiber.io/) for the web framework
- [testify](https://github.com/stretchr/testify) for testing

## Architecture

### Clean Architecture

The project is structured following the principles of the [Clean Architecture]. The main idea is to separate the business logic from the infrastructure and the framework. This way, the business logic is independent of the framework and can be easily tested. The framework is only used to connect the business logic with the outside world. The business logic is divided into three layers: entities, use cases, and interfaces. The entities layer contains the domain models. The use cases layer contains the business logic. The interfaces layer contains the interfaces that the use cases need to interact with the outside world. On top of it the project is structured following the principles of the [Layered Architecture].

Inside the `internal` folder, there are three main packages: `entities`, `service`, and `storage`. The `entities` package contains the domain models. The `service` package contains the business logic. The `storage` package contains the repository logic to interact with the database, cache, etc.

Inside the `internal` folder, there is a `server` package that contains the server setup. The server setup is responsible for setting up the server (http, grpc, etc.) and connecting the business logic with the outside world.

[Clean Architecture]: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
[Layered Architecture]: https://medium.com/@shershnev/layered-architecture-implementation-in-golang-6318a72c1e10

## Local development

### Requirements

- [Golang](https://go.dev/) as the main programming language
- [Air](https://github.com/air-verse/air) for live reload
- [Mockery](https://github.com/vektra/mockery) for mocking interfaces. Use version `v2.43.2`

### Setup

To enable live reload, you need to install [Air](https://github.com/air-verse/air). Air is a command-line utility for Go applications that monitors changes in the file system and restarts the application. To mock interfaces, you need to install [Mockery](https://github.com/vektra/mockery). Mockery is a tool for generating mocks for interfaces in Go. Inside the project folder, there is a `.env.example` file. You need to create a `.env` file with the same content and fill in the environment variables.

### Run

To run the project, you need to execute the following command:

**Use air for live reload**

```bash
air
```

**Without air**

```bash
go run main.go
```

### Test

Before running the tests, you need to create the mock files. To create the mock files, you need to execute the following command:

```bash
go install github.com/vektra/mockery/v2@v2.43.2 # install Mockery
mockery # create mock files
```

To run the tests, you need to execute the following command:

```bash
go test ./...
```

**NOTE:** Mockery is used to generate mocks for interfaces. The mocks are generated in the `_mocks` folder. To specify the output folder or to add created interfaces to the mocks, you can edit the `mockery.yml` file. The `mockery.yml` file is used to configure the behavior of Mockery. Running `go generate` will execute Mockery and generate the mocks. Also when running Air, the mocks will be generated automatically.
