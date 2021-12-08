package repositories

import (
	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/executionctx"
)

// Collection
const (
	UsersCollectionName    = "Users"
	CampainsCollectionName = "Campains"
)

// Repository Entity
type Repository struct {
	Factory    *mongodb.MongoFactory
	Repository mongodb.Repository
}

// NewRepo Creates a new Article Repository object
func NewRepository(collectionName string) *Repository {
	svc := executionctx.GetServiceProvider()
	connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
	databaseName := svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
	result := Repository{
		Factory: mongodb.NewFactory(connStr).WithDatabase(databaseName),
	}

	result.Repository = mongodb.NewRepository(result.Factory, databaseName, collectionName)

	return &result
}
