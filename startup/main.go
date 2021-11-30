package startup

import (
	database "github.com/cjlapao/common-go/database/mongo"
	"github.com/cjlapao/common-go/executionctx"
	"github.com/cjlapao/common-go/identity"
	"github.com/cjlapao/phishing-email-backend/controller"
	"github.com/cjlapao/phishing-email-backend/entities"
	"github.com/cjlapao/phishing-email-backend/repositories"
)

var svc = executionctx.GetServiceProvider()
var dbFactory *database.MongoFactory

func Init() {
	svc := executionctx.GetServiceProvider()
	connStr := svc.Context.Configuration.GetString("MONGODB_CONNECTION_STRING")
	databaseName := svc.Context.Configuration.GetString("MONGODB_DATABASENAME")
	dbFactory = database.NewFactory(connStr).WithDatabase(databaseName)
	userRepo := repositories.UserRepository()
	user1, _ := entities.NewUser("cjlapao@gmail.com")
	user2, _ := entities.NewUser("carlos.lapao@ivanti.com")
	userRepo.Repository.UpsertOne("email", user1.Email, user1)
	userRepo.Repository.UpsertOne("email", user2.Email, user2)
	identity.Seed(dbFactory, databaseName)
	controller.Init()
}
