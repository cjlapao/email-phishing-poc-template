package databaseservice

import (
	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/executionctx"
)

var svc = executionctx.GetServiceProvider()
var dbFactory *mongodb.MongoFactory
var databaseName string

func GetDatabase() (*mongodb.MongoFactory, string) {
	if dbFactory != nil {
		return dbFactory, databaseName
	}
	svc := executionctx.GetServiceProvider()
	connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
	databaseName = svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
	dbFactory = mongodb.NewFactory(connStr).WithDatabase(databaseName)
	return dbFactory, databaseName
}
