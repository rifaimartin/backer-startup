package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

//Repository interface
type Repository interface {
	GetByCampaignID(CampaignID int) ([]Transaction, error)
}

//NewRepository semothing
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}
