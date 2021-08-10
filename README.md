### Go Clean Architecture
Clean Architecture with Golang with Dependency Injection

### Run app with docker
- Update database env variables with credentials defined in `docker-compose.yml`
- Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if permission issues
    > Assumes: Docker is already installed in the machine. 

## Checking API documents with swagger UI
Browse to http://localhost:${SWAGGER_PORT}
- You can see all the documented endpoints in Swagger-UI from the API specification
- You can execute/test endpoint