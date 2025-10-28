package infrastructure

import "clean-architecture/pkg/framework"

// NewDBProvider is an Fx constructor that provides the concrete DBProvider.
func NewDBProvider(env *framework.Env) DBProvider {
	return providerDBFactory(env)
}

// NewDatabase receives a DBProvider via Fx and returns a connected Database.
func NewDatabase(logger framework.Logger, env *framework.Env, provider DBProvider) Database {
	db, err := provider.Connect(logger, env)
	if err != nil {
		logger.Panic("failed to connect database: ", err)
	}
	return *db
}
