package startup

import (
	mongodb "github.com/cjlapao/common-go/database/mongo"
	"github.com/cjlapao/common-go/executionctx"
	"github.com/cjlapao/phishing-email-backend/controller"
	"github.com/cjlapao/phishing-email-backend/entities"
)

var svc = executionctx.GetServiceProvider()

func Init() {
	connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
	databaseName := svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
	factory := mongodb.NewFactory(connStr).WithDatabase(databaseName)

	user, userErr := entities.NewUser("test@example.com")
	user2, user2Err := entities.NewUser("test2@example.com")
	if userErr == nil && user2Err == nil {
		factory.InsertOne("user", user)
		factory.InsertOne("user", user2)
	}

	controller.Init()
}
