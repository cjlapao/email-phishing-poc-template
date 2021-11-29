package controller

import "github.com/cjlapao/common-go/restapi"

var listener *restapi.HttpListener

func Init() {
	listener = restapi.GetHttpListener()
	listener.AddLogger().AddAuthentication().AddHealthCheck()
	listener.AddController(TestController, "/test", "GET")
	listener.Start()
}
