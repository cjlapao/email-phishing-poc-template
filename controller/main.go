package controller

import (
	"github.com/cjlapao/common-go/restapi"
	"github.com/cjlapao/phishing-email-backend/campaign"
)

var listener *restapi.HttpListener

func Init() {
	userCtx := UserContext{}
	listener = restapi.GetHttpListener()
	listener.WithAuthentication(userCtx).AddLogger().AddHealthCheck()
	listener.AddController(TestController, "/test", "GET")

	listener.AddAuthorizedController(AuthTestController, "/auth/test", "GET")

	listener.AddController(campaign.GetCampaignController, "/campaigns", "GET")
	listener.AddController(campaign.GetCampaignByIdController, "/campaigns/{id}", "GET")
	listener.AddController(campaign.PostCampaignController, "/campaign", "POST")
	listener.Start()
}
