package campaign

import (
	mongodb "github.com/cjlapao/common-go/database/mongo"
	"github.com/cjlapao/phishing-email-backend/constants"
	"github.com/cjlapao/phishing-email-backend/databaseservice"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

var repository mongodb.Repository

type CampaignRepository struct {
	Collection mongodb.Repository
}

func NewCampaignRepository() *CampaignRepository {
	factory, dbName := databaseservice.GetDatabase()
	repo := CampaignRepository{
		Collection: mongodb.NewRepository(factory, dbName, constants.CampaignCollection),
	}
	return &repo
}

func (c *CampaignRepository) Get() []Campaign {
	campaigns := make([]Campaign, 0)
	result := c.Collection.Filter("")
	for _, rec := range result {
		var campaign Campaign
		mapstructure.Decode(rec, campaign)
	}

	return campaigns
}

func (c *CampaignRepository) GetById(id string) Campaign {
	var campaign Campaign
	result := c.Collection.FindOne("_id", id)
	result.Decode(&campaign)

	return campaign
}

func (c *CampaignRepository) Upsert(campaign Campaign) bool {
	if campaign.ID == "" {
		campaign.ID = uuid.NewString()
	}

	updateResult := c.Collection.UpsertOne("_id", campaign.ID, campaign)

	return updateResult.UpsertedCount == 1
}
