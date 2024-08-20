# olaf-layout-basic - Basic Layout

Olaf is an application scaffold based on Golang. It is a composition of various popular libraries from the Golang ecosystem, which can help you quickly build efficient and reliable applications.

![Olaf](https://github.com/zxylon/olaf/blob/main/.github/assets/banner.png)

## Documentation
* [User Guide](https://github.com/zxylon/olaf/blob/main/docs/en/guide.md)
* [Architecture](https://github.com/zxylon/olaf/blob/main/docs/en/architecture.md)
* [Getting Started Tutorial](https://github.com/zxylon/olaf/blob/main/docs/en/tutorial.md)

## Features
- **Gin**: https://github.com/gin-gonic/gin
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- **Testify**: https://github.com/stretchr/testify
- **Sonyflake**: https://github.com/sony/sonyflake
- **gocron**:  https://github.com/go-co-op/gocron
- More...

## Highlights
* **Low Learning Curve and Customization**: Olaf encapsulates popular libraries that Gophers are familiar with. You can easily customize the application to meet specific requirements.
* **High Performance and Scalability**: Olaf aims to be high-performance and scalable. It utilizes the latest technologies and best practices to ensure your application can handle high traffic and large data volumes.
* **Security and Reliability**: Olaf uses stable and reliable third-party libraries to ensure the security and reliability of your application.
* **Modularity and Extensibility**: Olaf is designed to be modular and extensible. You can easily add new features and functionalities by using third-party libraries or writing your own modules.
* **Comprehensive Documentation and Test Coverage**: Olaf has comprehensive documentation and test coverage. It provides detailed documentation and examples to help you get started quickly. It also includes a test suite to ensure your application works as expected.

## Olaf CLI

![Olaf](https://github.com/zxylon/olaf/blob/main/.github/assets/screenshot.jpg)

## Directory Structure
```
.
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── config
│   ├── local.yml
│   └── prod.yml
├── internal
│   ├── handler
│   │   ├── handler.go
│   │   └── user.go
│   ├── middleware
│   │   └── cors.go
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   ├── repository.go
│   │   └── user.go
│   ├── server
│   │   └── http.go
│   └── service
│       ├── service.go
│       └── user.go
├── pkg
├── LICENSE
├── README.md
├── README_zh.md
├── go.mod
└── go.sum

```

This is a classic directory structure for a Golang project, which includes the following directories:

- cmd: Contains the entry points of the application, including the main function and dependency injection code.
  - server: The main entry point of the application, including the main function and dependency injection code.
    - main.go: The main function used to start the application.
    - wire.go: The dependency injection code generated using Wire.
    - wire_gen.go: The dependency injection code generated using Wire.

- config: Contains the configuration files of the application.
  - local.yml: The configuration file for the local environment.
  - prod.yml: The configuration file for the production environment.

- internal: Contains the internal code of the application.
  - handler: Contains the handlers for handling HTTP requests.
    - handler.go: The common handler for handling HTTP requests.
    - user.go: The handler for handling user-related HTTP requests.
  - middleware: Contains the middleware code.
    - cors.go: The CORS (Cross-Origin Resource Sharing) middleware.
  - model: Contains the data model code.
    - user.go: The user data model.
  - repository: Contains the data access code.
    - repository.go: The common interface for data access.
    - user.go: The implementation of the user data access interface.
  - server: Contains the server code.
    - http.go: The implementation of the HTTP server.
  - service: Contains the business logic code.
    - service.go: The common interface for business logic.
    - user.go: The implementation of the user business logic.

- pkg: Contains the public packages of the application.
- storage: Contains the storage files of the application.
- go.mod: The Go module file.
- go.sum: The dependency versions file for the Go module.

## Requirements
To use Olaf, you need to have the following software installed on your system:

* Golang 1.16 or higher
* Git

### Installation

You can install Olaf using the following command:

```bash
go install github.com/zxylon/olaf@latest
```

### Creating a New Project

You can create a new Golang project using the following command:

```bash
olaf new projectName
```

By default, it will pull from the GitHub repository, but you can also use a mirror repository for faster access:

```
// Use the basic template
olaf new projectName -r https://gitee.com/zxylon/olaf-layout-basic.git
// Use the advanced template
olaf new projectName -r https://gitee.com/zxylon/olaf-layout-advanced.git
```

This command will create a directory named `projectName` and generate an elegant Golang project structure within it.

### Creating Components

You can create handlers, services, repositories, and models for your project using the following commands:

```bash
olaf create handler user
olaf create service user
olaf create repository user
olaf create model user
```
or
```
olaf create all user
```

These commands will create components named `UserHandler`, `UserService`, `UserDao`, and `UserModel` respectively and place them in the correct directories.

### Starting the Project

You can quickly start your project using the following command:

```bash
olaf run
```

This command will start your Golang project and support hot-reloading of files.

### Compiling wire.go

You can quickly compile `wire.go` using the following command:

```bash
olaf wire
```

This command will compile your `wire.go` file and generate the required dependencies.

## Contribution

If you find any issues or have any improvement suggestions, please feel free to raise an issue or submit a pull request. Your contributions are highly appreciated!

## License

Olaf is released under the MIT License. See the [LICENSE](LICENSE) file for more information.