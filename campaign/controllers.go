package campaign

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCampaignController(w http.ResponseWriter, r *http.Request) {
	repo := NewCampaignRepository()
	campaign := repo.Get()
	if len(campaign) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(campaign)
}

func GetCampaignByIdController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	repo := NewCampaignRepository()
	campaign := repo.GetById(id)
	if campaign.ID != "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(campaign)
}

func PostCampaignController(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	campaign := Campaign{}
	err := json.Unmarshal(reqBody, &campaign)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := NewCampaignRepository()
	upsertResult := repo.Upsert(campaign)
	if !upsertResult {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(campaign)
}
