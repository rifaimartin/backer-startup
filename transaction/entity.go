package transaction

import (
	"backer-startup/campaign"
	"backer-startup/user"
	"time"
)

//Transaction struct
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
