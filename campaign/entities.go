package campaign

import "github.com/google/uuid"

type Campaign struct {
	ID      string           `json:"id" bson:"_id"`
	Name    string           `json:"name" bson:"name"`
	Targets CampaignsTargets `json:"targets" bson:"targets"`
	Summary CampaignSummary  `json:"summary" bson:"summary"`
}

type CampaignsTargets struct {
	LowThreshold     int
	MediumThreshold  int
	HighThreshold    int
	CurrentThreshold int
}

type CampaignSummary struct {
	Users       []string
	UsersPawned []string
}

func NewCampaign() Campaign {
	return Campaign{
		ID:      uuid.NewString(),
		Targets: CampaignsTargets{},
		Summary: CampaignSummary{},
	}
}
