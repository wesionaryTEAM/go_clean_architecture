## Adding API Endpoint in the architecture 

- If package name is not known, read package name from `go.mod` file when importing internal packages.
- If new feature is required, create inside `domain/<feature_name>/`. Feature generally has controller, route, service, module, serializer (dto) and repository all in separate files.
- Strictly adhere to the request and response structure if present, you need to create DTO layer to convert the db structure (created at `models/`) to req/resp structure (created at `domain/<feature_name>/`).
- Before adding dependencies for controller, route, service, repositories etc. check how its done in other features and make similar changes. Especially check for pointer or non-pointer dependencies in return type of provider function for the dependencies. 
  for example:
  ```go 
    package infrastructure 

    // NewDatabase creates a new database instance
    func NewDatabase(logger framework.Logger, env *framework.Env) Database {
    }
  ```
  the dependency Database is non pointer type. where as env is pointer type. 

- After all these files are created, the module should contain dependency injection setup for created controller, route, service and repository. Route registration is done using `fx.Invoke` which runs route registraion function on application start automatically.
  for example:
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
- The `domain/<feature_name>/module.go` module, should be linked with `domain/module.go` so that, it is added to dependency injection tree.

### Example of Using `infrastructure.Database` in Repository

When creating a repository, you can use `infrastructure.Database` to interact with the database. Below is an example of how to create a repository:

```go
package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// Repository represents the user repository structure
type Repository struct {
	infrastructure.Database
	logger framework.Logger
}

// NewRepository initializes a new user repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

// ExampleMethod demonstrates a database query using infrastructure.Database
func (r *Repository) ExampleMethod() error {
	r.logger.Info("[Repository...ExampleMethod]")

	var users []models.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return err 
	}

	return nil
}
```

### Defining Routes for a Feature

To define routes for a feature, you need to create a `route.go` file inside the respective feature's directory (e.g., `domain/<feature_name>/route.go`). Below is an example of how routes are defined for the `user` domain:

```go
package user

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

// NewRoute initializes a new Route instance
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

	api.POST("/user", r.controller.CreateUser)
	api.GET("/user/:id", r.controller.GetUserByID)
}
```

#### Key Points:
- The `Route` struct holds dependencies like the logger, router, and controller.
- The `NewRoute` function initializes the `Route` struct.
- The `RegisterRoute` function defines the actual routes and associates them with controller methods.
- Use the `Group` method of the router to group routes under a common prefix (e.g., `/api`).

This structure ensures that routes are modular and easy to manage for each feature.

### Using `fx.Invoke` for Route Registration

After defining the routes for a feature, you need to register them in the module file using `fx.Invoke`. This ensures that the route registration function is automatically executed when the application starts. Below is an example of how this is done for the `user` domain:

```go
package user

import (
	"go.uber.org/fx"
)

// Module provides the dependencies for the user domain
var Module = fx.Module("user",
	fx.Provide(
		NewRepository,
		NewService,
		NewController,
		NewRoute,
	),
	fx.Invoke(RegisterRoute),
)
```

#### Key Points:
- The `fx.Provide` function is used to declare the dependencies (e.g., repository, service, controller, and route) for the feature.
- The `fx.Invoke` function is used to call the `RegisterRoute` function, which sets up the routes for the feature.
- This setup ensures that the routes are registered as part of the application's dependency injection lifecycle.

By following this pattern, you can maintain a clean and modular structure for route registration in your application.

## Adding new models for a feature

- For adding new db models for a feature, models are added to `domain/models` folder. 
- After adding models, it is essential to diff the database with models and generate migration using atlas go. Since, makefile already contains the command for migration, you can check it. 
- The generated migrations, need to be run as well. 
- Some datatypes for new model generation; 
  UUID -> types.BinaryUUID
- Database we are using in MySQL so other variant of SQL in model definition might not work.

Here is a sample model definition, that you might need.
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

// BeforeCreate auto generate uuid before creating if it's not present already
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