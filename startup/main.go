package startup

import (
	"github.com/cjlapao/common-go/executionctx"
	"github.com/cjlapao/phishing-email-backend/controller"
	"github.com/cjlapao/phishing-email-backend/entities"
	"github.com/cjlapao/phishing-email-backend/repositories"
)

var svc = executionctx.GetServiceProvider()

func Init() {
	userRepo := repositories.UserRepository()
	user1, _ := entities.NewUser("cjlapao@gmail.com")
	user2, _ := entities.NewUser("carlos.lapao@ivanti.com")
	userRepo.Repository.UpsertOne("email", user1.Email, user1)
	userRepo.Repository.UpsertOne("email", user2.Email, user2)

	controller.Init()
}
