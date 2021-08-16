### Go Clean Architecture
Clean Architecture with Golang with Dependency Injection


### Run Migration Commands
> ⚓️ &nbsp; Add argument `p=host` after `make` if you want to run the migration runner from the host environment instead of docker environment. example; `make p=host migrate-up`

If you are not using docker; ensure that migrate is installed to use migration from the host environment. Browse https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation to install migrate and run migrate command.

<details>
    <summary>Migration commands available</summary>

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

</details>


### Run app with docker
- Update database env variables with credentials defined in `docker-compose.yml`
- Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if permission issues
    > Assumes: Docker is already installed in the machine. 

## Checking API documents with swagger UI
Browse to http://localhost:${SWAGGER_PORT}
- You can see all the documented endpoints in Swagger-UI from the API specification
- You can execute/test endpoint
You can read the article to know more on this: https://medium.com/wesionary-team/swagger-ui-on-docker-for-testing-rest-apis-5b3d5fcdee7