### Go Clean Architecture
Clean Architecture with Golang with Dependency Injection


### Use of Linter in the project
To install all the packages and run the linter in git pre-commit hook; Run
> Make lint-setup
<br/>


### Run Migration Commands
> ⚓️ &nbsp; Add argument `p=host` after `make` if you want to run the migration runner from the host environment instead of docker environment. example; `make p=host migrate-up`

If you are not using docker; ensure that sql-migrate is installed to use migration from the host environment.
To install sql-migrate:
> go get -v github.com/rubenv/sql-migrate/...

<details>
    <summary>Migration commands available</summary>

| Command              | Desc                                                       |
| -------------------- | ---------------------------------------------------------- |
| `make migrate-status`| Show migration status                                      |
| `make migrate-up`    | Migrates the database to the most recent version available |
| `make migrate-down`  | Undo a database migration                                  |
| `make redo`          | Reapply the last migration                                 |
| `make create`        | Create new migration file                                  |

</details>


### Run app with docker
- Update database env variables with credentials defined in `docker-compose.yml`
- Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if permission issues
    > Assumes: Docker is already installed in the machine. 

## Checking API documents with swagger UI
Browse to `http://localhost:${SWAGGER_PORT}`
- You can see all the documented endpoints in Swagger-UI from the API specification
- You can execute/test endpoint
You can read the article to know more on this: https://medium.com/wesionary-team/swagger-ui-on-docker-for-testing-rest-apis-5b3d5fcdee7

## Update Dependencies
<details>
    <summary><b>Steps to Update Dependencies</b></summary>
    
1. ```go get -u```
2. Remove all the dependencies packages that has ```// indirect``` from the modules
3. ```go mod tidy```
</details>

<details>
    <summary><b>Discovering available updates</b></summary>
    
List all of the modules that are dependencies of your current module, along with the latest version available for each:
> $ go list -m -u all

Display the latest version available for a specific module:
> $ go list -m -u example.com/theirmodule

<b>Example:</b>
> $ go list -m -u cloud.google.com/go/firestore<br/>
cloud.google.com/go/firestore v1.2.0 [v1.6.1]
</details>

<details>
    <summary><b>Getting a specific dependency version</b></summary>
    
To get a specific numbered version, append the module path with an @ sign followed by the version you want:
> $ go get example.com/theirmodule@v1.3.4

To get the latest version, append the module path with @latest:
> $ go get example.com/theirmodule@latest
</details>

<details>
    <summary><b>Synchronizing your code’s dependencies</b></summary>
 
> $ go mod tidy
</details>

