# Go Clean Architecture

Clean Architecture with [Gin Web Framework](https://github.com/gin-gonic/gin)

## Features :star:

-   Clean Architecture written in Go
-   Application backbone with [Gin Web Framework](https://github.com/gin-gonic/gin)
-   Dependency injection using [uber-go/fx](https://pkg.go.dev/go.uber.org/fx)
-   API endpoints documentation with [Swagger-UI](https://swagger.io/tools/swagger-ui/download/)
-   Uses fully featured [GORM](https://gorm.io/index.html)

## Linter setup

To add linter in git pre-commit hook, and install all the required packages for setting up linter.

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
- Setup `serviceAccountKey.json`. To get one create a firebase project. Go to Settings > Service Accounts and then click **"Generate New Private Key"**. and then confirm by clicking **"Generate Key"**.
Copy the key to `serviceAccountKey.json` file. You can see the example at `serviceAccountKey.json.example` file. 

### Locally

-   Run `go run main.go app:serve` to start the server.
-   There are other commands available as well. You can run `go run main.go -help` to know about other commands available.

### Using `Docker`

> Ensure Docker is already installed in the machine.

-   Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if there are permission issues.

---

## Folder Structure :file_folder:

| Folder Path                      | Description                                                                                         |
| -------------------------------- | --------------------------------------------------------------------------------------------------- |
| `/api`                           | contains all the `middlwares`, `controllers` and `routes` of the server in their respective folders |
| `/api-errors`                    | server error handlers                                                                               |
| `/bootstrap`                     | contains modules required to start the application                                                  |
| `/console`                       | server commands, run `go run main.go -help` for all the available server commands                   |
| `/constants`                     | global application constants                                                                        |
| `/docker`                        | `docker` files required for `docker compose`                                                        |
| `/docs`                          | API endpoints documentation using `swagger`                                                         |
| `/hooks`                         | `git` hooks                                                                                         |
| `/infrastructure`                | third-party services connections like `gmail`, `firebase`, `s3-bucket`, ...                         |
| `/lib`                           | contains library code                                                                               |
| `/migration`                     | database migration files                                                                            |
| `/models`                        | ORM models                                                                                          |
| `/repository`                    | contains repository part of clean architecture. Mainly database queries are added here.             |
| `/seeds`                         | seeds for already migrated tables                                                                   |
| `/services`                      | service layers, contains the functionality that compounds the core of the application               |
| `/tests`                         | includes application tests                                                                          |
| `/utils`                         | global utility/helper functions                                                                     |
| `.env.example`                   | sample environment variables                                                                        |
| `dbconfig.yml`                   | database configuration file for `sql-migrate` command                                               |
| `docker-compose.yml`             | `docker compose` file for service application via `Docker`                                          |
| `main.go`                        | entry-point of the server                                                                           |
| `Makefile`                       | stores frequently used commands; can be invoked using `make` command                                |
| `serviceAccountKey.json.example` | sample credentials file for accessing Google Cloud                                                  |

---

## Migration Commands

⚓️ &nbsp; If you want to run the migration runner from the host environment instead of the docker environment; ensure that `sql-migrate` is installed on your local machine.

### Install `sql-migrate`

> You can skip this step if `sql-migrate` has already been installed on your local machine.

**Note:** Starting in Go 1.17, installing executables with `go get` is deprecated. `go install` may be used instead. [Read more](https://go.dev/doc/go-get-install-deprecation)

```zsh
go install github.com/rubenv/sql-migrate/...@latest
```

If you're using Go version below `1.18`

```zsh
go get -v github.com/rubenv/sql-migrate/...
```

### Running migration

Add argument `p=host` after `make` command to run migration commands on local environment

<b>Example:</b>

```zsh
make p=host migrate-up
```

<details>
    <summary>Available migration commands</summary>

| Command               | Desc                                                       |
| --------------------- | ---------------------------------------------------------- |
| `make migrate-status` | Show migration status                                      |
| `make migrate-up`     | Migrates the database to the most recent version available |
| `make migrate-down`   | Undo a database migration                                  |
| `make redo`           | Reapply the last migration                                 |
| `make create`         | Create new migration file                                  |

</details>

---

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

---

## Checking API documents with swagger UI

Browse to `http://localhost:${SWAGGER_PORT}`

-   You can see all the documented endpoints in Swagger-UI from the API specification
-   You can execute/test endpoint. [Read article](https://medium.com/wesionary-team/swagger-ui-on-docker-for-testing-rest-apis-5b3d5fcdee7)

---

## Update Dependencies

<details>
    <summary><b>Steps to Update Dependencies</b></summary>
    
1. `go get -u`
2. Remove all the dependencies packages that has `// indirect` from the modules
3. `go mod tidy`
</details>

<details>
    <summary><b>Discovering available updates</b></summary>
    
List all of the modules that are dependencies of your current module, along with the latest version available for each:
```zsh 
go list -m -u all
```

Display the latest version available for a specific module:

```zsh
go list -m -u example.com/theirmodule
```

<b>Example:</b>

```zsh
go list -m -u cloud.google.com/go/firestore
cloud.google.com/go/firestore v1.2.0 [v1.6.1]
```

</details>

<details>
    <summary><b>Getting a specific dependency version</b></summary>
    
To get a specific numbered version, append the module path with an `@` sign followed by the `version` you want:

```zsh
go get example.com/theirmodule@v1.3.4
```

To get the latest version, append the module path with @latest:

```zsh
go get example.com/theirmodule@latest
```

</details>

<details>
    <summary><b>Synchronizing your code’s dependencies</b></summary>
 
```zsh
go mod tidy
```
</details>
