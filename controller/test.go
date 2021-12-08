package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go/executionctx"
	"github.com/cjlapao/phishing-email-backend/mailservice"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	mailSvc := mailservice.NewMailService()
	mailSvc.From = "cjlapao@gmail.com"
	mailSvc.Send("cjlapao@gmail.com", "testing templates", "test")
	response := "Healthy"
	ctx := executionctx.GetContext()
	if ctx.User != nil {
		response += " " + ctx.User.Name
	}
	json.NewEncoder(w).Encode(response)
}

func AuthTestController(w http.ResponseWriter, r *http.Request) {
	ctx := executionctx.GetContext()
	response := "Healthy"
	if ctx.User != nil {
		response += " " + ctx.User.Name
	}
	json.NewEncoder(w).Encode(response)
}
