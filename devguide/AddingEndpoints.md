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
			//If you want to enable auto-migrate add Migrate as shown below
			// fx.Invoke(Migrate, RegisterRoute),

			fx.Invoke(RegisterRoute),
		))
  ```
- The `domain/<feature_name>/module.go` module, should be linked with `domain/module.go` so that, it is added to dependency injection tree.

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