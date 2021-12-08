package startup

import (
	"encoding/json"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/identity"
	"github.com/cjlapao/phishing-email-backend/controller"
	"github.com/cjlapao/phishing-email-backend/databaseservice"
)

func Init() {
	dbFactory, databaseName := databaseservice.GetDatabase()
	// userRepo := repositories.UserRepository()
	// user1, _ := entities.NewUser("cjlapao@gmail.com")
	// user2, _ := entities.NewUser("carlos.lapao@ivanti.com")
	// userRepo.Repository.UpsertOne("email", user1.Email, user1)
	// userRepo.Repository.UpsertOne("email", user2.Email, user2)
	// test := mongodb.NewFilterBuilder()
	// test.And("lastName", "User")
	// test.Or("firstName", "Administrator")
	// test.Or("firstName", "Demo")
	// test.And("_id", "592D8E5C-6F5D-40A0-9348-80131B083715")
	// result := test.Build()

	builder := mongodb.NewFilterBuilder()
	builder.Field("test").GreaterThan(9).LowerThan(21).Build()
	result := builder.Build()
	j, _ := json.Marshal(result)
	println(j)
	identity.Seed(dbFactory, databaseName)
	controller.Init()
}
