package controller

import (
	mongodb "github.com/cjlapao/common-go/database/mongo"
	"github.com/cjlapao/common-go/executionctx"
	"github.com/cjlapao/common-go/identity"
)

var _factory *mongodb.MongoFactory
var _userRepo mongodb.Repository
var svc = executionctx.GetServiceProvider()

type UserContext struct{}

func (u UserContext) GetUserById(id string) *identity.User {
	var result identity.User
	repo := GetRepository()
	dbUsers := repo.FindOne("id", id)
	dbUsers.Decode(&result)
	return &result
}

func (u UserContext) GetUserByEmail(email string) *identity.User {
	var result identity.User
	repo := GetRepository()
	dbUsers := repo.FindOne("email", email)
	dbUsers.Decode(&result)
	return &result
}

func (u UserContext) UpsertUser(user identity.User) {

}

func GetRepository() mongodb.Repository {
	if _factory == nil {
		connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
		databaseName := svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
		_factory = mongodb.NewFactory(connStr).WithDatabase(databaseName)
		_userRepo = mongodb.NewRepository(_factory, "test", identity.IdentityUsersCollection)
	}

	return _userRepo
}
