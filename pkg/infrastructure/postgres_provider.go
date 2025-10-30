package infrastructure

import "clean-architecture/pkg/framework"

// PostgresProvider implements DBProvider for PostgreSQL.
type PostgresProvider struct{}

func (p *PostgresProvider) Type() string { return "postgres" }

func (p *PostgresProvider) Connect(logger framework.Logger, env *framework.Env) (*Database, error) {
	db, err := PostgresConnect(logger, env)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}
