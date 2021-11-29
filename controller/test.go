package controller

import (
	"encoding/json"
	"net/http"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	response := "Healthy"
	json.NewEncoder(w).Encode(response)
}
