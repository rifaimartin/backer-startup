package transaction

import "backer-startup/user"

//GetCampaignDetailInput this is struct
type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
