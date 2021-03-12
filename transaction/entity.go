package transaction

import (
	"backer-startup/campaign"
	"backer-startup/user"
	"time"

	"github.com/leekchan/accounting"
)

//Transaction struct
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//AmountFormatIDR fomat idr
func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "IDR ", Precision: 0, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
