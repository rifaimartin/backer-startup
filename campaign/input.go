package campaign

//GetCampaignDetailInput this is struct
type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
