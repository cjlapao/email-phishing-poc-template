package startup

import (
	"github.com/cjlapao/common-go/identity"
	"github.com/cjlapao/phishing-email-backend/controller"
	"github.com/cjlapao/phishing-email-backend/databaseservice"
	"github.com/cjlapao/phishing-email-backend/entities"
	"github.com/cjlapao/phishing-email-backend/repositories"
)

func Init() {
	dbFactory, databaseName := databaseservice.GetDatabase()
	userRepo := repositories.UserRepository()
	user1, _ := entities.NewUser("cjlapao@gmail.com")
	user2, _ := entities.NewUser("carlos.lapao@ivanti.com")
	userRepo.Repository.UpsertOne("email", user1.Email, user1)
	userRepo.Repository.UpsertOne("email", user2.Email, user2)
	identity.Seed(dbFactory, databaseName)
	controller.Init()
}
