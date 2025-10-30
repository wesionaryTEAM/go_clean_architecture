# Go Clean Architecture

Clean Architecture with [Gin Web Framework](https://github.com/gin-gonic/gin)

## Features :star:

-   Clean Architecture written in Go
-   Application backbone with [Gin Web Framework](https://github.com/gin-gonic/gin)
-   Dependency injection using [uber-go/fx](https://pkg.go.dev/go.uber.org/fx)
-   Uses fully featured [GORM](https://gorm.io/index.html)

## Linter setup

Need [Python3](https://www.python.org/) to setup linter in git pre-commit hook.

```zsh
make lint-setup
```

---

## Run application

-   Setup environment variables

```zsh
cp .env.example .env
```

-   Update your database credentials environment variables in `.env` file
-   Update `STORAGE_BUCKET_NAME` in `.env` with your AWS S3 bucket name.

### Locally

-   Run `go run main.go app:serve` to start the server.
-   There are other commands available as well. You can run `go run main.go -help` to know about other commands available.

### Using `Docker`

> Ensure Docker is already installed in the machine.

-   Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if there are permission issues.

---

## Folder Structure :file_folder:

| Folder Path                      | Description                                                                                            |
| -------------------------------- | ------------------------------------------------------------------------------------------------------ |
| `/bootstrap`                     | Contains modules required to start the application.                                                    |
| `/console`                       | Server commands; run `go run main.go -help` for all available commands.                                |
| `/docker`                        | Docker files required for `docker-compose`.                                                            |
| `/docs`                          | Contains project documentation.                                                                        |
| `/domain`                        | Contains models, constants, and a folder for each domain with controller, repository, routes, and services. |
| `/domain/constants`              | Global application constants.                                                                          |
| `/domain/models`                 | ORM models.                                                                                            |
| `/domain/<name>`                 | Controller, repository, routes, and service for a domain (e.g., `user` is a domain in this template).  |
| `/hooks`                         | Git hooks.                                                                                             |
| `/migrations`                    | Database migration files managed by Atlas.                                                             |
| `/pkg`                           | Contains shared packages for errors, framework utilities, infrastructure, middlewares, responses, services, types, and utils. |
| `/pkg/errorz`                    | Defines custom error types and handlers for the application.                                           |
| `/pkg/framework`                 | Core framework components like environment variable parsing, logger setup, etc.                        |
| `/pkg/infrastructure`            | Setup for third-party service connections (e.g., AWS, database, router).                               |
| `/pkg/middlewares`               | HTTP request middlewares used in the application.                                                        |
| `/pkg/responses`                 | Defines standardized HTTP response structures and error handling.                                        |
| `/pkg/services`                  | Shared application services or clients for external services (e.g., Cognito, S3, SES).                 |
| `/pkg/types`                     | Custom data types used throughout the application.                                                       |
| `/pkg/utils`                     | Global utility and helper functions.                                                                   |
| `/seeds`                         | Seed data for database tables.                                                                         |
| `/tests`                         | Application tests (unit, integration, etc.).                                                           |
| `.env.example`                   | sample environment variables                                                                           |
| `docker-compose.yml`             | `docker compose` file for service application via `Docker`                                             |
| `main.go`                        | entry-point of the server                                                                              |
| `Makefile`                       | stores frequently used commands; can be invoked using `make` command                                   |

---

## 🚀 Running Migrations

This project uses [Atlas](https://atlasgo.io/) for database schema migrations. Atlas enables declarative, versioned, and diff-based schema changes.

---

### 🧰 Prerequisites

Make sure you have the following set up:

- **Atlas CLI**: Install Atlas by running:

  ```sh
  curl -sSf https://atlasgo.sh | sh
  ```

  > For other installation methods or details, visit the [official installation guide](https://atlasgo.io/getting-started/installation).

- **`.env` file** at the project root with the following environment variables:

  ```env
  DB_USER=root
  DB_PASS=secret
  DB_NAME=exampledb
  DB_FORWARD_PORT=3306
  ```

---

### 📦 Available Migration Commands

Below are the supported `make` commands for managing database migrations:

| Make Command          | Description                                                                 |
| --------------------- | --------------------------------------------------------------------------- |
| `make migrate-status` | Show the current migration status                                           |
| `make migrate-diff`   | Generate a new migration by comparing models to the current DB (`gorm` env) |
| `make migrate-apply`  | Apply all pending migrations                                                |
| `make migrate-down`   | Roll back the most recent migration (`gorm` env)                            |
| `make migrate-hash`   | Hash migration files for integrity checking                                 |

---

📚 For more on schema management and best practices, refer to the [Atlas documentation](https://atlasgo.io).

## Testing

The framework comes with unit and integration testing support out of the box. You can check examples written in tests directory.

To run the test just run:

```zsh
go test ./... -v
```

### For test coverage

```zsh
go test ./... -v -coverprofile cover.txt -coverpkg=./...
go tool cover -html=cover.txt -o index.html
```

### Update Dependencies
See [UPDATING_DEPENDENCIES.md](./UPDATING_DEPENDENCIES.md) file for more information on how to update project dependencies.




### Contribute 👩‍💻🧑‍💻

We are happy that you are looking to improve go clean architecture. Please check out the [contributing guide](contributing.md)

Even if you are not able to make contributions via code, please don't hesitate to file bugs or feature requests that needs to be implemented to solve your use case.

### Authors

<div align="center">
    <a href="https://github.com/wesionaryTEAM/go_clean_architecture/graphs/contributors">
        <img src="https://contrib.rocks/image?repo=wesionaryTEAM/go_clean_architecture" />
    </a>
</div>
