package infrastructure

import "clean-architecture/pkg/framework"

// MySQLProvider implements DBProvider for MySQL.
type MySQLProvider struct{}

func (p *MySQLProvider) Type() string { return "mysql" }

func (p *MySQLProvider) Connect(logger framework.Logger, env *framework.Env) (*Database, error) {
	db, err := MySQLConnect(logger, env)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}
