## Adding API Endpoint in the Architecture

- If the package name is not known, read the package name from the `go.mod` file when importing internal packages.
- If a new feature is required, create it inside `domain/<feature_name>/`. A feature generally includes a controller, route, service, module, serializer (DTO), and repository, all in separate files.
- Strictly adhere to the request and response structure if present. You need to create a DTO layer to convert the database structure (created in `models/`) to the request/response structure (created in `domain/<feature_name>/`).
- Before adding dependencies for the controller, route, service, repositories, etc., check how it is done in other features and make similar changes. Especially check for pointer or non-pointer dependencies in the return type of the provider function for the dependencies. For example:

  ```go
  package infrastructure

  // NewDatabase creates a new database instance
  func NewDatabase(logger framework.Logger, env *framework.Env) Database {
  }
  ```

  The dependency `Database` is a non-pointer type, whereas `env` is a pointer type.

- After all these files are created, the module should contain dependency injection setup for the created controller, route, service, and repository. Route registration is done using `fx.Invoke`, which runs the route registration function on application start automatically. For example:

  ```go
  var Module = fx.Module("user",
      fx.Options(
          fx.Provide(
              NewRepository,
              NewService,
              NewController,
              NewRoute,
          ),
          fx.Invoke(RegisterRoute),
      ))
  ```

- The `domain/<feature_name>/module.go` module should be linked with `domain/module.go` so that it is added to the dependency injection tree.

### Defining Routes in the Framework

To define routes in the framework, you need to create a route file inside the `domain/<feature_name>/` folder. The route file should include the necessary imports, route definitions, and a function to register the routes. Below is an updated example of how to define routes for a feature:

```go
package <feature_name>

import (
    "clean-architecture/pkg/framework"
    "clean-architecture/pkg/infrastructure"
)

// Route struct
type Route struct {
    logger     framework.Logger
    handler    infrastructure.Router
    controller *Controller
}

// NewRoute creates a new Route instance
func NewRoute(
    logger framework.Logger,
    handler infrastructure.Router,
    controller *Controller,
) *Route {
    return &Route{
        handler:    handler,
        logger:     logger,
        controller: controller,
    }
}

// RegisterRoute sets up the routes for the feature
func RegisterRoute(r *Route) {
    r.logger.Info("Setting up routes")

    api := r.handler.Group("/api")

    api.POST("/<feature_name>", r.controller.Create)
    api.GET("/<feature_name>/:id", r.controller.GetByID)
}
```

### Explanation

1. **Route Struct**: The `Route` struct encapsulates the logger, router, and controller dependencies required for setting up routes.
2. **Route Initialization**: The `NewRoute` function initializes a new `Route` instance with the required dependencies.
3. **Route Registration**: The `RegisterRoute` function defines the HTTP methods and their corresponding handler functions for the feature.
4. **Dynamic Parameters**: Use `:id` in the route path to define dynamic parameters.

## Adding New Models for a Feature

- For adding new database models for a feature, models are added to the `domain/models` folder.
- After adding models, it is essential to diff the database with models and generate migration using Atlas Go. Since the Makefile already contains the command for migration, you can check it.
- The generated migrations need to be run as well.
- Some data types for new model generation:
  - UUID -> `types.BinaryUUID`
- The database we are using is MySQL, so other variants of SQL in model definition might not work.

Here is a sample model definition that you might need:

```go
package models

import (
    "clean-architecture/domain/constants"
    "clean-architecture/pkg/types"

    _ "ariga.io/atlas-provider-gorm/gormschema"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User model
type User struct {
    gorm.Model
    UUID       types.BinaryUUID `json:"uuid" gorm:"index;notnull;unique"`
    CognitoUID *string          `json:"-" gorm:"index;size:50;unique"`

    FirstName   string `json:"first_name" gorm:"size:255"`
    LastName    string `json:"last_name" gorm:"size:255"`

    Email string             `json:"email" gorm:"notnull;index,unique;size:255"`
    Role  constants.UserRole `json:"role" gorm:"size:25" copier:"-"`
}

// BeforeCreate auto-generates a UUID before creating if it's not present already
func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.UUID.String() == (types.BinaryUUID{}).String() {
        id, err := uuid.NewRandom()
        u.UUID = types.BinaryUUID(id)
        return err
    }
    return nil
}

func (*User) TableName() string {
    return "users"
}
```

### 📦 Available Migration Commands

Below are the supported `make` commands for managing database migrations:

| Make Command          | Description                                                                 |
| --------------------- | --------------------------------------------------------------------------- |
| `make migrate-status` | Show the current migration status                                           |
| `make migrate-diff`   | Generate a new migration by comparing models to the current DB (`gorm` env) |
| `make migrate-apply`  | Apply all pending migrations                                                |
| `make migrate-down`   | Roll back the most recent migration (`gorm` env)                            |
| `make migrate-hash`   | Hash migration files for integrity checking                                 |