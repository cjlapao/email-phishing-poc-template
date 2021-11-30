package repositories

import (
	database "github.com/cjlapao/common-go/database/mongo"
	"github.com/cjlapao/common-go/executionctx"
)

// Collection
const (
	UsersCollectionName    = "Users"
	CampainsCollectionName = "Campains"
)

// Repository Entity
type Repository struct {
	Factory    *database.MongoFactory
	Repository database.Repository
}

// NewRepo Creates a new Article Repository object
func NewRepository(collectionName string) *Repository {
	svc := executionctx.GetServiceProvider()
	connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
	databaseName := svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
	result := Repository{
		Factory: database.NewFactory(connStr).WithDatabase(databaseName),
	}

	result.Repository = database.NewRepository(result.Factory, databaseName, collectionName)

	return &result
}
