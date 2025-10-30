package infrastructure

import (
	"clean-architecture/pkg/framework"
	"strings"
)

// DBProvider interface
type DBProvider interface {
	Connect(logger framework.Logger, env *framework.Env) (*Database, error)
	Type() string
}

// providerDBFactory chooses correct provider based on env.DBType.
func providerDBFactory(env *framework.Env) DBProvider {
	dbType := strings.ToLower(strings.TrimSpace(env.Database.Type))
	if dbType == "" {
		return &MySQLProvider{}
	}
	switch dbType {
	case "mysql":
		return &MySQLProvider{}
	case "postgres", "postgresql":
		return &PostgresProvider{}
	default:
		return &MySQLProvider{} // fallback for unknown types
	}
}
